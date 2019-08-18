package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
)

type Config struct {
	Proxy struct {
		Enabled bool   `yaml:"enabled"`
		Host    string `yaml:"host"`
	}
	Search Search
}

type Search struct {
	Url string `yaml:"url"`
}

var DEBUG = true
var config Config

func main() {
	setup()
	yamlFile, err := ioutil.ReadFile("./.config")
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Proxy enabled: ", config.Proxy.Enabled)
	fmt.Println("Search url: ", config.Search.Url)

	if _, err := os.Stat("./.result"); os.IsNotExist(err) {
		keyword := "도시어부"
		search_url := strings.Replace(config.Search.Url, "%s", keyword, 1)
		fmt.Println(search_url)
		res, err := http.Get(search_url)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s\n", string(data))
		fmt.Println("Connected")
	} else {
		result, err := ioutil.ReadFile("./.result")
		if err != nil {
			panic(err)
		}
		content := string(result)

		fmt.Println("Used file")
	}

}

func setup() {
	if runtime.GOOS == "darwin" {
		DEBUG = true
		fmt.Println(runtime.GOOS)
		fmt.Println(runtime.GOARCH)
	} else {
		DEBUG = false
		fmt.Println(runtime.GOOS)
		fmt.Println(runtime.GOARCH)
	}
}
