package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v2"
)

var Config struct {
	Env struct {
		
	}

	Client struct {
		Name string `yaml:"username"`
	}
}

func loadConfig() {
	// check for file in .config
	// if not return

	read, err := os.ReadFile("/home/popi/.config/svchat/config.yml")
	if os.IsNotExist(err) {
		// fmt.Println("err: ", err.Error())
		// loadDefaultConfig()
	} else {
		fmt.Println(string(read))
		err := yaml.Unmarshal(read, &Config)
		chk(err)
		fmt.Println(Config)
	}
}
