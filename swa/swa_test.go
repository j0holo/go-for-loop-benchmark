package swa

import "testing"

func BenchmarkAllocation(b *testing.B) {
	var inv *inventory
	for i := 0; i < b.N; i++ {
		inv = &inventory{}
	}
	if inv == nil {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItemsRangeLoopValue(b *testing.B) {
	inventory := &inventory{}

	inventory.weight[0] = 10
	inventory.weight[2] = 33
	inventory.weight[5] = 44
	inventory.weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = totalWeightOfIndividualItems_RangeLoopValue(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItemsRangeLoopIndex(b *testing.B) {
	inventory := &inventory{}

	inventory.weight[0] = 10
	inventory.weight[2] = 33
	inventory.weight[5] = 44
	inventory.weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = totalWeightOfIndividualItems_RangeLoopIndex(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItems(b *testing.B) {
	inventory := &inventory{}

	inventory.weight[0] = 10
	inventory.weight[2] = 33
	inventory.weight[5] = 44
	inventory.weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = totalWeightOfIndividualItems(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfInventory(b *testing.B) {
	inventory := &inventory{}

	inventory.weight[0] = 10
	inventory.weight[2] = 33
	inventory.weight[5] = 44
	inventory.weight[9] = 500

	inventory.amount[0] = 1
	inventory.amount[2] = 1
	inventory.amount[5] = 1
	inventory.amount[9] = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = totalWeightOfInventory(inventory)
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}
