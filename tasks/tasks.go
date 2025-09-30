package tasks

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Tasks struct {
	Taskname    string
	Description string
	Priority    string
	Completed   bool
}

func tasksFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".godo")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return filepath.Join(dir, "tasks.json"), nil
}
func SaveTasks(ts []Tasks) error {
	path, err := tasksFilePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(ts, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

func LoadTasks() ([]Tasks, error) {
	path, err := tasksFilePath()
	if err != nil {
		return []Tasks{}, err
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return []Tasks{}, err
	}
	var ts []Tasks
	if err := json.Unmarshal(b, &ts); err != nil {
		return nil, err
	}
	return ts, nil
}
