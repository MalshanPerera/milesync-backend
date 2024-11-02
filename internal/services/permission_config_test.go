package services_test

import (
	"jira-for-peasants/internal/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProjectPermissions(t *testing.T) {
	_, e := services.GetAllProjectPermissions()
	assert.Nil(t, e)
}

func TestGetAllTaskPermissions(t *testing.T) {
	_, e := services.GetAllTaskPermissions()
	assert.Nil(t, e)
}
