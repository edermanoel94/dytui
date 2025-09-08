package awsutil

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

type Credentials struct {
	Name            string
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

// Profiles maps profile names to their credentials
// LoadAWSCredentials reads the AWS credentials file and parses it into a map of profiles
func LoadAWSCredentials() ([]Credentials, error) {
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

	profiles := make([]Credentials, 0)

	var currentProfile string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, ";") {
			continue
		}

		var credentials Credentials

		// profile header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentProfile = strings.Trim(line, "[]")
			credentials.Name = currentProfile
			continue
		}

		// key = value
		if credentials.Name != "" && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			switch key {
			case "aws_access_key_id":
				credentials.AccessKeyID = value
			case "aws_secret_access_key":
				credentials.SecretAccessKey = value
			case "region":
				credentials.Region = value
			}

			profiles = append(profiles, credentials)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return profiles, nil
}
