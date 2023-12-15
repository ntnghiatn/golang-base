package abstract_factory

import "errors"

type ICombo interface {
	GetShoe() shoe
	GetShort() Short
}

// GetWinterCombo - abstract facetory la day
func GetWinterCombo(seasion string) (ICombo, error) {
	if seasion == "winter" {
		return NewWinterCombo(), nil
	}

	if seasion == "summer" {
		return NewSummerCombo(), nil
	}

	return nil, errors.New("not found this compo")
}
