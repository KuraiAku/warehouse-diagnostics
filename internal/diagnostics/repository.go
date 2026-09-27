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

func getOrderDetails(orderID int) ([]OrderLineDetails, error) { // this function goes through the database and returns the details of a all the items in that customers order.
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
			ol.Quantity,
			ol.PickedQuantity,
			ol.Quantity - ol.PickedQuantity AS RemainingToPick,
			h.QuantityOnHand
		FROM Sales.Orders o
		JOIN Sales.OrderLines ol
			ON o.OrderID = ol.OrderID
		JOIN Warehouse.StockItems si
			ON ol.StockItemID = si.StockItemID
		JOIN Warehouse.StockItemHoldings h
			ON ol.StockItemID = h.StockItemID
		WHERE o.OrderID = @p1
		ORDER BY o.OrderDate ASC, o.OrderID ASC;
	`, orderID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []OrderLineDetails{}

	for rows.Next() {
		var item OrderLineDetails

		if err := rows.Scan(
			&item.OrderID,
			&item.OrderDate,
			&item.StockItemName,
			&item.OrderedQuantity,
			&item.PickedQuantity,
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

func getRiskSummary() (RiskSummary, error) { // This function gives us a summary of all the items that are at risk.
	db, err := database.OpenFromEnv("WWI_DB_URL")
	if err != nil {
		return RiskSummary{}, err
	}
	defer db.Close()

	row := db.QueryRow(`
		SELECT
			COUNT(DISTINCT o.OrderID),
			COUNT(*),
			COUNT(DISTINCT si.StockItemID),
			SUM(ol.Quantity - ol.PickedQuantity),
			CONVERT(varchar(10), MIN(o.OrderDate), 23)
		FROM Sales.Orders o
		JOIN Sales.OrderLines ol
			ON o.OrderID = ol.OrderID
		JOIN Warehouse.StockItems si
			ON ol.StockItemID = si.StockItemID
		JOIN Warehouse.StockItemHoldings h
			ON ol.StockItemID = h.StockItemID
		WHERE ol.PickedQuantity < ol.Quantity
			AND (ol.Quantity - ol.PickedQuantity) > h.QuantityOnHand;
	`)

	var summary RiskSummary

	if err := row.Scan(
		&summary.RiskyOrders,
		&summary.RiskyLines,
		&summary.AffectedStockItems,
		&summary.TotalRemaining,
		&summary.OldestOrderDate,
	); err != nil {
		return RiskSummary{}, err
	}

	return summary, nil
}

func getItemRiskSummary() ([]ItemRiskSummary, error) {
	db, err := database.OpenFromEnv("WWI_DB_URL")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
	SELECT
		si.StockItemName,
		COUNT(DISTINCT o.OrderID),
		COUNT(*) AS RiskyLines,
		SUM(ol.Quantity - ol.PickedQuantity) AS TotalRemaining,
		h.QuantityOnHand
	FROM
		Sales.Orders o
	JOIN Sales.OrderLines ol
		ON o.OrderID = ol.OrderID
	JOIN Warehouse.StockItems si
		ON ol.StockItemID = si.StockItemID
	JOIN Warehouse.StockItemHoldings h
		ON si.StockItemID = h.StockItemID
	WHERE (ol.PickedQuantity < ol.Quantity)
		AND	(ol.Quantity - ol.PickedQuantity) > h.QuantityOnHand
	GROUP BY si.StockItemID, si.StockItemName, h.QuantityOnHand
	ORDER BY TotalRemaining DESC;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []ItemRiskSummary{}

	for rows.Next() {
		var item ItemRiskSummary

		if err := rows.Scan(
			&item.StockItemName,
			&item.RiskyOrders,
			&item.RiskyLines,
			&item.TotalRemaining,
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
