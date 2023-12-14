package abstract_factory

import "errors"

type ICombo interface {
	// define tat cac cac phuong thuc chung cho tat ca cac Factory

	GetShoe() shoe
	GetShort() Short
	//ISummerCombo
}

// GetWinterCombo - abstract facetory la day
func GetWinterCombo(season string) (ICombo, error) {
	if season == "winter" {
		return NewWinterCombo(), nil
	}

	if season == "summer" {
		return ConcreteSummerCombo{}, nil
	}

	return nil, errors.New("not found this compo")
}
