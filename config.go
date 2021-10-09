package main

import (
	"fmt"
	"os"
	// "gopkg.in/yaml.v2"
	"gopkg.in/ini.v1"
)

var Config struct {
	Env struct {
		LocalColor 				string
		RemoteColor 			string
		BackgroundColor 		string
		InputPointerColor 		string
		InputTextColor 			string
		CustomPointer 			string
		OverrideCustomColors	bool
	}

	Client struct {
		Name string
		CustomColor string
	}
}

func loadConfig() {
	cfg, err := ini.Load("/home/popi/.config/svchat/config.ini")

	if os.IsNotExist(err) {
		loadDefaultConfig()
	} else {
		Config.Client.Name = cfg.Section("client").Key("username").String()
		Config.Client.CustomColor = cfg.Section("client").Key("custom_color").String()

		Config.Env.LocalColor = cfg.Section("env").Key("local_color").String()
		Config.Env.RemoteColor = cfg.Section("env").Key("remote_color").String()
		Config.Env.BackgroundColor = cfg.Section("env").Key("background_color").String()
		Config.Env.InputPointerColor = cfg.Section("env").Key("input_pointer_color").String()
		Config.Env.InputTextColor = cfg.Section("env").Key("input_text_color").String()
		Config.Env.CustomPointer = cfg.Section("env").Key("custom_pointer").String()
		Config.Env.OverrideCustomColors = cfg.Section("env").Key("override_custom_colors").MustBool()

		loadEmptyConfigValues()
		fmt.Println(Config.Env)
	}
}

func loadDefaultConfig(){
	fmt.Println("Loading default config...")

	Config.Client.Name = ""
	Config.Client.CustomColor = ""

	Config.Env.LocalColor = "red"
	Config.Env.RemoteColor = "white"
	Config.Env.BackgroundColor = "black"
	Config.Env.InputPointerColor = "red"
	Config.Env.InputTextColor = "white"
	Config.Env.CustomPointer = "> "
	Config.Env.OverrideCustomColors = false
}

func loadEmptyConfigValues(){
	replaceEmpty(&Config.Client.Name, "")
	replaceEmpty(&Config.Client.CustomColor, "red")
	replaceEmpty(&Config.Env.LocalColor, "red")
	replaceEmpty(&Config.Env.RemoteColor, "white")
	replaceEmpty(&Config.Env.BackgroundColor, "black")
	replaceEmpty(&Config.Env.InputPointerColor, "red")
	replaceEmpty(&Config.Env.InputTextColor, "white")
	replaceEmpty(&Config.Env.CustomPointer, "> ")
}

func replaceEmpty(s* string, v string){
	if len(*s) == 0 {
		s = &v
	}
}