package gojson

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestStreamEncoder(t *testing.T) {
	writer, _ := os.Create("customer_output.json")
	encoder := json.NewEncoder(writer)

	customer := Customer{
		FirstName:  "Azie",
		MiddleName: "Melza",
		LastName:   "Pratama",
		Age:        19,
		IsMarried:  false,
	}
	encoder.Encode(customer)

	fmt.Println(customer)
}
