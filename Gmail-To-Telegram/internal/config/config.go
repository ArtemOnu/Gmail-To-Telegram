package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Token    string `json:"Token"`
	ChatID   int64  `json:"Chat-ID"`
	Mail     string `json:"Mail-addres"`
	Password string `json:"Password"`
}

func GetConfig(input any) error {
	read, err := os.ReadFile("../../configArtemka.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(read, input)
	return err
}
