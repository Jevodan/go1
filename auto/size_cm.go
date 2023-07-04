package auto

var size = sizesCM{}

type sizesCM struct {
	length float64
	width  float64
	height float64
}

func NewSizeCM(sizes []float64) sizesCM {
	size.length = sizes[0]
	size.width = sizes[1]
	size.height = sizes[2]
	return size
}

func (s sizesCM) Length() Unit {
	unit := Unit{Value: s.length, T: CM}
	return unit
}

func (s sizesCM) Width() Unit {
	unit := Unit{Value: s.width, T: CM}
	return unit
}

func (s sizesCM) Height() Unit {
	unit := Unit{Value: s.height, T: CM}
	return unit
}
