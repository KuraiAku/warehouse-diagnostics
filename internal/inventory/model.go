package inventory

type Inventory struct {
	WarehouseCode string
	SKU      string
	Quantity int
}


type OrderSummary struct {

	Status	string
	OrderCount	int
	TotalAmount float64
	
}
