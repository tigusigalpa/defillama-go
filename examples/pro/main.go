// Pro API usage of the DefiLlama Go SDK.
//
// This example performs real network calls and requires a Pro key:
//
//	DEFILLAMA_API_KEY=... go run ./examples/pro
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	defillama "github.com/tigusigalpa/defillama-go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	key := os.Getenv("DEFILLAMA_API_KEY")
	if key == "" {
		return fmt.Errorf("set DEFILLAMA_API_KEY before running this example")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := defillama.New(
		defillama.WithAPIKey(key),
		defillama.WithTimeout(10*time.Second),
		defillama.WithRetryPolicy(defillama.RetryPolicy{
			MaxAttempts: 3,
			BaseDelay:   250 * time.Millisecond,
			MaxDelay:    2 * time.Second,
			Jitter:      true,
		}),
	)
	if err != nil {
		return fmt.Errorf("create client: %w", err)
	}

	usage, err := client.Account().GetUsage(ctx)
	if err != nil {
		var rateLimit *defillama.RateLimitError
		if errors.As(err, &rateLimit) {
			return fmt.Errorf("usage request was rate limited; retry after %s: %w", rateLimit.RetryAfter, err)
		}
		return fmt.Errorf("get API usage: %w", err)
	}
	if credits, ok := usage.CreditsLeft(); ok {
		fmt.Printf("credits remaining: %.0f\n", credits)
	} else {
		fmt.Printf("usage response: %v\n", usage.Raw)
	}

	chart, err := client.TVL().GetProtocolTVLChart(ctx, "aave", nil)
	if err != nil {
		return fmt.Errorf("get Aave TVL chart: %w", err)
	}
	fmt.Printf("Aave TVL chart points: %d\n", len(chart))
	return nil
}
