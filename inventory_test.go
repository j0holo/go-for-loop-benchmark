package inventory_benchmark

import "testing"

const inventorySize = 15

func BenchmarkOOPAllocation(b *testing.B) {
	var inventory *OOPInventory
	for i := 0; i < b.N; i++ {
		inventory = NewOOPInventory(inventorySize)
	}
	if inventory == nil {
		b.Fail()
	}
}

func BenchmarkAOSAllocation(b *testing.B) {
	var inventory *AOSInventory
	for i := 0; i < b.N; i++ {
		inventory = &AOSInventory{}
	}
	if inventory == nil {
		b.Fail()
	}
}

func BenchmarkSWAAllocation(b *testing.B) {
	var inventory *SWAInventory
	for i := 0; i < b.N; i++ {
		inventory = &SWAInventory{}
	}
	if inventory == nil {
		b.Fail()
	}
}

func BenchmarkOOPInventory_TotalWeightOfIndividualItems_RangeLoopValue(b *testing.B) {
	inventory := NewOOPInventory(inventorySize)

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.TotalWeightOfIndividualItems_RangeLoopValue()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkOOPInventory_TotalWeightOfIndividualItems_RangeLoopIndex(b *testing.B) {
	inventory := NewOOPInventory(inventorySize)

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.TotalWeightOfIndividualItems_RangeLoopIndex()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkOOPInventory_TotalWeightOfIndividualItems(b *testing.B) {
	inventory := NewOOPInventory(inventorySize)

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.TotalWeightOfIndividualItems()
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkAOS_TotalWeightOfIndividualItems_RangeLoopValue(b *testing.B) {
	inventory := &AOSInventory{}

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = AOS_TotalWeightOfIndividualItems_RangeLoopValue(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkAOS_TotalWeightOfIndividualItems_RangeLoopIndex(b *testing.B) {
	inventory := &AOSInventory{}

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = AOS_TotalWeightOfIndividualItems_RangeLoopIndex(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkAOS_TotalWeightOfIndividualItems(b *testing.B) {
	inventory := &AOSInventory{}

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = AOS_TotalWeightOfIndividualItems(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkSWA_TotalWeightOfIndividualItems_RangeLoopValue(b *testing.B) {
	inventory := &SWAInventory{}

	inventory.Weight[0] = 10
	inventory.Weight[2] = 33
	inventory.Weight[5] = 44
	inventory.Weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = SWA_TotalWeightOfIndividualItems_RangeLoopValue(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkSWA_TotalWeightOfIndividualItems_RangeLoopIndex(b *testing.B) {
	inventory := &SWAInventory{}

	inventory.Weight[0] = 10
	inventory.Weight[2] = 33
	inventory.Weight[5] = 44
	inventory.Weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = SWA_TotalWeightOfIndividualItems_RangeLoopIndex(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkSWA_TotalWeightOfIndividualItems(b *testing.B) {
	inventory := &SWAInventory{}

	inventory.Weight[0] = 10
	inventory.Weight[2] = 33
	inventory.Weight[5] = 44
	inventory.Weight[9] = 500

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = SWA_TotalWeightOfIndividualItems(inventory)
	}
	if totalWeight != 10+33+44+500 {
		b.Fail()
	}
}

func BenchmarkOOPInventory_TotalWeightOfInventory(b *testing.B) {
	inventory := NewOOPInventory(inventorySize)

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	inventory.slots[0].Amount = 1
	inventory.slots[2].Amount = 1
	inventory.slots[5].Amount = 1
	inventory.slots[9].Amount = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.TotalWeightOfInventory()
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}

func BenchmarkOOPInventory_TotalWeightOfInventory_ExtraCall(b *testing.B) {
	inventory := NewOOPInventory(inventorySize)

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	inventory.slots[0].Amount = 1
	inventory.slots[2].Amount = 1
	inventory.slots[5].Amount = 1
	inventory.slots[9].Amount = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = inventory.TotalWeightOfInventory_ExtraCall()
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}

func BenchmarkAOS_TotalWeightOfInventory(b *testing.B) {
	inventory := &AOSInventory{}

	inventory.slots[0].Weight = 10
	inventory.slots[2].Weight = 33
	inventory.slots[5].Weight = 44
	inventory.slots[9].Weight = 500

	inventory.slots[0].Amount = 1
	inventory.slots[2].Amount = 1
	inventory.slots[5].Amount = 1
	inventory.slots[9].Amount = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = AOS_TotalWeightOfInventory(inventory)
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}

func BenchmarkSWA_TotalWeightOfInventory(b *testing.B) {
	inventory := &SWAInventory{}

	inventory.Weight[0] = 10
	inventory.Weight[2] = 33
	inventory.Weight[5] = 44
	inventory.Weight[9] = 500

	inventory.Amount[0] = 1
	inventory.Amount[2] = 1
	inventory.Amount[5] = 1
	inventory.Amount[9] = 2

	totalWeight := 0
	for i := 0; i < b.N; i++ {
		totalWeight = SWA_TotalWeightOfInventory(inventory)
	}
	if totalWeight != 10+33+44+500*2 {
		b.Fail()
	}
}
