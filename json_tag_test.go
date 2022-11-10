package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string  `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	procuct := Product{
		Id:       "asdao",
		Name:     "Samsodent",
		Price:    12.6,
		ImageURL: "/images/asdao",
	}

	byte, err := json.Marshal(procuct)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(byte))
}
