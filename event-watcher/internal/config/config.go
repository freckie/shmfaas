package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type WatcherConfig struct {
	Namespace     string        `yaml:"namespace"`
	FilterTargets bool          `yaml:"filterTargets"`
	Targets       []TargetEvent `yaml:"targets"`
}

type TargetEvent struct {
	Reason       string `yaml:"reason"`
	ResourceKind string `yaml:"resourceKind"`
}

func LoadConfig(filename string) (*WatcherConfig, error) {
	cfg := &WatcherConfig{}

	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
