package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
)

var configFilePath string

type IntegrationRecord struct {
	Token      string
	DatabaseId string
}

type Config struct {
	Integrations map[string]IntegrationRecord
}

func NewConfig(configFile string) *Config {
	configFilePath = configFile
	return &Config{
		Integrations: make(map[string]IntegrationRecord),
	}
}

func (c *Config) read() error {
	f, err := os.OpenFile(configFilePath, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	_, err = toml.NewDecoder(reader).Decode(c)
	return err
}

func (c *Config) overwrite() error {
	f, err := os.OpenFile(configFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	return toml.NewEncoder(writer).Encode(*c)
}

func (c *Config) integrationExists(Name string) bool {
	_, found := c.Integrations[Name]
	return found
}

func (c *Config) addIntegration(Name string, Token string, DatabaseId string) {
	c.Integrations[Name] = IntegrationRecord{
		Token:      Token,
		DatabaseId: DatabaseId,
	}
}

func (c *Config) removeIntegration(Name string) {
	delete(c.Integrations, Name)
}

func (c *Config) Get() error {
	return c.read()
}

func (c *Config) AddIntegrationRecord(Name string, Token string, DatabaseId string) error {
	if Name = strings.TrimSpace(Name); Name == "" {
		return fmt.Errorf("Integration Record is required to have a symbolic name.\n")
	}

	if Token = strings.TrimSpace(Token); Token == "" {
		return fmt.Errorf("Token can not be an empty string (notion supplies it to you when new Integration is created)\n")
	}

	if DatabaseId = strings.TrimSpace(DatabaseId); DatabaseId == "" {
		return fmt.Errorf("DatabaseId can not be an empty string (lookup DatabaseId that you added Integration to)\n")
	}

	err := c.read()
	if err != nil {
		return err
	}

	if c.integrationExists(Name) {
		return fmt.Errorf("A Notion Integration Record with the same name already exists. Choose a different name or delete the old entry.\n")
	}

	c.addIntegration(Name, Token, DatabaseId)
	return c.overwrite()
}

func (c *Config) RemoveIntegrationRecord(Name string) error {
	if Name == "" {
		return fmt.Errorf("Name of Integration Record can not be empty\n")
	}

	err := c.read()
	if err != nil {
		return err
	}

	if !c.integrationExists(Name) {
		return fmt.Errorf("Integration Record with Name - %s was not found in the config file. Can not remove.\n", Name)
	}

	c.removeIntegration(Name)
	return c.overwrite()
}
