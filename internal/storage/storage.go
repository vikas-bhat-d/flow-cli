package storage

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func getStatePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".flow", "state.json")
}

func LoadState() (*State, error) {
	path := getStatePath()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return &State{Version: 1}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var state State
	json.Unmarshal(data, &state)

	return &state, nil
}

func SaveState(state *State) error {
	path := getStatePath()

	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	data, _ := json.MarshalIndent(state, "", "  ")

	return os.WriteFile(path, data, 0644)
}
