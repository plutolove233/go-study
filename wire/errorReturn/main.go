package main

import (
	"errors"
	"fmt"
)

type Config struct {
	RemoteAddr string
}

type APIClient struct {
	Config
}

func NewAPIClient(c Config) (*APIClient, error) {
	if c.RemoteAddr == "" {
		return nil, errors.New("no remote addr")
	}
	return &APIClient{Config: c}, nil
}

type Service struct {
	client *APIClient
}

func NewService(client *APIClient) *Service {
	return &Service{client: client}
}

func main() {
	fmt.Println(InitClient(Config{RemoteAddr: ""}))
}
