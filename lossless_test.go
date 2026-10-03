package defillama

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

const (
	dl1Hash = "e2645a63ed69051971f8101ad6bd078074063d8456395983e279e22bb0fe683f"
	dl2Hash = "543f250751224544f34cd3e1c1938558cc490fc7dd900a2eab8568e362beac48"
	dl3Hash = "bac03ef57d68fab1cbaa54811bf2009f2a37b9057f59736f890392fcbe30b43c"
	dl4Hash = "3595adf624f091f64509f593705a0d3d63c1515eaf22d134de4782519d153bb0"
)

func TestLosslessProtocolReceipts(t *testing.T) {
	protocolsFixture := losslessFixture(t, "DL1-protocols.json")
	detailsFixture := losslessFixture(t, "DL2-protocol-details.json")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-Fixture", "protocol")
		switch r.URL.Path {
		case "/protocols":
			_, _ = w.Write(protocolsFixture)
		case "/protocol/provider-aave":
			_, _ = w.Write(detailsFixture)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	client, err := New(WithBaseURLsForTesting(map[string]string{
		"https://api.llama.fi": server.URL,
	}))
	if err != nil {
		t.Fatal(err)
	}

	before := time.Now().UTC()
	protocolsReceipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	after := time.Now().UTC()
	if err != nil {
		t.Fatal(err)
	}
	assertReceipt(t, protocolsReceipt, "GET /protocols", server.URL+"/protocols", dl1Hash, protocolsFixture)
	if protocolsReceipt.DocsURL != "https://api-docs.defillama.com/#tag/tvl/get/protocols" {
		t.Errorf("DocsURL = %q", protocolsReceipt.DocsURL)
	}
	if protocolsReceipt.CapturedAt.Before(before) || protocolsReceipt.CapturedAt.After(after) {
		t.Errorf("CapturedAt = %s, outside request interval", protocolsReceipt.CapturedAt)
	}
	if got := protocolsReceipt.Header().Get("X-Fixture"); got != "protocol" {
		t.Errorf("receipt header = %q", got)
	}

	var generic []any
	if err := protocolsReceipt.Decode(&generic); err != nil {
		t.Fatal(err)
	}
	if got := generic[0].(map[string]any)["tvl"].(json.Number).String(); got != "1234567890.12345678901234567890" {
		t.Errorf("UseNumber TVL = %q", got)
	}

	var protocols LosslessProtocols
	if err := protocolsReceipt.Decode(&protocols); err != nil {
		t.Fatal(err)
	}
	entries := protocols.Entries()
	if len(entries) != 1 {
		t.Fatalf("entries = %d, want 1", len(entries))
	}
	if got, present, err := entries[0].ID(); err != nil || !present || got != "provider-aave" {
		t.Errorf("protocol ID = (%q, %v, %v)", got, present, err)
	}
	if got, present, err := entries[0].TVL(); err != nil || !present || got.String() != "1234567890.12345678901234567890" {
		t.Errorf("protocol TVL = (%q, %v, %v)", got, present, err)
	}
	chainTVLs, present := entries[0].Fields().Field("chainTvls")
	if !present {
		t.Fatal("chainTvls is absent")
	}
	chains, err := chainTVLs.Object()
	if err != nil {
		t.Fatal(err)
	}
	ethereum, present := chains.Field("Ethereum")
	if !present {
		t.Fatal("Ethereum chain TVL is absent")
	}
	ethTVL, err := ethereum.Number()
	if err != nil || ethTVL.String() != "987654321.12345678901234567890" {
		t.Errorf("Ethereum TVL = (%q, %v)", ethTVL, err)
	}
	if !entries[0].Fields().Has("unknownProviderField") {
		t.Error("unknown provider field was discarded")
	}

	detailsReceipt, err := client.TVL().GetProtocolReceipt(context.Background(), "provider-aave")
	if err != nil {
		t.Fatal(err)
	}
	assertReceipt(t, detailsReceipt, "GET /protocol/{protocol}", server.URL+"/protocol/provider-aave", dl2Hash, detailsFixture)
	var details LosslessProtocolDetails
	if err := detailsReceipt.Decode(&details); err != nil {
		t.Fatal(err)
	}
	if got, present, err := details.ID(); err != nil || !present || got != "provider-aave" {
		t.Errorf("details ID = (%q, %v, %v)", got, present, err)
	}
	history, present, err := details.TVLHistory()
	if err != nil || !present || len(history) != 1 {
		t.Fatalf("history = (%d entries, %v, %v)", len(history), present, err)
	}
	date, present, err := numberField(history[0], "date")
	if err != nil || !present || date.String() != "1728000000" {
		t.Errorf("history date = (%q, %v, %v)", date, present, err)
	}
	historyTVL, present, err := numberField(history[0], "totalLiquidityUSD")
	if err != nil || !present || historyTVL.String() != "1234567890.12345678901234567890" {
		t.Errorf("history TVL = (%q, %v, %v)", historyTVL, present, err)
	}
	chains, present, err = details.ChainTVLs()
	if err != nil || !present {
		t.Fatalf("detail chainTvls = (%v, %v)", present, err)
	}
	ethereum, present = chains.Field("Ethereum")
	if !present {
		t.Fatal("detail Ethereum chain TVL is absent")
	}
	chainTVL, err := ethereum.Number()
	if err != nil || chainTVL.String() != "987654321.12345678901234567890" {
		t.Errorf("detail Ethereum TVL = (%q, %v)", chainTVL, err)
	}
	if detailsReceipt.CapturedAt.Equal(time.Unix(1728000000, 0).UTC()) {
		t.Error("capture time was conflated with provider history time")
	}
	if !details.Fields().Has("unknownMethodology") {
		t.Error("unknown methodology field was discarded")
	}
}

func TestLosslessYieldReceipts(t *testing.T) {
	poolsFixture := losslessFixture(t, "DL3-yield-pools.json")
	chartFixture := losslessFixture(t, "DL4-yield-chart.json")
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/pools":
			_, _ = w.Write(poolsFixture)
		case "/chart/provider-pool-main":
			_, _ = w.Write(chartFixture)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)

	client, err := New(WithBaseURLsForTesting(map[string]string{
		"https://yields.llama.fi": server.URL,
	}))
	if err != nil {
		t.Fatal(err)
	}

	poolsReceipt, err := client.Yields().GetPoolsReceipt(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	assertReceipt(t, poolsReceipt, "GET /pools", server.URL+"/pools", dl3Hash, poolsFixture)
	var poolsResponse LosslessYieldPools
	if err := poolsReceipt.Decode(&poolsResponse); err != nil {
		t.Fatal(err)
	}
	pools, present, err := poolsResponse.Pools()
	if err != nil || !present || len(pools) != 3 {
		t.Fatalf("pools = (%d entries, %v, %v)", len(pools), present, err)
	}
	if got, present, err := pools[0].PoolID(); err != nil || !present || got != "provider-pool-main" {
		t.Errorf("pool ID = (%q, %v, %v)", got, present, err)
	}
	if got, present, err := pools[0].APY(); err != nil || !present || got.String() != "12.345678901234567890" {
		t.Errorf("APY = (%q, %v, %v)", got, present, err)
	}
	reward, present := pools[0].APYReward()
	if !present || !reward.IsNull() {
		t.Errorf("null apyReward = (present=%v, null=%v)", present, reward.IsNull())
	}
	zeroReward, present := pools[1].APYReward()
	if !present || zeroReward.IsNull() {
		t.Errorf("zero apyReward = (present=%v, null=%v)", present, zeroReward.IsNull())
	}
	if got, err := zeroReward.Number(); err != nil || got.String() != "0" {
		t.Errorf("zero apyReward number = (%q, %v)", got, err)
	}
	if _, present := pools[2].APYReward(); present {
		t.Error("absent apyReward was reported as present")
	}
	rewardTokens, present := pools[0].RewardTokens()
	if !present || !rewardTokens.IsNull() {
		t.Errorf("null rewardTokens = (present=%v, null=%v)", present, rewardTokens.IsNull())
	}
	if _, present := pools[2].RewardTokens(); present {
		t.Error("absent rewardTokens was reported as present")
	}
	predictions, present := pools[0].Fields().Field("predictions")
	if !present {
		t.Fatal("predictions field was discarded")
	}
	predictionObject, err := predictions.Object()
	if err != nil {
		t.Fatal(err)
	}
	probability, present, err := numberField(predictionObject, "predictedProbability")
	if err != nil || !present || probability.String() != "0.98765432109876543210" {
		t.Errorf("predictedProbability = (%q, %v, %v)", probability, present, err)
	}

	chartReceipt, err := client.Yields().GetPoolChartReceipt(context.Background(), "provider-pool-main")
	if err != nil {
		t.Fatal(err)
	}
	assertReceipt(t, chartReceipt, "GET /chart/{pool}", server.URL+"/chart/provider-pool-main", dl4Hash, chartFixture)
	var chart LosslessYieldChart
	if err := chartReceipt.Decode(&chart); err != nil {
		t.Fatal(err)
	}
	points, present, err := chart.Points()
	if err != nil || !present || len(points) != 1 {
		t.Fatalf("points = (%d entries, %v, %v)", len(points), present, err)
	}
	if got, present, err := points[0].Timestamp(); err != nil || !present || got != "2024-01-02T03:04:05.678Z" {
		t.Errorf("timestamp = (%q, %v, %v)", got, present, err)
	}
	if got, present, err := points[0].APY(); err != nil || !present || got.String() != "5.200000000000000001" {
		t.Errorf("chart APY = (%q, %v, %v)", got, present, err)
	}
	chartReward, present := points[0].APYReward()
	if !present || !chartReward.IsNull() {
		t.Errorf("chart null apyReward = (present=%v, null=%v)", present, chartReward.IsNull())
	}
	if !points[0].Fields().Has("providerCorrection") {
		t.Error("chart correction field was discarded")
	}
}

func TestReceiptObserverCapturesEveryRetryResponse(t *testing.T) {
	fixture := losslessFixture(t, "DL1-protocols.json")
	var calls int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		calls++
		if calls == 1 {
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte(`{"error":"retry"}`))
			return
		}
		w.Header().Set("X-Immutable", "original")
		_, _ = w.Write(fixture)
	}))
	t.Cleanup(server.Close)

	var receipts []ResponseReceipt
	client, err := New(
		WithBaseURLsForTesting(map[string]string{"https://api.llama.fi": server.URL}),
		WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
		WithReceiptObserver(func(receipt ResponseReceipt) {
			receipts = append(receipts, receipt)
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(receipts) != 2 {
		t.Fatalf("receipts = %d, want 2", len(receipts))
	}
	if receipts[0].StatusCode != http.StatusServiceUnavailable || receipts[0].Attempt != 1 {
		t.Errorf("first receipt = status %d, attempt %d", receipts[0].StatusCode, receipts[0].Attempt)
	}
	if receipts[1].StatusCode != http.StatusOK || receipts[1].Attempt != 2 || receipts[1].BodySHA256() != dl1Hash {
		t.Errorf("second receipt = status %d, attempt %d, hash %s", receipts[1].StatusCode, receipts[1].Attempt, receipts[1].BodySHA256())
	}
	body := receipts[1].Body()
	body[0] = 'X'
	if bytes.Equal(body, receipts[1].Body()) {
		t.Error("Body exposed mutable receipt storage")
	}
	header := receipts[1].Header()
	header.Set("X-Immutable", "mutated")
	if got := receipts[1].Header().Get("X-Immutable"); got != "original" {
		t.Errorf("Header exposed mutable receipt storage: %q", got)
	}
}

func TestReceiptOptionRejectsNilObserver(t *testing.T) {
	if _, err := New(WithReceiptObserver(nil)); err == nil {
		t.Error("nil receipt observer was accepted")
	}
}

func losslessFixture(t *testing.T, name string) []byte {
	t.Helper()
	body, err := os.ReadFile("testdata/lossless/" + name)
	if err != nil {
		t.Fatal(err)
	}
	return body
}

func assertReceipt(t *testing.T, receipt *ResponseReceipt, route, sourceURL, hash string, fixture []byte) {
	t.Helper()
	if receipt == nil {
		t.Fatal("receipt is nil")
	}
	if receipt.Route != route || receipt.SourceURL != sourceURL {
		t.Errorf("receipt route/URL = %q / %q, want %q / %q", receipt.Route, receipt.SourceURL, route, sourceURL)
	}
	if receipt.StatusCode != http.StatusOK || receipt.Attempt != 1 {
		t.Errorf("receipt status/attempt = %d / %d", receipt.StatusCode, receipt.Attempt)
	}
	if receipt.BodySHA256() != hash {
		t.Errorf("receipt hash = %s, want %s", receipt.BodySHA256(), hash)
	}
	if !bytes.Equal(receipt.Body(), fixture) {
		t.Error("receipt body differs from the immutable fixture bytes")
	}
}
