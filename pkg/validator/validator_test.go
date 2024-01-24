package validator

import (
	"fmt"
	"testing"
)

func TestValidatorv10(t *testing.T) {

	type Person struct {
		ID   int64  `json:"id"`
		Name string `json:"name" validate:"requiredIfEmpty=ID"`
	}

	person := Person{
		ID:   1,
		Name: "",
	}

	errMap := ValidateStruct(person)
	fmt.Println(errMap)
}
