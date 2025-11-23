package gcp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type GoogleCloudConfig struct {
	ProjectId string `json:"project_id"`
	Region    string `json:"region"`
}

func NewGCPConfig() *GoogleCloudConfig {
	return &GoogleCloudConfig{}
}

func (g *GoogleCloudConfig) SetRegionOrLocation(value string) error {
	if value == "" {
		return fmt.Errorf("error: region value cannot be empty")
	}
	g.Region = value
	return nil
}

func (g *GoogleCloudConfig) SetProjectOrSubscription(value string) error {
	if value == "" {
		return fmt.Errorf("error: project-id value cannot be empty")
	}
	g.ProjectId = value
	return nil
}

func (g *GoogleCloudConfig) Save() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configDir := filepath.Join(home, ".config/.blov")
	if err := os.MkdirAll(configDir, 0o755); err != nil {
		return err
	}

	configPath := filepath.Join(configDir, "config.json")
	data, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, data, 0o600)
}

func (g *GoogleCloudConfig) Read() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	configPath := filepath.Join(home, ".config/.blov", "config.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &g)
	return err
}

func (g *GoogleCloudConfig) Show() any {
	return GoogleCloudConfig{
		ProjectId: g.ProjectId,
		Region:    g.Region,
	}
}
