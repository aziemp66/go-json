package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Homie struct {
	Name      string
	Hobbies   []string
	Addresses []Address
}

type Address struct {
	Street     string
	City       string
	PostalCode string
}

func TestJsonArray(t *testing.T) {
	ledib := Homie{
		Name:    "Ledib",
		Hobbies: []string{"Gaming", "Badminton"},
		Addresses: []Address{
			{
				Street:     "Jl Jl",
				City:       "Karawang",
				PostalCode: "Jepang",
			},
			{
				Street:     "Sn Sn",
				City:       "Jelit",
				PostalCode: "Juning",
			},
		},
	}

	bytes, _ := json.Marshal(ledib)
	fmt.Println(string(bytes))
}
