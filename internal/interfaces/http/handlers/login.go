package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type LoginUs struct {
	ApiUrl string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (l *LoginUs) Login(ctx context.Context, email, password string) (string, error) {
	loginReq := LoginRequest{
		Email:    email,
		Password: password,
	}
	body, err := json.Marshal(loginReq)
	req, err := http.NewRequestWithContext(ctx, "POST", l.ApiUrl, bytes.NewReader(body))
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
	var parsed LoginResponse
	_ = json.NewDecoder(resp.Body).Decode(&parsed)

	return parsed.Token, nil
}
