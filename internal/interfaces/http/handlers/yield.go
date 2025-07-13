package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Yields struct {
	ApiURL string
}

func (y *Yields) GetCleanYield(ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", y.ApiURL, nil)
	if err != nil {
		return "", fmt.Errorf("Cant make request on api, error: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("Cant take the body, error: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Cant get the body, status code: %v", resp.StatusCode)
	}
	var parsed balanceResponse
	if err := json.NewDecoder(resp.Body).Decode(&parsed); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}
	clean, _ := strconv.Atoi(parsed.Balance)
	res := clean * 10 / 100
	result := strconv.Itoa(res)
	return result, nil
}
