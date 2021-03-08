package errs

import (
	"github.com/fatih/camelcase"
)

func transformField(field string) string {
	hardTransforms := map[string]string{"id": "ID"}
	value, ok := hardTransforms[field]
	if ok {
		return value
	}

	fieldForMsg := ""

	fieldWords := camelcase.Split(field)
	next := ""

	for i, w := range fieldWords {
		next = w
		if i != len(fieldWords)-1 {
			next += " "
		}

		fieldForMsg += next
	}
	return fieldForMsg
}
