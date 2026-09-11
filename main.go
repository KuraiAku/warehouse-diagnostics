package main

import (
	"fmt"
	"os"
	"warehouse-diagnostics/internal/inventory"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command")
		return
	}

	command := os.Args[1]

	switch command {

	case "inventory":
		if len(os.Args) < 3 {
			fmt.Println("Please add SKU")
			return
		}
		sku := os.Args[2]
		result, err := inventorySummary(sku)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(result)

	case "orders-summary":
		if len(os.Args) < 3 {
			fmt.Println("Please add status")
			return
		}
		status := os.Args[2]
		result, err := ordersSummary(status)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(result)

	case "inventory-all":
		items := inventory.InventoryAll()

		for _, item := range items {

			fmt.Println(item.SKU, item.Quantity)

		}

	default:
		fmt.Println("Unknown command:", command)
	}
}

func ordersSummary(status string) (string, error) {
	if status == "" {
		return "", fmt.Errorf("Please add status")
	}
	return "order-summary " + status, nil
}

func inventorySummary(sku string) (string, error) {

	if sku == "" {
		return "", fmt.Errorf("SKU cannot be empty")
	}

	item, err := inventory.FindInventory(sku)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("inventory %s quantity %d", item.SKU, item.Quantity), nil
}
