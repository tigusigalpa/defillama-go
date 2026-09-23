// Basic Free API usage of the DefiLlama Go SDK.
//
// This example performs real network calls — run it manually:
//
//	go run ./examples/basic
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	defillama "github.com/tigusigalpa/defillama-go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := defillama.New(defillama.WithTimeout(10 * time.Second))
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	protocols, err := client.TVL().GetProtocols(ctx)
	if err != nil {
		return fmt.Errorf("get protocols: %w", err)
	}
	if len(protocols) == 0 {
		return fmt.Errorf("get protocols: DefiLlama returned no protocols")
	}
	fmt.Printf("protocols: %d (first: %s)\n", len(protocols), protocols[0].Name)

	prices, err := client.Prices().GetCurrentPrices(ctx, []string{"coingecko:bitcoin"})
	if err != nil {
		return fmt.Errorf("get BTC price: %w", err)
	}
	btc, ok := prices["coingecko:bitcoin"]
	if !ok {
		return fmt.Errorf("get BTC price: response did not include bitcoin")
	}
	fmt.Printf("btc: $%.0f\n", btc.Price)
	return nil
}
