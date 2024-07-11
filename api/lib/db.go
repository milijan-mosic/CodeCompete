package lib

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getDBCredentials(env map[string]string) string {
	var connection string

	switch env["APP_MODE"] {
	case "LOCAL":
		connection = "postgresql://" + env["LOCAL_DB_ADMIN_USERNAME"] + ":" + env["LOCAL_DB_ADMIN_PASSWORD"] + "@postgres:" + env["DB_PORT"] + "/" + env["DB_NAME"]
	case "DEV":
		connection = "postgresql://" + env["DEV_DB_ADMIN_USERNAME"] + ":" + env["DEV_DB_ADMIN_PASSWORD"] + "@postgres:" + env["DB_PORT"] + "/" + env["DB_NAME"]
	case "PROD":
		connection = "postgresql://" + env["PROD_DB_ADMIN_USERNAME"] + ":" + env["PROD_DB_ADMIN_PASSWORD"] + "@postgres:" + env["DB_PORT"] + "/" + env["DB_NAME"]
	}

	return connection
}

func LoadEnvVariables(filePath string) (map[string]string, error) {
	err := godotenv.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	envVars := make(map[string]string)

	for _, env := range os.Environ() {
		key, value, err := parseEnv(env)

		if err != nil {
			return nil, fmt.Errorf("error parsing environment variable: %w", err)
		}

		envVars[key] = value
	}

	return envVars, nil
}

func parseEnv(env string) (key, value string, err error) {
	pair := strings.SplitN(env, "=", 2)

	if len(pair) != 2 {
		return "", "", fmt.Errorf("malformed environment variable: %s", env)
	}

	return pair[0], pair[1], nil
}
