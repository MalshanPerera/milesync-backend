package services

import (
	"encoding/json"
	"io"
	"os"
)

type Permissions []string

type AppRole struct {
	Key         string      `json:"key"`
	Name        string      `json:"name"`
	Permissions Permissions `json:"permissions"`
	DependsOn   []string    `json:"depends_on"`
}

var projectPermissionsCache []AppRole
var taskPermissionsCache []AppRole

// GetAllProjectPermissions retrieves all project permissions from the cache or reads them from the project_permissions.json file.
// If the permissions are already cached, it returns the cached permissions. Otherwise, it reads the permissions from the JSON file,
// caches them, and returns the permissions.
func GetAllProjectPermissions() ([]AppRole, error) {
	if projectPermissionsCache != nil {
		return projectPermissionsCache, nil
	}

	jsonFile, err := os.Open("../auth/permission/project_permissions.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var permissions []AppRole
	json.Unmarshal(byteValue, &permissions)

	projectPermissionsCache = permissions

	return permissions, nil
}

// GetAllTaskPermissions retrieves all task permissions from the task_permissions.json file.
// If the permissions are already cached, it returns the cached permissions.
// Otherwise, it reads the permissions from the file, caches them, and returns them.
func GetAllTaskPermissions() ([]AppRole, error) {
	if taskPermissionsCache != nil {
		return taskPermissionsCache, nil
	}

	jsonFile, err := os.Open("../auth/permission/task_permissions.json")
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)
	var permissions []AppRole
	json.Unmarshal(byteValue, &permissions)

	taskPermissionsCache = permissions

	return permissions, nil
}
