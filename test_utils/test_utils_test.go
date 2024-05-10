package test_utils_test

import (
	testUtils "jira-for-peasants/test_utils"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMigrationScripts(t *testing.T) {
	testUtils.SkipCI(t)
	defer testUtils.RemoveTempMigrations()
	_, e := testUtils.GetMigrationScripts()
	assert.Nil(t, e)
}
func TestCopyTestScripts(t *testing.T) {
	testUtils.SkipCI(t)
	defer testUtils.RemoveTempTestScripts()
	testScripts := []string{"script1.sql", "script2.sql", "script3.sql"}
	err := CreateTestScripts(testScripts)
	if err != nil {
		t.Fatal(err)
	}
	defer RemoveTestScripts(testScripts)
	tempScripts, err := testUtils.CopyTestScripts(0, testScripts)
	assert.Nil(t, err)

	// Verify the existence of the temporary test scripts
	for _, script := range tempScripts {
		_, err := os.Stat(script)
		assert.False(t, os.IsNotExist(err))
	}
}

func CreateTestScripts(scripts []string) error {
	for _, script := range scripts {
		file, err := os.Create(filepath.Join("..", "testdata", script))
		if err != nil {
			return err
		}
		defer file.Close()
	}
	return nil
}

func RemoveTestScripts(scripts []string) error {
	for _, script := range scripts {
		err := os.Remove(filepath.Join("..", "testdata", script))
		if err != nil {
			return err
		}
	}
	return nil
}
