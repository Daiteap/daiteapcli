package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type IConfig struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func getConfigLocation() (string, error) {
	cfgPath, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(cfgPath, "daiteap"), nil
}

func InitConfig() error {
	daiteapCfgDir, err := getConfigLocation()
	if err != nil {
		return err
	}

	// create daiteap config directory
	if _, err = os.Stat(daiteapCfgDir); os.IsNotExist(err) {
		err = os.MkdirAll(daiteapCfgDir, 0o700)
		if err != nil {
			return err
		}
	}

	return err
}

func SaveConfig(cfg *IConfig) error {
	cfgDir, err := getConfigLocation()
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%s/%s", cfgDir, "config.json")
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}

	if _, err = os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0o600)

	if err != nil {
		return fmt.Errorf("%v: %w", "unable to save config", err)
	}
	return nil
}

func SaveToken(token string) error {
	cfgDir, err := getConfigLocation()
	if err != nil {
		return err
	}

	file := fmt.Sprintf("%s/%s", cfgDir, "config.json")

	var cfg *IConfig = &IConfig{
		AccessToken: token,
	}

	data, err := json.MarshalIndent(&cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}

	if _, err = os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0o600)
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to save config", err)
	}
	return nil
}

func GetToken() (string, error) {
	cfgDir, err := getConfigLocation()
	if err != nil {
		return "", err
	}

	file := fmt.Sprintf("%s/%s", cfgDir, "config.json")

	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("%v: %w", "unable to save config", err)
	}

	var f interface{}
	json.Unmarshal(content, &f)
	m := f.(map[string]interface{})

	var cfg *IConfig = &IConfig{
		AccessToken:  m["access_token"].(string),
		RefreshToken: m["refresh_token"].(string),
	}

	return cfg.AccessToken, nil
}
