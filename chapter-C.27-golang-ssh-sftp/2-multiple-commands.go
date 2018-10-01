package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func main() {
	const SSH_ADDRESS = "0.0.0.0:22"
	const SSH_USERNAME = "user"
	const SSH_KEY = ""
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
			// PublicKeyFile(SSH_KEY),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	session, err := client.NewSession()
	if session != nil {
		defer session.Close()
	}
	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}

	var stdout, stderr bytes.Buffer
	session.Stdout = &stdout
	session.Stderr = &stderr

	stdin, err := session.StdinPipe()
	if err != nil {
		log.Fatal("Error getting stdin pipe. " + err.Error())
	}

	err = session.Start("/bin/bash")
	if err != nil {
		log.Fatal("Error starting bash. " + err.Error())
	}

	commands := []string{
		"cd /where/is/the/gopath",
		"cd src/myproject",
		"ls",
		"exit",
	}
	for _, cmd := range commands {
		if _, err = fmt.Fprintln(stdin, cmd); err != nil {
			log.Fatal(err)
		}
	}

	err = session.Wait()
	if err != nil {
		log.Fatal(err)
	}

	outputErr := stderr.String()
	fmt.Println("============== ERROR")
	fmt.Println(strings.TrimSpace(outputErr))

	outputString := stdout.String()
	fmt.Println("============== OUTPUT")
	fmt.Println(strings.TrimSpace(outputString))
}
