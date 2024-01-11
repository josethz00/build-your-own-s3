package utils

import (
	"bufio"
	"log"
	"os"
	"path"
)

func CheckAccessKey(accessKey string) bool {
	f, err := os.Open(path.Join("..", "accesskeys.txt"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if accessKey == scanner.Text() {
			return true
		}
	}

	return false
}

func CheckSecretAccessKey(secretAccessKey string) bool {
	f, err := os.Open(path.Join("..", "secretaccesskeys.txt"))

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if secretAccessKey == scanner.Text() {
			return true
		}
	}

	return false
}

func CheckApiCredentials(accessKey, secretAccessKey string) bool {
	return CheckAccessKey(accessKey) && CheckSecretAccessKey(secretAccessKey)
}
