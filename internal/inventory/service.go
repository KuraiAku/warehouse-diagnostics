package inventory

func InventoryAll() ([]Inventory, error) {
	return getAllInventory()
}

func FindInventory(warehouseCode, sku string) (Inventory, error) {
	return findInventory(warehouseCode, sku)
}

func LowInventory(threshold int) ([]Inventory, error) {

	results := []Inventory{}

	items, err := InventoryAll()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.Quantity < threshold {
			results = append(results, item)
		}
	}

	return results, nil

}

func GetOrdersSummary(status string) (OrderSummary, error) {
	return OrdersSummary(status)
}
