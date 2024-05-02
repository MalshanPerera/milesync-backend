package test_utils_test

import (
	testUtils "jira-for-peasants/test_utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMigrationScripts(t *testing.T) {
	testUtils.SkipCI(t)
	defer testUtils.RemoveTempMigrations()
	_, e := testUtils.GetMigrationScripts()
	assert.Nil(t, e)
}
