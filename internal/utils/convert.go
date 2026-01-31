package utils

func Map[Tin any, Tout any](slice []Tin, fnMap func(int, Tin) Tout) []Tout {
	newSlice := make([]Tout, len(slice))
	for i, v := range slice {
		o := fnMap(i, v)
		newSlice[i] = o
	}
	return newSlice
}

func SliceAs[Tin any, Tout any](slice []Tin) []Tout {
	return Map(slice, func(_ int, v Tin) Tout {
		var a any = v
		return a.(Tout)
	})
}
