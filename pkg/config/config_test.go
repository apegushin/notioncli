package config_test

import (
	//"math/rand"
	"os"
	"testing"

	"github.com/apegushin/notioncli/pkg/config"
)

const (
	emptyConfigFileName string = "empty-config-file-do-not-add-records-to.toml"
	testConfigFileName  string = "test-config.toml"

	testIntegrationName       string = "TestIntegRecName"
	testIntegrationToken      string = "TestIntegToken"
	testIntegrationDatabaseId string = "TestIntegDatabaseId"
)

func TestGetEmptyConfig(t *testing.T) {
	c := config.NewConfig(emptyConfigFileName)
	if len(c.Integrations) != 0 {
		t.Errorf("Expected to read empty Config object from an empty config file. Fail!\n")
	}
}

func TestAddIntegrationRecordInvalidArgs(t *testing.T) {
	emptyTestStrings := []string{
		" ", "  ", "\t", " \t", " \t", " \t ", "\n", " \n", "\n ", " \n ", "\t \n ", " \t \n", " \t \n ", "\n \t ", " \n \t", " \n \t ",
	}
	nonEmptyTestString := "definitely-not-empty"

	c := config.NewConfig(testConfigFileName)

	for _, emptyTestString := range emptyTestStrings {
		err := c.AddIntegrationRecord(emptyTestString, nonEmptyTestString, nonEmptyTestString)
		if err == nil {
			t.Errorf("Expected AddIntegrationRecord to error out on Name being empty string. No error received. Fail!")
		}

		err = c.AddIntegrationRecord(nonEmptyTestString, emptyTestString, nonEmptyTestString)
		if err == nil {
			t.Errorf("Expected AddIntegrationRecord to error out on Token being empty string. No error received. Fail!")
		}

		err = c.AddIntegrationRecord(nonEmptyTestString, nonEmptyTestString, emptyTestString)
		if err == nil {
			t.Errorf("Expected AddIntegrationRecord to error out on DatabaseId being empty string. No error received. Fail!")
		}
	}
}

func TestAddIntegrationRecordOneRecord(t *testing.T) {
	c := config.NewConfig(testConfigFileName)
	defer os.Remove(testConfigFileName)

	err := c.AddIntegrationRecord(testIntegrationName, testIntegrationToken, testIntegrationDatabaseId)
	if err != nil {
		t.Errorf("Expected no error when adding one valid record to empty config. Integration Name - %s, Token - %s, DatabaseId - %s. Fail!\n",
			testIntegrationName, testIntegrationToken, testIntegrationDatabaseId)
	}

	err = c.Get()
	if err != nil {
		t.Errorf("Expected no error when reading test config. Fail!\n")
	}

	if len(c.Integrations) != 1 {
		t.Errorf("Expected only one Integration record in the test config. Fail!\n")
	}
	record, found := c.Integrations[testIntegrationName]
	if !found {
		t.Errorf("Expected Integration record with Name - %s. Fail!\n", testIntegrationName)
	}

	if record.Token != testIntegrationToken || record.DatabaseId != testIntegrationDatabaseId {
		t.Errorf("Expected Integration record with Token - %s and DatabaseId - %s. Fail!\n", testIntegrationToken, testIntegrationDatabaseId)
	}
}

func generateRecords(numRecords uint) []config.IntegrationRecord {
	records := []config.IntegrationRecord{}
	return records
}

//func TestAddIntegrationRecordRandRecords(t *testing.T) {
//
//}
//
//func TestRemoveIntegrationRandRecord(t *testing.T) {
//}
