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
	case "inventory-all":
		handleInventoryAll()

	case "inventory":
		handleInventory()

	case "orders-summary":
		handleOrdersSummary()

	case "inventory-low":
		handleLowInventory()

	default:
		fmt.Println("Unknown command:", command)
	}
}
