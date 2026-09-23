package inventory

func InventoryAll() ([]Inventory, error) {
	return getAllInventory()
}

func FindInventory(warehouseCode, sku string) (Inventory, error) {
	return findInventory(warehouseCode, sku)
}

func LowInventory(threshold int) ([]Inventory, error) {
	return findLowInventory(threshold)

}

func GetOrdersSummary(status string) (OrderSummary, error) {
	return OrdersSummary(status)
}

