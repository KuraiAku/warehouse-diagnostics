package diagnostics

func PickingBacklog() ([]PickingBacklogItem, error) {
	return getPickingBacklog()
}

func InventoryRisk() ([]PickingBacklogItem, error) {
	return getInventoryRisk()

}
