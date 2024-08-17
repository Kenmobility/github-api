package helpers

import "os"

// IsLocal returns true or false depending on APP_ENV environmental variable's value
func IsLocal() bool {
	return os.Getenv("APP_ENV") == "" || os.Getenv("APP_ENV") == "local"
}

// Getenv gets the env variable value or set a default if empty
func Getenv(variable string, defaultValue ...string) string {
	env := os.Getenv(variable)
	if env == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return ""
	}
	return env
}
