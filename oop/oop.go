package oop

import "inventory_benchmark"

// Object-Oriented Programming OOP

type inventory struct {
	slots [15]item
}

func newInventory() *inventory {
	return &inventory{}
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

func (i *item) totalWeight() int {
	return i.weight * i.amount
}

func (i *inventory) totalWeightOfIndividualItems_RangeLoopValue() int {
	var amount int
	for _, s := range i.slots {
		amount += s.weight
	}
	return amount
}

func (i *inventory) totalWeightOfIndividualItems_RangeLoopIndex() int {
	var amount int
	for j, _ := range i.slots {
		amount += i.slots[j].weight
	}
	return amount
}

func (i *inventory) totalWeightOfIndividualItems() int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].weight
	}
	return amount
}

func (i *inventory) totalWeightOfInventory() int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		// Pprof says this is inlined.
		amount += i.slots[j].totalWeight()
	}
	return amount
}
