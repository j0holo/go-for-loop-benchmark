package swa

import "inventory_benchmark"

// Struct with arrays SWA

type inventory struct {
	// 8 bytes * 15 = 120 bytes
	weight [15]int
	// 8 bytes * 15 = 120 bytes
	amount              [15]int
	material            [15]inventory_benchmark.Material
	damage              [15]int
	canBeUsedInMainHand [15]bool
	canBeUsedInOffHand  [15]bool
	name                [15]string
	description         [15]string
	texture             [15][inventory_benchmark.SpriteWidth * inventory_benchmark.SpriteHeight]byte
}

func totalWeightOfIndividualItems(i *inventory) int {
	var amount int

	for j := 0; j < len(i.weight); j++ {
		amount += i.weight[j]
	}

	return amount
}

func totalWeightOfIndividualItems_RangeLoopValue(i *inventory) int {
	var amount int

	for _, w := range i.weight {
		amount += w
	}

	return amount
}

func totalWeightOfIndividualItems_RangeLoopIndex(i *inventory) int {
	var amount int

	for j, _ := range i.weight {
		amount += i.weight[j]
	}

	return amount
}

func totalWeightOfInventory(i *inventory) int {
	var amount int

	for j := 0; j < len(i.weight); j++ {
		amount += i.weight[j] * i.amount[j]
	}

	return amount
}
