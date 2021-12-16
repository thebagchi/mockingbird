package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"mockingbird/common"
	"mockingbird/pathtree"
	"net/http"
	"os"
	"strings"
)

type Endpoint struct {
	Request struct {
		Path   string `json:"path" yaml:"path"`
		Method string `json:"method" yaml:"method"`
	} `json:"request" yaml:"request"`

	Response struct {
		Code    int               `json:"code" yaml:"code"`
		Body    string            `json:"body" yaml:"body"`
		Headers map[string]string `json:"headers" yaml:"headers"`
	} `json:"response" yaml:"response"`
}

type Configuration struct {
	Port      int
	Endpoints []Endpoint `json:"endpoints" yaml:"endpoints"`
}

func LoadConfig(path string, port int) *Configuration {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	configuration := &Configuration{
		Port: port,
	}
	err = yaml.NewDecoder(file).Decode(configuration)
	if err != nil {
		return nil
	}
	return configuration
}

func StartServer(configuration *Configuration) {
	if nil == configuration {
		return
	}
	if len(configuration.Endpoints) <= 0 {
		return
	}

	node := pathtree.New()
	for _, endpoint := range configuration.Endpoints {
		path := endpoint.Request.Path
		method := endpoint.Request.Method
		_ = node.Add(fmt.Sprintf("/%s%s", strings.ToLower(method), path), endpoint)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		body, _ := ioutil.ReadAll(request.Body)
		if nil == body {
			body = []byte("")
		}
		fmt.Println(fmt.Sprintf(
			"Time: %s Method: %s Path: %s Body: %s\n",
			common.LocalTimeString(),
			request.Method,
			request.URL.Path,
			string(body),
		))

		_ = common.AppendStringToFile("mock-server.log", fmt.Sprintf(
			"Time: %s Method: %s Path: %s Body: %s\n",
			request.Method,
			request.URL.Path,
			string(body),
		))

		leaf, _ := node.Find(fmt.Sprintf(fmt.Sprintf("/%s%s", strings.ToLower(request.Method), request.URL.Path)))
		if nil == leaf {
			fmt.Println("No path found for:", request.Method, request.URL.Path)
			writer.WriteHeader(http.StatusBadRequest)
		} else {
			endpoint, ok := leaf.Value.(Endpoint)
			if !ok {
				fmt.Println("No endpoint registered for:", request.Method, request.URL.Path)
				writer.WriteHeader(http.StatusBadRequest)
			} else {
				for key, value := range endpoint.Response.Headers {
					writer.Header().Add(key, value)
				}
				writer.WriteHeader(endpoint.Response.Code)
				if len(endpoint.Response.Body) > 0 {
					content, err := ioutil.ReadFile(endpoint.Response.Body)
					if nil == err {
						_, _ = writer.Write(content)
					}
				}
			}
		}
	})
	fmt.Println("Starting http server on port: ", configuration.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil))
}

func main() {
	var (
		config = flag.String("config", "config.yaml", "configuration file path, defaults to config.yaml")
		port   = flag.Int("port", 8080, "port for http server, defaults to 8080")
	)
	flag.Parse()
	fmt.Println("Loading configurations from:", *config)
	configuration := LoadConfig(*config, *port)
	if nil == configuration {
		fmt.Println("Bad configurations, exiting")
		os.Exit(-1)
	}
	StartServer(configuration)
}
