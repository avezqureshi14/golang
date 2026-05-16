package payment

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type FraudHTTPClient struct {
	baseURL string
	client  *http.Client
}

func NewFraudHTTPClient(baseURL string) *FraudHTTPClient {
	return &FraudHTTPClient{
		baseURL: baseURL,
		client: &http.Client{
			Timeout: 3 * time.Second,
		},
	}
}

type fraudRequest struct {
	UserID int64   `json:"user_id"`
	Amount float64 `json:"amount"`
}

type fraudResponse struct {
	IsFraud bool `json:"IsFraud"`
}

func (f *FraudHTTPClient) CheckFraud(ctx context.Context, userID int64, amount float64) (*fraudResponse, error) {

	reqBody := fraudRequest{
		UserID: userID,
		Amount: amount,
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	// ✅ FIX 1: ensure proper URL with scheme + safe join
	base, err := url.Parse(f.baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base url: %w", err)
	}

	base.Path = "/fraud/check"
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		base.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http call failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fraud service returned status: %d", resp.StatusCode)
	}

	// ✅ FIX 2: DO NOT use pointer here
	var result fraudResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}

	return &result, nil
}

func (f *FraudHTTPClient) Close() error {
	return nil
}
