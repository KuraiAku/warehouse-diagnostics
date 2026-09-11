package main

import (
	"fmt"
	"os"
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
		items := inventoryAll()

		for _, item := range items {

			fmt.Println(item.SKU, item.Quantity)

		}

	default:
		fmt.Println("Unknown command:", command)
	}
}

type Inventory struct {
	SKU      string
	Quantity int
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

	item, err := findInventory(sku)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("inventory %s quantity %d", item.SKU, item.Quantity), nil
}

func inventoryAll() []Inventory {

	items := []Inventory{
		{
			SKU:      "BOX-1001",
			Quantity: 100,
		},

		{
			SKU:      "PAL-2001",
			Quantity: 45,
		},
	}

	return items
}

func findInventory(sku string) (Inventory, error) {

	items := inventoryAll()

	for _, item := range items {
		if item.SKU == sku {
			return item, nil
		}
	}

	return Inventory{}, fmt.Errorf("SKU %s does not exist", sku)
}
