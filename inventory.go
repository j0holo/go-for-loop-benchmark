package inventory_benchmark

const spriteWidth = 150
const spriteHeight = 150

type material uint

const (
	unknown material = iota
	wood
	steel
	stone
	leather
)

// Object-Oriented Programming OOP

type OOPInventory struct {
	slots []OOPItem
}

func NewOOPInventory(size int) *OOPInventory {
	return &OOPInventory{
		slots: make([]OOPItem, size),
	}
}

// OOPItem is 22566 bytes
type OOPItem struct {
	Weight              int
	Amount              int
	CanBeUsedInMainHand bool
	CanBeUsedInOffHand  bool
	Damage              int
	Material            material
	Name                string
	Description         string
	Texture             [spriteWidth * spriteHeight]byte
}

func (i *OOPItem) TotalWeight() int {
	return i.Weight * i.Amount
}

func (i *OOPInventory) TotalWeightOfIndividualItems_RangeLoopValue() int {
	var amount int
	for _, s := range i.slots {
		amount += s.Weight
	}
	return amount
}

func (i *OOPInventory) TotalWeightOfIndividualItems_RangeLoopIndex() int {
	var amount int
	for j, _ := range i.slots {
		amount += i.slots[j].Weight
	}
	return amount
}

func (i *OOPInventory) TotalWeightOfIndividualItems() int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].Weight
	}
	return amount
}

func (i *OOPInventory) TotalWeightOfInventory() int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].Weight * i.slots[j].Amount
	}
	return amount
}

func (i *OOPInventory) TotalWeightOfInventory_ExtraCall() int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		// Pprof says this is inlined.
		amount += i.slots[j].TotalWeight()
	}
	return amount
}

// Array of structs AOS

type AOSInventory struct {
	slots [15]AOSItem
}

// AOSItem is 22566 bytes
type AOSItem struct {
	Weight              int
	Amount              int
	CanBeUsedInMainHand bool
	CanBeUsedInOffHand  bool
	Damage              int
	Material            material
	Name                string
	Description         string
	Texture             [spriteWidth * spriteHeight]byte
}

func AOS_TotalWeightOfIndividualItems_RangeLoopValue(i *AOSInventory) int {
	var amount int
	for _, s := range i.slots {
		amount += s.Weight
	}
	return amount
}

func AOS_TotalWeightOfIndividualItems_RangeLoopIndex(i *AOSInventory) int {
	var amount int
	for j, _ := range i.slots {
		amount += i.slots[j].Weight
	}
	return amount
}

func AOS_TotalWeightOfIndividualItems(i *AOSInventory) int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].Weight
	}
	return amount
}

func AOS_TotalWeightOfInventory(i *AOSInventory) int {
	var amount int
	for j := 0; j < len(i.slots); j++ {
		amount += i.slots[j].Weight * i.slots[j].Amount
	}
	return amount
}

// Struct with arrays SWA

type SWAInventory struct {
	// 8 bytes * 15 = 120 bytes
	Weight [15]int
	// 8 bytes * 15 = 120 bytes
	Amount              [15]int
	Material            [15]material
	Damage              [15]int
	CanBeUsedInMainHand [15]bool
	CanBeUsedInOffHand  [15]bool
	Name                [15]string
	Description         [15]string
	Texture             [15][spriteWidth * spriteHeight]byte
}

func SWA_TotalWeightOfIndividualItems(i *SWAInventory) int {
	var amount int

	for j := 0; j < len(i.Weight); j++ {
		amount += i.Weight[j]
	}

	return amount
}

func SWA_TotalWeightOfIndividualItems_RangeLoopValue(i *SWAInventory) int {
	var amount int

	for _, w := range i.Weight {
		amount += w
	}

	return amount
}

func SWA_TotalWeightOfIndividualItems_RangeLoopIndex(i *SWAInventory) int {
	var amount int

	for j, _ := range i.Weight {
		amount += i.Weight[j]
	}

	return amount
}

func SWA_TotalWeightOfInventory(i *SWAInventory) int {
	var amount int

	for j := 0; j < len(i.Weight); j++ {
		amount += i.Weight[j] * i.Amount[j]
	}

	return amount
}
