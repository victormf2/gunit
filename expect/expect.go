package expect

type Expector struct {
	value any
}

func It(value any) *Expector {
	return &Expector{
		value: value,
	}
}
