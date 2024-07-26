package inventory_benchmark

const SpriteWidth = 150
const SpriteHeight = 150

type Material uint

const (
	unknown Material = iota
	wood
	steel
	stone
	leather
)
