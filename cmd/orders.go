package cmd

import (
	"fmt"
	"os"
	"warehouse-diagnostics/internal/inventory"
)

func handleOrdersSummary() {
	if len(os.Args) < 3 {
		fmt.Println("Please add status")
		return
	}

	status := os.Args[2]

	summary, err := inventory.GetOrdersSummary(status)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(summary.Status, summary.OrderCount, summary.TotalAmount)
}
