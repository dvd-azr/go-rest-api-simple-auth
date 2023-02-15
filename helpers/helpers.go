package helpers

import (
	"fmt"
	"log"
	"os"
)

// Handle to return the error it self
func ErrorReturnHandler(err error) error {
	if err != nil {
		return err
	}
	return nil
}

// Handle the error with log.fatal()
func ErrorLogFatalHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Get value from .env by variable key with oc.Getenv()
func Env(key string, defaultOpt ...string) string {
	if ok := os.Getenv(key); ok != "" && key != "" {
		return ok
	}

	if defaultOpt != nil && checkEnvKey(defaultOpt[0]) {
		return os.Getenv(defaultOpt[0])
	}
	return ""
}

//Checking that an environment variable key is present or not.
func checkEnvKey(key string) bool {
	envKey, ok := os.LookupEnv(key)
	if !ok {
		fmt.Printf("%s is not present\n", key)
		return false
	} else {
		fmt.Printf("%s : %s\n", key, envKey)
		return true
	}
}
