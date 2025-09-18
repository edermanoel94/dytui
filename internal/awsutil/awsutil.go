package awsutil

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrCredentialNotFound = errors.New("credentials not found in the .aws/credentials")
)

type Credential struct {
	Name            string
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

func LoadProfile(profile string) (*Credential, error) {

	credentials, err := loadAWSCredentials()

	if err != nil {
		return nil, err
	}

	for _, cred := range credentials {
		if cred.Name == profile {
			return &cred, nil
		}
	}

	return nil, ErrCredentialNotFound
}

func loadAWSCredentials() ([]Credential, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(home, ".aws", "credentials")
	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	credentials := make([]Credential, 0)
	credentialsMap := make(map[string]Credential)
	var currentProfileName string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		// profile header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentProfileName = strings.Trim(line, "[]")
			credentialsMap[currentProfileName] = Credential{Name: currentProfileName}
			continue
		}

		// key = value
		if currentProfileName != "" && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)

			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			cred := credentialsMap[currentProfileName]

			switch key {
			case "aws_access_key_id":
				cred.AccessKeyID = value
			case "aws_secret_access_key":
				cred.SecretAccessKey = value
			case "region":
				cred.Region = value
			}

			credentialsMap[currentProfileName] = cred
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	for _, cred := range credentialsMap {
		credentials = append(credentials, cred)
	}

	return credentials, nil
}
