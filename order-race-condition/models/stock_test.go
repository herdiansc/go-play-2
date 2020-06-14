package models

import "testing"

func TestStock_Order(t *testing.T) {
	stock := NewStock(2)
	currentStock, status, _ := stock.Order()
	if currentStock != stock.Total-1 {
		t.Errorf("Expected: %v, Got: %v\n", stock.Total-1, currentStock)
	}
	if status != "success" {
		t.Errorf("Expected: %v, Got: %v\n", "success", status)
	}
}
