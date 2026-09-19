package main

import (
	"fmt"

	"warehouse-diagnostics/internal/database"
)

func main() {
	db, err := database.Open()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	fmt.Println("connected")

	for i := 1; i <= 5; i++ {
		sku := fmt.Sprintf("GEN-%06d", i)
		name := fmt.Sprintf("Generated Product %d", i)

		_, err := db.Exec(`
			INSERT INTO products (sku, product_name, category, unit_cost)
			VALUES (@p1, @p2, @p3, @p4);
		`, sku, name, "Generated", 5.00)
		
		if err != nil {
			fmt.Println(err)
			return
		}
		
		fmt.Println("inserted", sku)
	}
}
