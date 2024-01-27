package utils

import (
	"bufio"
	"log"
	"os"
)

func CheckAccessKey(accessKey string) bool {
	f, err := os.Open("accesskeys.txt")

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
	f, err := os.Open("secretaccesskeys.txt")

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
