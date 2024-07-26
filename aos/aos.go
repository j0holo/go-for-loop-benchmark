package aos

import "inventory_benchmark"

// Array of structs AOS

type inventory struct {
	slots [15]item
}

// item is 22566 bytes
type item struct {
	weight              int
	amount              int
	canBeUsedInMainHand bool
	canBeUsedInOffHand  bool
	damage              int
	material            inventory_benchmark.Material
	name                string
	description         string
	texture             [inventory_benchmark.SpriteWidth * inventory_benchmark.SpriteHeight]byte
}

func totalWeightOfIndividualItems_RangeLoopValue(i *inventory) int {
	var amount int
	for _, s := range i.slots {
		amount += s.weight
	}
	return amount
}

func totalWeightOfIndividualItems_RangeLoopIndex(i *inventory) int {
	var amount int
	for j, _ := range i.slots {
		amount += i.slots[j].weight
	}
	return amount
}

func totalWeightOfIndividualItems(i *inventory) int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].weight
	}
	return amount
}

func totalWeightOfInventory(i *inventory) int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].weight * i.slots[j].amount
	}
	return amount
}
