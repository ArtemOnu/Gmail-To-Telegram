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
	Host     string `json:"Host"`
}

func GetConfig(input any, addres string) error {
	read, err := os.ReadFile(addres)
	if err != nil {
		return err
	}
	err = json.Unmarshal(read, input)
	return err
}
