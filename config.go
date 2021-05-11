package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Actions []Program `yaml:"Actions"`
}

func ParseConfig(c []byte) *Config {
	color.Cyan("Now parsing configuration information..")
	cfg := &Config{}
	err := yaml.Unmarshal(c, &cfg)
	if err != nil {
		color.Red("Error parsing configuration information, please try again latter.")
		fmt.Println(err)
		os.Exit(1)
	}

	color.Cyan("Succesfully parsed configuration information.")
	color.Cyan("We will install/run the following: ")

	for _, a := range cfg.Actions {
		a.parseArgs()
		fmt.Printf("---- Name: %s | Path: %s | Type: %s ----\n", a.Name, a.Path, a.Kind)
		if len(a.Ressources) > 0 {
			fmt.Printf("---- Ressources: %s ----\n", a.Ressources)
		}
		if len(a.Args) > 0 {
			fmt.Printf("---- Args: %s ----\n", strings.Join(a.Args, " "))
		}

		fmt.Println("-------------------------------------------------------------")
	}

	color.Cyan("Configuration successfully parsed.")

	return cfg
}
