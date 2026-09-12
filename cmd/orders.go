package cmd

import (
	"fmt"
	"os"
)

func handleOrdersSummary() {
	if len(os.Args) < 3 {
		fmt.Println("Please add status")
		return
	}
	status := os.Args[2]

	fmt.Println("order-summary " + status)
}
