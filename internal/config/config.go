package config

import (
    "encoding/json"
    "errors"
    "fmt"
    "os"
    "path/filepath"
)

type Config struct {
    BasePath    string `json:"base_path"`
    TemplatePath string `json:"template_path,omitempty"`
    MakefilePath string `json:"makefile_path,omitempty"`
    ConfigPath   string `json:"config_path"`
}

func DefaultPath() string {
    cfgDir, _ := os.UserConfigDir()
    if cfgDir == "" {
        cfgDir = filepath.Join(os.Getenv("HOME"), ".config")
    }
    return filepath.Join(cfgDir, "puc_rio_cp", "config.json")
}

func Load(path string) (Config, error) {
    if path == "" {
        path = DefaultPath()
    }
    cfg := Config{ConfigPath: path}
    b, err := os.ReadFile(path)
    if err != nil {
        if errors.Is(err, os.ErrNotExist) {
            return cfg, nil
        }
        return cfg, err
    }
    if err := json.Unmarshal(b, &cfg); err != nil {
        return cfg, fmt.Errorf("invalid configuration: %w", err)
    }
    if cfg.ConfigPath == "" {
        cfg.ConfigPath = path
    }
    return cfg, nil
}

func Save(cfg Config) error {
    if cfg.ConfigPath == "" {
        cfg.ConfigPath = DefaultPath()
    }
    if err := os.MkdirAll(filepath.Dir(cfg.ConfigPath), 0o755); err != nil {
        return err
    }
    b, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(cfg.ConfigPath, b, 0o644)
}

