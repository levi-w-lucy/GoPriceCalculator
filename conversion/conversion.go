package conversion

import (
	"fmt"
	"strconv"
)

func StringsToFloats(input []string) ([]float64, error) {
	floatVals := make([]float64, len(input))

	for stringIndex, str := range input {
		float, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, fmt.Errorf("failed to convert %s to float", str)
		}

		floatVals[stringIndex] = float
	}
	return floatVals, nil
}
