package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var c Config
var emptyConf Config
var testDir string

type Config struct {
	Env         string
	ES256       es256
	PostgresURI postgresURI
	S3          s3
}

type postgresURI struct {
	Prod string
	Dev  string
	Test string
}

type s3 struct {
	BucketNames bucketNames
	Credentials s3Credentials
	Endpoint    string
	Region      string
}

type bucketNames struct {
	Photos photos
}

type photos struct {
	Prod string
	Dev  string
	Test string
}

type s3Credentials struct {
	Id     string
	Secret string
}

type es256 struct {
	PrivateKey string
	PublicKey  string
}

func LoadConfigurations() {
	if c != emptyConf {
		return
	}
	_, file, _, _ := runtime.Caller(0)
	projectRoot := filepath.Dir(path.Join(path.Dir(file)))
	configDir := filepath.Join(projectRoot, "config")
	testDir = filepath.Join(projectRoot, "test")

	viper.AddConfigPath(configDir)
	viper.SetConfigName("secrets")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("Improper config file: secrets.yaml, ", err)
	}
	if err := viper.Unmarshal(&c); err != nil {
		log.Fatalln("Failed to unmarshal configurations, ", err)
	}
	if env := os.Getenv("env"); env != "" {
		c.Env = env
	}
	log.Println("Configurations are successfully prepared")
}

func Get() Config {
	if c == emptyConf {
		LoadConfigurations()
	}
	return c
}

func TestDirPath() string {
	if testDir == "" {
		LoadConfigurations()
	}
	return testDir
}
