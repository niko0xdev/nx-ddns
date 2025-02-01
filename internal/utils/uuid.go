package utils

import (
	"log"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func InitSnowflake(nodeID int64) {
	var err error
	node, err = snowflake.NewNode(nodeID)
	if err != nil {
		log.Fatalf("Failed to initialize Snowflake node: %v", err)
	}
}

func GenerateID() snowflake.ID {
	return node.Generate()
}
