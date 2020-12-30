package main

type Config struct {
	ConfigFile string `json:"config_file" yaml:"configFile"`
	Data       string `yaml:"data"        json:"data"`
}
