package cmd

import (
	"fmt"

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
