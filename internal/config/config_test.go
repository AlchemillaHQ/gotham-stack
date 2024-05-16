package config

import (
	"io/ioutil"
	"os"
	"testing"
)

// Sample configuration for testing
const testConfigContent = `{
	"environment": "test",
	"port": 8080,
	"database_url": "postgres://user:password@localhost:5432/testdb",
	"log_level": "debug",
	"session_secret": "secret",
	"admins": [
		{
			"username": "admin1",
			"email": "admin1@example.com",
			"password": "password1"
		}
	],
	"data_path": "./testdata"
}`

// Helper function to create a test config file
func createTestConfigFile(t *testing.T) string {
	file, err := ioutil.TempFile("", "config_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	_, err = file.Write([]byte(testConfigContent))
	if err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}

	file.Close()
	return file.Name()
}

// TestParseConfig tests the ParseConfig function
func TestParseConfig(t *testing.T) {
	configFile := createTestConfigFile(t)
	defer os.Remove(configFile) // Clean up the temp file

	config := ParseConfig(configFile)

	if config.Environment != "test" {
		t.Errorf("Expected environment to be 'test', got '%s'", config.Environment)
	}

	if config.Port != 8080 {
		t.Errorf("Expected port to be 8080, got %d", config.Port)
	}

	if config.DatabaseURL != "postgres://user:password@localhost:5432/testdb" {
		t.Errorf("Expected database_url to be 'postgres://user:password@localhost:5432/testdb', got '%s'", config.DatabaseURL)
	}

	if config.Admins[0].Username != "admin1" {
		t.Errorf("Expected first admin username to be 'admin1', got '%s'", config.Admins[0].Username)
	}
}

// TestSetupDataPath tests the SetupDataPath function
func TestSetupDataPath(t *testing.T) {
	// Test case: data path is not set in the config
	ParsedConfig = &Config{}
	err := SetupDataPath()
	if err != nil {
		t.Fatalf("Failed to set up data path: %v", err)
	}

	if ParsedConfig.DataPath != "./data" {
		t.Errorf("Expected data path to be './data', got '%s'", ParsedConfig.DataPath)
	}

	// Clean up the test directory
	os.RemoveAll("./data")

	// Test case: data path is set in the config
	ParsedConfig = &Config{DataPath: "./testdata"}
	err = SetupDataPath()
	if err != nil {
		t.Fatalf("Failed to set up data path: %v", err)
	}

	// Check if the directory exists
	if _, err := os.Stat("./testdata"); os.IsNotExist(err) {
		t.Fatalf("Expected directory './testdata' to exist")
	}

	// Clean up the test directory
	os.RemoveAll("./testdata")
}
