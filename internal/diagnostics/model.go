package diagnostics

type PickingBacklogItem struct {
	OrderID         int
	OrderDate       string
	StockItemName   string
	RemainingToPick int
	QuantityOnHand  int
}
