package cmd

import (
	"fmt"
	"os"
	"strconv"

	"warehouse-diagnostics/internal/diagnostics"
)

func handlePickingBacklog() {
	items, err := diagnostics.PickingBacklog()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf(
			"%d %s %s %d %d\n",
			item.OrderID,
			item.OrderDate,
			item.StockItemName,
			item.RemainingToPick,
			item.QuantityOnHand,
		)
	}
}

func handleInventoryRisk() {
	items, err := diagnostics.InventoryRisk()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf(
			"%d %s %s %d %d\n",
			item.OrderID,
			item.OrderDate,
			item.StockItemName,
			item.RemainingToPick,
			item.QuantityOnHand,
		)
	}
}

func handleOrderDetails() {
	if len(os.Args) < 3 {
		fmt.Println("Please add OrderID")
		return
	}

	orderID, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("OrderID must be a number")
		return
	}

	items, err := diagnostics.OrderDetails(orderID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {
		fmt.Printf(
			"%d %s %s %d %d %d %d\n",
			item.OrderID,
			item.OrderDate,
			item.StockItemName,
			item.OrderedQuantity,
			item.PickedQuantity,
			item.RemainingToPick,
			item.QuantityOnHand,
		)
	}
}

func handleRiskSummary() {
	summary, err := diagnostics.GetRiskSummary()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf(
		"Risky Orders: %d\nRisky Lines: %d\nAffected Stock Items: %d\nTotal Remaining: %d\nOldest Order Date: %s\n",
		summary.RiskyOrders,
		summary.RiskyLines,
		summary.AffectedStockItems,
		summary.TotalRemaining,
		summary.OldestOrderDate,
	)
}

func handleItemRiskSummary() {
	items, err := diagnostics.GetItemRiskSummary()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range items {

		fmt.Printf(
			"Stock Item Name: %s\nRisky Orders: %d\nRiskyLines: %d\nTotal Remaining: %d\nQuantityOnHand: %d\n",
			item.StockItemName,
			item.RiskyOrders,
			item.RiskyLines,
			item.TotalRemaining,
			item.QuantityOnHand,
		)
	}
}
