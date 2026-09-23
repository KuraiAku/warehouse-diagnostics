package diagnostics

func PickingBacklog() ([]PickingBacklogItem, error) {
	return getPickingBacklog()
}

func InventoryRisk() ([]PickingBacklogItem, error) {
	return getInventoryRisk()

}

func OrderDetails(OrderID int) ([]OrderLineDetails, error) {
	return getOrderDetails(OrderID)
	
}

func GetOrderLineDetails(orderID int) ([]OrderLineDetails, error) {
	return getOrderDetails(orderID)
}

