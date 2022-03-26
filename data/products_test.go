package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "Coffe",
		Price: 1.00,
		SKU: "abs-ced-frg",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
