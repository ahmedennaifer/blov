package config

import "github.com/ahmedennaifer/blov/internal/config/gcp"

const SAVE_PATH = ".config/.blov/config.json"

type Configurer interface {
	SetRegionOrLocation(value string) error
	SetProjectOrSubscription(value string) error
	Save() error
	Read() error
	Show() any
}

type ProviderConfig struct {
	Provider string
	Config   Configurer
}

func NewProviderConfig(provider string) *ProviderConfig {
	switch provider {
	case "gcp":
		return &ProviderConfig{
			"gcp",
			&gcp.GoogleCloudConfig{},
		}
	default:
		return &ProviderConfig{}
	}
}
