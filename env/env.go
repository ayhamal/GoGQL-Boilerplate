package env

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

// Create Fiber Server struct
type AppCredentials struct {
	Debug      bool   `yaml:"debug"`
	SigningKey string `yaml:"signing_key"`
	Port       uint16 `yaml:"port"`
}

// Create Database Credentials struct
type PostgresCredentials struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	SslMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timezone"`
}

// Create Smtp Broker Credentials struct
type SmtpCredentials struct {
	Host     string `yaml:"host"`
	Port     uint16 `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Security string `yaml:"security"`
}

// Create Configuration struct
type Env struct {
	App        AppCredentials      `yaml:"app"`
	PostgresDb PostgresCredentials `yaml:"postgresql"`
	SmtpBroker SmtpCredentials     `yaml:"smtp_broker"`
}

// Load env configuration assets
func Load(filename string) (*Env, error) {
	// Reference to Env Pointer
	env := &Env{}
	// Read env file
	content, err := os.ReadFile(filename)
	// Handle reading errors
	if err != nil {
		return nil, err
	}
	// Unmarshal env yaml file
	err = yaml.Unmarshal(content, env)
	// Handle unmarshal errors
	if err != nil {
		return nil, err
	}
	// Return reference if no has errors
	return env, nil
}

// The New function creates a new instance of the Env struct by loading the environment path.
func New() (*Env, error) {
	return Load(GetEnvPath())
}

func GetEnvPath() string {
	// Get the absolute path of the current file
	absPath, err := filepath.Abs(os.Args[0])
	// Handle errors
	if err != nil {
		log.Println("Error obtaining the absolute path of .env.yml file:", err)
		return ""
	}
	// Get the parent directory using filepath.Dir()
	parentDir := filepath.Dir(absPath)
	// Get the parent directory using filepath.Dir()
	secondParentDir := filepath.Dir(parentDir)
	// Return the absolute path of the .env.yml file
	return secondParentDir+"/.env.yml"
}
