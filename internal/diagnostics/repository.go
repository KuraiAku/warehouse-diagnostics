package diagnostics

import "warehouse-diagnostics/internal/database"

func getPickingBacklog() ([]PickingBacklogItem, error) {
	db, err := database.OpenFromEnv("WWI_DB_URL")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			o.OrderID,
			CONVERT(varchar(10), o.OrderDate, 23),
			si.StockItemName,
			ol.Quantity - ol.PickedQuantity AS RemainingToPick,
			h.QuantityOnHand
		FROM Sales.Orders o
		JOIN Sales.OrderLines ol
			ON o.OrderID = ol.OrderID
		JOIN Warehouse.StockItems si
			ON ol.StockItemID = si.StockItemID
		JOIN Warehouse.StockItemHoldings h
			ON ol.StockItemID = h.StockItemID
		WHERE ol.PickedQuantity < ol.Quantity
		ORDER BY o.OrderDate ASC, o.OrderID ASC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []PickingBacklogItem{}

	for rows.Next() {
		var item PickingBacklogItem

		if err := rows.Scan(
			&item.OrderID,
			&item.OrderDate,
			&item.StockItemName,
			&item.RemainingToPick,
			&item.QuantityOnHand,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func getInventoryRisk() ([]PickingBacklogItem, error) {

	db, err := database.OpenFromEnv("WWI_DB_URL")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT
			o.OrderID,
			CONVERT(varchar(10), o.OrderDate, 23),
			si.StockItemName,
			ol.Quantity - ol.PickedQuantity AS RemainingToPick,
			h.QuantityOnHand
		FROM Sales.Orders o
		JOIN Sales.OrderLines ol
			ON o.OrderID = ol.OrderID
		JOIN Warehouse.StockItems si
			ON ol.StockItemID = si.StockItemID
		JOIN Warehouse.StockItemHoldings h
			ON ol.StockItemID = h.StockItemID
		WHERE ol.PickedQuantity < ol.Quantity
			AND (ol.Quantity - ol.PickedQuantity) > h.QuantityOnHand
		ORDER BY o.OrderDate ASC, o.OrderID ASC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []PickingBacklogItem{}

	for rows.Next() {
		var item PickingBacklogItem

		if err := rows.Scan(
			&item.OrderID,
			&item.OrderDate,
			&item.StockItemName,
			&item.RemainingToPick,
			&item.QuantityOnHand,
		); err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
