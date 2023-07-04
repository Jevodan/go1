package auto

var sizeI = sizesInch{}

type sizesInch struct {
	length float64
	width  float64
	height float64
}

func NewSizeInch(sizes []float64) sizesInch {
	sizeI.length = sizes[0]
	sizeI.width = sizes[1]
	sizeI.height = sizes[2]
	return sizeI
}

func (size sizesInch) Length() Unit {
	unit := Unit{Value: size.length, T: Inch}
	return unit
}

func (size sizesInch) Width() Unit {
	unit := Unit{Value: size.width, T: Inch}
	return unit
}

func (size sizesInch) Height() Unit {
	unit := Unit{Value: size.height, T: Inch}
	return unit
}
