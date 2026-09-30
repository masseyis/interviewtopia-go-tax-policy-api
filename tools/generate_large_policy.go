package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type band struct {
	MinimumIncomeCents int64  `json:"minimumIncomeCents"`
	MaximumIncomeCents *int64 `json:"maximumIncomeCents"`
	RateBasisPoints    int    `json:"rateBasisPoints"`
}

type policy struct {
	PolicyID  string `json:"policyId"`
	Residents int    `json:"residents"`
	Currency  string `json:"currency"`
	Bands     []band `json:"bands"`
}

func main() {
	const count = 1000
	const width = int64(10_000_000)

	bands := make([]band, 0, count)
	for i := 0; i < count; i++ {
		minimum := int64(i) * width
		var maximum *int64
		if i < count-1 {
			value := minimum + width - 1
			maximum = &value
		}
		rate := i * 5
		if rate > 5000 {
			rate = 5000
		}
		bands = append(bands, band{minimum, maximum, rate})
	}

	rand.New(rand.NewSource(20260930)).Shuffle(len(bands), func(i, j int) {
		bands[i], bands[j] = bands[j], bands[i]
	})

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(policy{
		PolicyID:  "interviewtopia-go-2026-household-6-large",
		Residents: 6,
		Currency:  "USD",
		Bands:     bands,
	}); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
