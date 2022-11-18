package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	server   Server   `yaml:"server"`
	database Database `yaml:"database"`
}

type Server struct {
	port       int    `yaml:"port"`
	runmode    string `yaml:"runmode"`
	debuglevel string `yaml:"debuglevel"`
}

type Database struct {
	dbtype   string `yaml:"type"`
	host     string `yaml:"host"`
	port     int    `yaml:"port"`
	username string `yaml:"username"`
	password string `yaml:"password"`
	dbname   string `yaml:"dbname"`
}

// tostring methods
// func (s Server) String() string {return ""}

func InitYaml(filepath string) int {
	// resolve relative path problem
	_, pre, _, _ := runtime.Caller(0)
	filepath = path.Join(path.Dir(pre), filepath)
	_, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("file no exist.")
	}
	yamlfile, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Println("err read file")
	}
	c := new(Conf)
	err = yaml.Unmarshal(yamlfile, &c)
	return 0
}
