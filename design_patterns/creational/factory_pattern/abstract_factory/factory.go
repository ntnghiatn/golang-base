package abstract_factory

import "errors"

// GetWinterCombo - abstract facetory la day
func GetWinterCombo(seasion string) (IwinterCombo, error) {
	if seasion == "winter" {
		return ConcreteWinterCombo{}, nil
	}

	if seasion == "summer" {
		return ConcreteSummerCombo{}, nil
	}

	return nil, errors.New("not found this compo")
}
