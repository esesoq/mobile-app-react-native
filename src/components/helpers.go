package mobileappreactnative

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GenerateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err!= nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(uuid), nil
}

func GetTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetProjectRoot() string {
	cwd, err := os.Getwd()
	if err!= nil {
		log.Fatal(err)
	}
	return filepath.Dir(filepath.Dir(cwd))
}

func GetProjectName() string {
	return strings.TrimPrefix(filepath.Base(GetProjectRoot()), "mobile-app-react-native-")
}

func GetCommitHash() string {
	hash := os.Getenv("COMMIT_HASH")
	if hash == "" {
		return ""
	}
	return hash
}

func GetEnvironmentVariable(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal(fmt.Sprintf("Environment variable %s is not set", key))
	}
	return value
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StrToBool(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	return b, err
}

func GetOsName() string {
	if strings.Contains(os.Getenv("OS"), "Windows") {
		return "Windows"
	} else if strings.Contains(os.Getenv("OS"), "Darwin") {
		return "Darwin"
	} else {
		return "Linux"
	}
}