package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	handlers "github.com/niko0xdev/nx-ddns/internal/app/handler"
	"github.com/niko0xdev/nx-ddns/internal/config"
	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/repository"
	"github.com/niko0xdev/nx-ddns/internal/utils"
	"github.com/niko0xdev/nx-ddns/pkg/ddns"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/niko0xdev/nx-ddns/cmd/nxddns/docs"
)

// @BasePath /api
func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	database.InitDBConnection(cfg)

	// Check database connection
	if err := database.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection successfully established")

	// Initialize Snowflake ID generator
	// TODO: add node id here
	utils.InitSnowflake(1)

	// Initialize repository and handler
	dnsHandler := handlers.NewDNSHandler()

	// Set up the Gin router
	r := gin.Default()

	// Serve Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API Group
	api := r.Group("/api")
	{
		// Register routes
		api.GET("/records", dnsHandler.GetDNSRecords)
		api.GET("/records/:id", dnsHandler.GetDNSRecord)
		api.POST("/records", dnsHandler.CreateDNSRecord)
		api.PUT("/records/:id", dnsHandler.UpdateDNSRecord)
		api.DELETE("/records/:id", dnsHandler.DeleteDNSRecord)

		// register change logs tracking
		api.GET("/logs/:dnsRecordId", dnsHandler.GetDNSLogs)
	}

	// Start background task to check IP changes
	go startIPChangeChecker(cfg)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server with graceful shutdown support
	log.Printf("Starting server on port %s...", port)

	// Create a channel to listen for termination signals (like SIGINT)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	// Start server in a goroutine to handle shutdown signals
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for SIGINT or SIGTERM signals for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a deadline to wait for active connections to finish
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}

	log.Println("Server gracefully stopped")
}

// startIPChangeChecker runs the check_ip_changes function every minute
func startIPChangeChecker(cfg *config.Config) {
	ticker := time.NewTicker(ddns.IpCheckInterval)
	defer ticker.Stop()

	// init db repository
	repo := repository.NewDNSRecordRepository(database.DB)

	for range ticker.C {
		fmt.Println("Checking for IP changes...")
		// get active dns records
		records, err := repo.GetDNSRecords()
		if err != nil {
			panic(err)
		}

		// get current public ip
		ip, err := utils.GetPublicIP()
		if err != nil {
			panic(err)
		}

		// update dns records
		for _, record := range records {

			if record.IPAddress != ip {
				// update dns record
				_, err := ddns.UpdateDNSRecord(repo, &record, ip, cfg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
