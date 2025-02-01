package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/niko0xdev/nx-ddns/internal/app/dto"
	"github.com/niko0xdev/nx-ddns/internal/database"
	"github.com/niko0xdev/nx-ddns/internal/repository"
	"github.com/niko0xdev/nx-ddns/internal/utils"
)

type DNSHandler struct {
	repository *repository.DNSRecordRepository
}

func NewDNSHandler() *DNSHandler {
	return &DNSHandler{repository: repository.NewDNSRecordRepository(database.DB)}
}

// @Summary Get all DNS records
// @Description Retrieve a list of all DNS records
// @Tags DNSRecords
// @Produce json
// @Success 200 {array} dto.DNSRecord
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /records [get]
func (handler *DNSHandler) GetDNSRecords(c *gin.Context) {
	records, err := handler.repository.GetDNSRecords()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, records)
}

// @Summary Create a new DNS record
// @Description Creates a new DNS record and logs the creation
// @Tags DNSRecords
// @Accept json
// @Produce json
// @Param dnsRecord body dto.DNSRecordRequest true "DNS Record"
// @Success 201 {object} dto.DNSRecord
// @Failure 400 {object} dto.ErrorResponse "Invalid input"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /records [post]
func (handler *DNSHandler) CreateDNSRecord(c *gin.Context) {
	var request dto.DNSRecordRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create DNS record
	dnsRecord := dto.DNSRecordFromDTO(request)
	record, err := handler.repository.CreateDNSRecord(&dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Log the creation of this DNS record
	dnsLog := database.DNSLog{
		BaseModel: database.BaseModel{
			ID:        utils.GenerateID().String(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		DNSRecordID: record.ID,
		OldValue:    "",
		NewValue:    record.IPAddress,
	}

	_, err = handler.repository.CreateDNSLog(&dnsLog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, record)
}

// @Summary Get a DNS record by ID
// @Description Retrieve a specific DNS record by ID
// @Tags DNSRecords
// @Produce json
// @Param id path int true "DNS Record ID"
// @Success 200 {object} dto.DNSRecord
// @Failure 404 {object} dto.ErrorResponse "DNS record not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /records/{id} [get]
func (handler *DNSHandler) GetDNSRecord(c *gin.Context) {
	id := c.Param("id")

	record, err := handler.repository.GetDNSRecordByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DNS record not found"})
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Summary Update a DNS record
// @Description Update the details of an existing DNS record
// @Tags DNSRecords
// @Accept json
// @Produce json
// @Param id path int true "DNS Record ID"
// @Param dnsRecord body dto.DNSRecordRequest true "DNS Record Update"
// @Success 200 {object} dto.DNSRecord
// @Failure 400 {object} dto.ErrorResponse "Invalid input"
// @Failure 404 {object} dto.ErrorResponse "DNS record not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /records/{id} [put]
func (handler *DNSHandler) UpdateDNSRecord(c *gin.Context) {
	id := c.Param("id")

	var request dto.DNSRecordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create DNS record
	dnsRecord := dto.DNSRecordFromDTO(request)
	record, err := handler.repository.UpdateDNSRecord(id, &dnsRecord)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, record)
}

// @Summary Delete a DNS record
// @Description Delete a DNS record by ID
// @Tags DNSRecords
// @Produce json
// @Param id path int true "DNS Record ID"
// @Success 204 {object} nil
// @Failure 404 {object} dto.ErrorResponse "DNS record not found"
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /records/{id} [delete]
func (handler *DNSHandler) DeleteDNSRecord(c *gin.Context) {
	id := c.Param("id")

	if err := handler.repository.DeleteDNSRecord(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "DNS record not found"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// @Summary Get DNS logs for a record
// @Description Retrieve all logs related to a specific DNS record
// @Tags DNSLogs
// @Produce json
// @Param dnsRecordId path string true "DNS Record ID"
// @Success 200 {array} dto.DNSLog
// @Failure 500 {object} dto.ErrorResponse "Internal server error"
// @Router /logs/{dnsRecordId} [get]
func (handler *DNSHandler) GetDNSLogs(c *gin.Context) {
	dnsRecordID := c.Param("dnsRecordId")

	logs, err := handler.repository.GetDNSLogs(dnsRecordID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
