package diagnostics

type PickingBacklogItem struct {
	OrderID         int
	OrderDate       string
	StockItemName   string
	RemainingToPick int
	QuantityOnHand  int
}

type OrderLineDetails struct {
	OrderID         int
	OrderDate       string
	StockItemName   string
	OrderedQuantity int
	PickedQuantity  int
	RemainingToPick int
	QuantityOnHand  int
}
