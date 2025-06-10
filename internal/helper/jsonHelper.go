package helper

import (
	"encoding/json"
	"fmt"
	"os"
)

// ReadJSONFile
func ReadJSONFile[T any](path string) (T, error) {
	var result T

	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return result, fmt.Errorf("configuration file does not exist: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return result, fmt.Errorf("failed to read configuration file %s: %w", path, err)
	}

	// Check if file is empty
	if len(data) == 0 {
		return result, fmt.Errorf("configuration file is empty: %s", path)
	}

	if err := json.Unmarshal(data, &result); err != nil {
		return result, fmt.Errorf("failed to parse JSON configuration from %s: %w", path, err)
	}

	return result, nil
}
