package data

import (
	"testing"
)

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name: "test",
		Price: 100,
		SKU: "skfjsd-sdfkj-sdfjk",
	}

	err := p.Validate()

	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
	}

}