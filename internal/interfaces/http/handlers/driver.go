package handlers

import (
	"context"
	"encoding/json"
	"net/http"
)

type DriverReq struct {
	ApiUrl string
}

type DriverResponse struct {
	DriverData string `json:"driver_data"`
}

func (d *DriverReq) GetDriverData(ctx context.Context) (string, error) {
	var parsed DriverResponse
	req, err := http.NewRequestWithContext(ctx, "GET", d.ApiUrl, nil)
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
	return parsed.DriverData, nil
}
