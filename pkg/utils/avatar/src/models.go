package src

import "image/color"

type LabelConfiguration struct {
	Text      string
	Font      string
	FontSize  float64
	YPosition int
	Color     color.Color
}
