package config

import (
	"os"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
)

type Address struct {
	Ip   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Config struct {
	Address Address `yaml:"address"`
	LogDir  string  `yaml:"log"`
	Mysql   string  `yaml:"mysql"`
	Redis   string  `yaml:"redis"`
}

func (c *Config) LoadPath(path string) error {
	return Load(c, path)
}
func Load(c interface{}, path string) error {
	file, e := os.Open(path)
	if e != nil {
		return e
	}
	info, e := file.Stat()
	if e != nil {
		return e
	}
	defer file.Close()
	data := make([]byte, info.Size())
	n, e := file.Read(data)
	if e != nil {
		return e
	}
	if int64(n) < info.Size() {
		return errors.New(fmt.Sprintf("cannot read %v bytes from %v", info.Size(), path))
	}

	e = yaml.Unmarshal([]byte(data),c)
	return e
}