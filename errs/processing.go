package errs

import (
	"github.com/fatih/camelcase"
)

func splitField(field string) string {
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
