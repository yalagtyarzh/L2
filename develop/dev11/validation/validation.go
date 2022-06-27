package validation

import (
	"fmt"
	"net/url"
	"strings"
)

// ValidateURLValues валидирует url string с переданным словарей url параметров, проверяя их наличие
func ValidateURLValues(values url.Values, dict ...string) (map[string]string, error) {
	kvalue := make(map[string]string)
	missedValues := make([]string, 0)

	for _, v := range dict {
		value := values.Get(v)
		if value == "" {
			missedValues = append(missedValues, v)
		}

		kvalue[v] = value
	}

	if len(missedValues) != 0 {
		valstring := strings.Join(missedValues, ", ")
		return map[string]string{}, fmt.Errorf("not enough parameters: %s", valstring)
	}

	return kvalue, nil
}
