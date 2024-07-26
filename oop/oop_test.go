package oop

import "testing"

func BenchmarkAllocation(b *testing.B) {
	var inventory *inventory
	for i := 0; i < b.N; i++ {
		inventory = newInventory()
	}
	if inventory == nil {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItemsRangeLoopValue(b *testing.B) {
	inventory := newInventory()

	inventory.slots[0].weight = 10
	inventory.slots[2].weight = 33
	inventory.slots[5].weight = 44
	inventory.slots[9].weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.totalWeightOfIndividualItems_RangeLoopValue()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItemsRangeLoopIndex(b *testing.B) {
	inventory := newInventory()

	inventory.slots[0].weight = 10
	inventory.slots[2].weight = 33
	inventory.slots[5].weight = 44
	inventory.slots[9].weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.totalWeightOfIndividualItems_RangeLoopIndex()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfIndividualItems(b *testing.B) {
	inventory := newInventory()

	inventory.slots[0].weight = 10
	inventory.slots[2].weight = 33
	inventory.slots[5].weight = 44
	inventory.slots[9].weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.totalWeightOfIndividualItems()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkTotalWeightOfInventory(b *testing.B) {
	inventory := newInventory()

	inventory.slots[0].weight = 10
	inventory.slots[2].weight = 33
	inventory.slots[5].weight = 44
	inventory.slots[9].weight = 500

	inventory.slots[0].amount = 1
	inventory.slots[2].amount = 1
	inventory.slots[5].amount = 1
	inventory.slots[9].amount = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.totalWeightOfInventory()
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}
