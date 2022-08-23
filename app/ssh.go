package app

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

type FileItem struct {
	fileName string
	data     []byte
}

func Upload(paramFile FileItem) {
	fileName := paramFile.fileName
	content := paramFile.data

	textFileDirectory := filepath.Join("/dev/shm")

	err := os.WriteFile(filepath.Join(textFileDirectory, path.Clean(fileName)), content, 0666)

	if err != nil {
		log.Fatal(err)
	}

	hostKeyCallback, err := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))

	if err != nil {
		log.Fatal(err)
	}

	config := &ssh.ClientConfig{
		User:            os.Getenv("SSH_USERNAME"),
		Auth:            []ssh.AuthMethod{publicKey()},
		HostKeyCallback: hostKeyCallback,
	}

	// connect
	conn, err := ssh.Dial("tcp", os.Getenv("SSH_HOST")+":22", config)

	if err != nil {
		fmt.Println(err)
	}

	defer func(conn *ssh.Client) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	// create new SFTP client
	client, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Println(err)
	}
	defer func(client *sftp.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)

	// create destination file
	dstFile, err := client.Create(os.Getenv("SSH_DESTINATION_FOLDER") + "/" + fileName)

	if err != nil {
		fmt.Println("error: ", err)
	}

	defer func(dstFile *sftp.File) {
		err := dstFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(dstFile)

	// open source file
	srcFile, err := os.Open(filepath.Join(textFileDirectory, path.Clean(fileName)))

	if err != nil {
		fmt.Println(err)
	}

	// copy source file to destination file
	bytes, err := io.Copy(dstFile, srcFile)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d bytes copied\n", bytes)

	targetPath := filepath.Join(textFileDirectory, path.Clean(fileName))

	if FileExists(targetPath) {
		err = os.Remove(targetPath)

		if err != nil {
			fmt.Println(err)
		}
	}
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return !info.IsDir()
}

func publicKey() ssh.AuthMethod {
	pemBytes, err := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".ssh", "id_rsa"))

	if err != nil {
		fmt.Println(err)
	}

	signer, err := ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(os.Getenv("SSH_PASSWORD")))

	if err != nil {
		panic(err)
	}

	return ssh.PublicKeys(signer)
}
