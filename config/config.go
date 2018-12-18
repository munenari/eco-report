package config

import (
	"github.com/BurntSushi/toml"
)

// EcoReport config struct
type EcoReport struct {
	Origin          string `toml:"eco-origin"`
	Password        string `toml:"password"`
	FilterValue     string `toml:"filter-value"`
	ElasticOrigin   string `toml:"elastic"`
	BatteryDeviceID string `toml:"battery-deviceid"`
	loadedFilepath  string `toml:"-"`
}

// LoadByPath eco report config file
func LoadByPath(p string) (*EcoReport, error) {
	c := new(EcoReport)
	err := Load(p, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Load toml config into v
func Load(p string, v *EcoReport) error {
	_, err := toml.DecodeFile(p, v)
	v.loadedFilepath = p
	return err
}

// Reload config file by a previous loaded filepath
func (c *EcoReport) Reload() error {
	return Load(c.loadedFilepath, c)
}
