package diagnostics

func PickingBacklog() ([]PickingBacklogItem, error) {
	return getPickingBacklog()
}

func InventoryRisk() ([]PickingBacklogItem, error) {
	return getInventoryRisk()
}

func OrderDetails(orderID int) ([]OrderLineDetails, error) {
	return getOrderDetails(orderID)
}
