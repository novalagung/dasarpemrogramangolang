package main

import (
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"os"
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
	// const SSH_ADDRESS = "0.0.0.0:22"
	// const SSH_USERNAME = "user"
	// const SSH_KEY = ""
	// const SSH_PASSWORD = "password"

	// sshConfig := &ssh.ClientConfig{
	// 	User:            SSH_USERNAME,
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// 	Auth: []ssh.AuthMethod{
	// 		ssh.Password(SSH_PASSWORD),
	// 		// PublicKeyFile(SSH_KEY),

	const SSH_ADDRESS = "go.eaciit.com:22"
	const SSH_USERNAME = "developer"
	const SSH_KEY = "/Users/novalagung/Documents/developer.pem"
	const SSH_PASSWORD = ""

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			// ssh.Password(SSH_PASSWORD),
			PublicKeyFile(SSH_KEY),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		log.Fatal("Failed create client sftp client. " + err.Error())
	}

	fDestination, err := sftpClient.Create("/data/nginx/files/test-file.txt")
	if err != nil {
		log.Fatal("Failed to create destination file. " + err.Error())
	}

	fSource, err := os.Open("/Users/novalagung/Desktop/test-file.txt")
	if err != nil {
		log.Fatal("Failed to read source file. " + err.Error())
	}

	_, err = io.Copy(fDestination, fSource)
	if err != nil {
		log.Fatal("Failed copy source file into destination file. " + err.Error())
	}

	log.Println("File copied.")
}
