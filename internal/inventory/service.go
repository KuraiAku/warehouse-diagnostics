package inventory

import "fmt"

func InventoryAll() []Inventory {

	items := []Inventory{
		{
			SKU:      "BOX-1001",
			Quantity: 100,
		},

		{
			SKU:      "PAL-2001",
			Quantity: 45,
		},
	}

	return items
}

func FindInventory(sku string) (Inventory, error) {

	items := InventoryAll()

	for _, item := range items {
		if item.SKU == sku {
			return item, nil
		}
	}

	return Inventory{}, fmt.Errorf("SKU %s does not exist", sku)
}

