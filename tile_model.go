package main

import (
	"image"
)

type Tile struct {
	Origin         image.Point `json:"origin"`
	OppositeCorner image.Point `json:"oppositeCorner"`
}

func newTile(originX, originY, oppositeX, oppositeY int) Tile {
	origin := image.Point{originX, originY}
	opposite := image.Point{oppositeX, oppositeY}
	return Tile{origin, opposite}
}

func (this Tile) getDimensions() (x, y int, rect image.Rectangle) {
	x = this.OppositeCorner.X - this.Origin.X
	y = this.OppositeCorner.Y - this.Origin.Y

	rect = image.Rect(0, 0, x, y)
	return
}
