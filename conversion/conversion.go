package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(parameterString []string) ([]float64, error) {
	var resultFloats []float64

	for _, parameterStringVal := range parameterString {
		floatPrice, err := strconv.ParseFloat(parameterStringVal, 64)

		if err != nil {
			return nil, errors.New("Error Converting string to float")
		}

		resultFloats = append(resultFloats, floatPrice)
	}

	return resultFloats, nil
}
