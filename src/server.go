package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"errors"
	// "regexp"
	"os/exec"
	"path/filepath"
	"gopkg.in/yaml.v3"
)


type Config struct {
	Server ServerConfig `yaml:"server"`
	Client ClientConfig `yaml:"client"`
}
type ServerConfig struct {
	Port   int    `yaml:"port"`
	Domain string `yaml:"domain"`
}
type ClientConfig struct {
	ProjectDir string `yaml:"projectDir"`
}

func isValidDomain(domain string) bool {
	// create normal check
	// var domainRegex = regexp.MustCompile(`^(?!-)[A-Za-z0-9-]{1,63}(?<!-)(\.[A-Za-z]{2,})+$`)
	if len(domain) < 4 || len(domain) > 253 {
		return false
	}
	return true
	// return domainRegex.MatchString(domain)
}

func buildClient(domain string, projectDir string) error {
	
	if !isValidDomain(domain) {
		return errors.New("Wrong domain!")
	}

	clientSrcDir := filepath.Join(projectDir, "client_src")

	if err := os.Chdir(clientSrcDir); err != nil {
		return err
	}

	cmdTidy := exec.Command("go", "mod", "tidy")
	if _, err := cmdTidy.CombinedOutput(); err != nil {
		return err
	}

	cmdBuild := exec.Command("go", "build", fmt.Sprintf("-ldflags=-X 'main.DNSName=%s'", domain),"-o","dnser_c","client.go")
	if _, err := cmdBuild.CombinedOutput(); err != nil {
		return err
	}

	fmt.Printf("The client was successfully builded.\nPath: %s\n", clientSrcDir+"/dnser_c")
	return nil
}

func main() {
	banner := `    __
.--|  |.-----..-----..-----..----.
|  _  ||     ||__ --||  -__||   _|
|_____||__|__||_____||_____||__| 
 server v1.0

 by: fedoik
 Repo: https://github.com/fedoik/dnser/tree/main
 Try for help: server -h
`

	fmt.Println(banner)

	configPath := flag.String("config", "./test_config.yaml", "Path to the configuration file (required)")
	build := flag.Bool("build", false, "Build the client (optional)")
	serve := flag.Bool("serve", false, "Server start (optional)")
	flag.Parse()

	// Init config
	file, err := os.Open(*configPath)
	if err != nil {
		log.Fatalf("[X]Error opening file: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("[X]Error decoding YAML: %v", err)
	}
	// arg for build client
	// from config and from cli

	if *build {
		err = buildClient(config.Server.Domain, config.Client.ProjectDir)
		if err != nil {
			log.Fatalf("[X] Client build error: %v", err)
		}
	}

	if *serve {

	}

}
