package inventory

import "warehouse-diagnostics/internal/database"

func getAllInventory() ([]Inventory, error) {

	db, err := database.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
        SELECT w.warehouse_code, p.sku, i.quantity
    FROM inventory i
    JOIN warehouses w
        ON i.warehouse_id = w.warehouse_id
    JOIN products p
        ON i.product_id = p.product_id
    ORDER BY w.warehouse_code, p.sku;
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Inventory{}

	for rows.Next() {
		var item Inventory

		if err := rows.Scan(&item.WarehouseCode, &item.SKU, &item.Quantity); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func findInventory(warehouseCode string, sku string) (Inventory, error) {
	db, err := database.Open()
	if err != nil {
		return Inventory{}, err
	}

	defer db.Close()

	var item Inventory

	row := db.QueryRow(`
		        SELECT w.warehouse_code, p.sku, i.quantity
		    FROM inventory i
		    JOIN warehouses w
		        ON i.warehouse_id = w.warehouse_id
		    JOIN products p
		        ON i.product_id = p.product_id
		    WHERE w.warehouse_code = @p1 AND p.sku = @p2
		   `, warehouseCode, sku)

	if err := row.Scan(&item.WarehouseCode, &item.SKU, &item.Quantity); err != nil {
		return Inventory{}, err
	}
	return item, nil

}
