package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

func ensureConfigFolderExists() error {
	configLocation, err := getConfigLocation("/")
	if err != nil {
		return err
	}

	if _, err := os.Stat(configLocation); os.IsNotExist(err) {
		err := os.MkdirAll(configLocation, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}

func readDataFromDisk(filename string) (data []byte, err error, notFound bool) {
	configLocation, err := getConfigLocation(filename)
	if err != nil {
		return []byte{}, err, false
	}

	fileData, err := os.ReadFile(configLocation)
	if err != nil {
		if os.IsNotExist(err) {
			return []byte{}, err, true
		}
		return []byte{}, err, false
	}

	return fileData, nil, false
}

func writeDataToDisk(_data any, filename string) error {
	data, err := json.MarshalIndent(_data, "", "  ")
	if err != nil {
		return err
	}

	configLocation, err := getConfigLocation(filename)
	if err != nil {
		return err
	}

	err = os.WriteFile(configLocation, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigLocation(filename string) (string, error) {
	if filename == "" {
		return "", errors.New("filename cannot be empty")
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configLocation := filepath.Join(homeDir, CONFIG_LOCATION_FOLDER, filename)
	return configLocation, nil
}
