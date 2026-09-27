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

type RiskSummary struct {
	RiskyOrders        int
	RiskyLines         int
	AffectedStockItems int
	TotalRemaining     int
	OldestOrderDate    string
}

type ItemRiskSummary struct {
	StockItemName  string
	RiskyOrders    int
	RiskyLines     int
	TotalRemaining int
	QuantityOnHand int
}
