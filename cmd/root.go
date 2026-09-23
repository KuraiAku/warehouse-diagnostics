package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {
	case "picking-backlog":
		handlePickingBacklog()
	case "inventory-risk":
		handleInventoryRisk()
	case "order-details":
		handleOrderDetails()
	default:
		fmt.Println("Unknown command:", command)
	}
}
