// Package color provides structs relating to RGB colors.

package color

// RGB represents a color with its values normalized between 0 to 1
type RGB struct {
	R float64
	G float64
	B float64
}

func NewNormRGB(r, g, b float64) RGB {
	return RGB{r, g, b}
}

// AverageRGB is used to calculate the average of many RGBs
type AverageRGB struct {
	rgb   RGB
	count uint
}

// Add adds a color for the average
func (argb *AverageRGB) Add(rgb RGB) {
	argb.rgb.R += rgb.R
	argb.rgb.G += rgb.G
	argb.rgb.B += rgb.B
	argb.count++
}

// Average returns a RGB which is the average of all added colors
func (argb AverageRGB) Average() RGB {
	c := float64(argb.count)
	return NewNormRGB(argb.rgb.R / c, argb.rgb.G / c, argb.rgb.B / c)
}