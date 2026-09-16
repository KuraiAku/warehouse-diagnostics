package inventory

import (
	"fmt"
)

func InventoryAll() ([]Inventory, error) {
	return getAllInventory()
}
func FindInventory(sku string) (Inventory, error) {

	items, err := InventoryAll()
	if err != nil {
		return Inventory{}, err
	}

	for _, item := range items {
		if item.SKU == sku {
			return item, nil
		}
	}

	return Inventory{}, fmt.Errorf("SKU %s does not exist", sku)
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
