package main

import (
	"bufio"
	"encoding/base64"
	"crypto/md5"
	"github.com/miekg/dns"
	"encoding/hex"
	"fmt"
	"os"
	"net")

var DNSName string

func resolver(domain string, qtype uint16) []dns.RR {
	//ONLY FOR LINUX
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(domain), qtype)
    m.RecursionDesired = true

    c := &dns.Client{}

    response, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if err != nil {
        return nil
    }

    if response == nil {
        return nil
    }

    return response.Answer
}

func integrityCheck(message string)string{
	hash := md5.New()
    hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

func sending(message string) string{

	if message == "" {
		return "empty message"
	}

	crc := integrityCheck(message)
	startingToken := ".ZG5zZXJjX3N0YXJ0X21lc3NhZ2Ug."

	init_resp := resolver(crc+startingToken+DNSName, dns.TypeA)
	fmt.Println(init_resp)
	
	return ""
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var message string

	for scanner.Scan() {
		line := scanner.Text()
		message += line + "\n"
	}
	message = base64.StdEncoding.EncodeToString([]byte(message))

	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Reading error:", err)
        return
    }

	err := sending(message)
	if err != "" {
		fmt.Println("Sending error:", err)
	}
}

// go build -ldflags="-X 'main.DNSName=my.dns.example.com'" main.go