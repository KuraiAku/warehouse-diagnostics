package cmd

import (
	"fmt"
	"os"
	"strconv"
	"warehouse-diagnostics/internal/inventory"
)

func handleInventoryAll() {
	items, err := inventory.InventoryAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range items {
		fmt.Println(item.WarehouseCode, item.SKU, item.Quantity)
	}
}

func handleInventory() {

	if len(os.Args) < 4 {
		fmt.Println("Please add warehouse code and SKU")
		return
	}

	warehouseCode := os.Args[2]
	sku := os.Args[3]
	item, err := inventory.FindInventory(warehouseCode, sku)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(item.WarehouseCode, item.SKU, item.Quantity)

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

	items, err := inventory.LowInventory(threshold)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Println(item.WarehouseCode, item.SKU, item.Quantity)
	}
}
