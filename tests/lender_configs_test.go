package tests

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func ReadJSONFile(path string) (LenderConfig, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		logrus.Fatalf("Failed to read file: %v\n", err)
	}

	var lenderConfig LenderConfig
	_ = json.NewDecoder(bytes.NewBuffer(b)).Decode(&lenderConfig)

	return lenderConfig, nil
}

func TestLenderConfigs(t *testing.T) {
	var path = "1.json"
	var lenderConfig, err = ReadJSONFile(path)
	if err != nil {
		logrus.Fatalf("Failed to read file: %v\n", err)
	}
	logrus.Infof("%#v", lenderConfig)

}
