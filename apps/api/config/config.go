package config

import (
	"context"
	"errors"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

const (
	fPath = "settings.yml"
)

const (
	EnvDevelopment Environment = "development"
	EnvProduction  Environment = "production"
)

// Initialize validator once and reuse it
var validate = validator.New()

func New(ctx context.Context) (*Settings, error) {
	settings := &Settings{
		Environment: EnvDevelopment, // Default to development
	}

	// Try to read settings file; ignore if it does not exist
	cf, err := os.ReadFile(fPath)
	if err == nil {
		if err := yaml.Unmarshal(cf, settings); err != nil {
			return nil, err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		// If there was an error other than file not existing, return it
		return nil, err
	}

	// Override with environment variables
	overrideFromEnv(settings)

	// Validate
	if err := validate.Struct(settings); err != nil {
		return nil, err
	}

	return settings, nil
}
