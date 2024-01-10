package utils

func CheckAccessKey(accessKey string) bool {
	return accessKey == "123456"
}

func CheckSecretAccessKey(secretAccessKey string) bool {
	return secretAccessKey == "123456"
}

func CheckApiCredentials(accessKey, secretAccessKey string) bool {
	return CheckAccessKey(accessKey) && CheckSecretAccessKey(secretAccessKey)
}
