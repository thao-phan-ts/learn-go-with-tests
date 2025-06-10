package utils

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
	"learn-go-with-tests/internal/entity/dj"
)

func ReadJSONFile(path string) (dj.LenderConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatalf("Failed to read file: %v\n", err)
	}

	var lenderConfig dj.LenderConfig
	_ = json.NewDecoder(bytes.NewBuffer(b)).Decode(&lenderConfig)

	return lenderConfig, nil
}
