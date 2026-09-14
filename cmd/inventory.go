package cmd

import (
	"fmt"
	"os"
	"strconv"
	"warehouse-diagnostics/internal/inventory"
)

func handleInventoryAll() {
	items := inventory.InventoryAll()

	for _, item := range items {
		fmt.Println(item.SKU, item.Quantity)
	}
}

func handleInventory() {

	if len(os.Args) < 3 {
		fmt.Println("Please add SKU")
		return
	}
	sku := os.Args[2]
	item, err := inventory.FindInventory(sku)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item.SKU, item.Quantity)

}

func handleLowInventory() {

	if len(os.Args) < 3 {
		fmt.Println("Please add stock quantity")
		return
	}

	threshold, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Threshold must be a number")
		return
	}
	items := inventory.LowInventory(threshold)

	for _, item := range items {
		fmt.Println(item.SKU, item.Quantity)
	}
}
