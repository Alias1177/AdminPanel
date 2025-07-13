package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type ClientReq struct {
	ApiUrl string
}
type ClientResponse struct {
	ClientData string `json:"client_data"`
}

func (c *ClientReq) GetClientData(ctx context.Context) (string, error) {
	var parsed ClientResponse
	req, err := http.NewRequestWithContext(ctx, "GET", c.ApiUrl, nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", err
	}
	_ = json.NewDecoder(resp.Body).Decode(&parsed)
	return parsed.ClientData, nil
}
