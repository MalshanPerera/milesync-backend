package services

import datastore "jira-for-peasents/db"

type ProjectService struct {
	db *datastore.DB
}

func NewProjectService(db *datastore.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}
