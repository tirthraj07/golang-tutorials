package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
JSON
Type : 		Text, self-describing
Schema: 	Optional (none built-in)
Size: 		Moderate
Speed: 		Moderate
Readable: 	Yes
*/

type Item struct {
	SKU      string
	Quantity int32
	Price    float64
}

type Order struct {
	ID         int64     `json:"id"`
	CustomerID string    `json:"customer_id"`
	Items      []Item    `json:"items"`
	Total      float64   `json:"total"`
	CreatedAt  time.Time `json:"created_at"`
	Status     string    `json:"status"`
	// Never include this field
	Internal string `json:"-"` // never serialized
	// Skip field if zero/empty
	MaybeEmpty string `json:"empty,omitempty"`
}

func main() {
	order := Order{
		ID:         42,
		CustomerID: "cust-007",
		Items: []Item{
			{
				SKU:      "SKU-1",
				Quantity: 2,
				Price:    9.99,
			},
		},
		Total:     19.98,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	// Serialize (Marshal)
	data, err := json.Marshal(order)
	if err != nil {
		fmt.Println("Error while serializing order ", err)
		return
	}

	fmt.Printf("JSON: %s\n", data)

	// Pretty-print for debugging
	pretty, _ := json.MarshalIndent(order, "", "  ")
	fmt.Printf("Pretty:\n%s\n", pretty)

	// Deserialize (Unmarshal)
	var decoded Order
	if err := json.Unmarshal(data, &decoded); err != nil {
		panic(err)
	}
	fmt.Printf("Decoded ID: %d\n", decoded.ID)

}
