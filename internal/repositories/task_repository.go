package repositories

import (
	"context"
	datastore "jira-for-peasants/db"
	db "jira-for-peasants/db/sqlc"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type TaskModel db.Task
type TaskCommentModel db.TaskComment
type TaskAttachmentModel db.TaskAttachment
type TaskAssigneeModel db.TaskAssignee
type CommentAttachmentModel db.CommentAttachment

type CreateTaskParams struct {
	Title          string
	Description    string
	UserID         string
	AssigneeID     string
	ReporterID     string
	OrganizationID string
	ProjectID      string
	StatusID       string
	Priority       int16
	DueDate        time.Time
	OrderIndex     int32
}

type UpdateTaskParams struct {
	ID             string
	Title          string
	Description    string
	UserID         string
	AssigneeID     string
	ReporterID     string
	OrganizationID string
	ProjectID      string
	StatusID       string
	Priority       int16
	DueDate        time.Time
	OrderIndex     int32
}

type CreateTaskCommentParams struct {
	TaskID  string
	UserID  string
	Comment string
}

type UpdateTaskCommentParams struct {
	ID      string
	Comment string
}

type CreateTaskAttachmentParams struct {
	TaskID   string
	UserID   string
	FileName string
	FilePath string
	FileSize int64
	MimeType string
}

type CreateTaskAssigneeParams struct {
	TaskID string
	UserID string
}

type CreateCommentAttachmentParams struct {
	CommentID string
	UserID    string
	FileName  string
	FilePath  string
	FileSize  int64
	MimeType  string
}

type TaskRepository struct {
	*datastore.Trx
	db *datastore.DB
}

func NewTaskRepository(db *datastore.DB) *TaskRepository {
	return &TaskRepository{
		db:  db,
		Trx: datastore.NewTrx(db),
	}
}

func (repo *TaskRepository) CreateTask(ctx context.Context, tx pgx.Tx, params CreateTaskParams) (TaskModel, error) {
	task, e := repo.db.GetQuery().WithTx(tx).CreateTask(ctx, db.CreateTaskParams{
		Title:          params.Title,
		Description:    &params.Description,
		UserID:         params.UserID,
		AssignerID:     &params.AssigneeID,
		ReporterID:     &params.ReporterID,
		OrganizationID: params.OrganizationID,
		ProjectID:      params.ProjectID,
		StatusID:       params.StatusID,
		Priority:       params.Priority,
		DueDate:        pgtype.Timestamp{Time: params.DueDate, Valid: true},
		OrderIndex:     &params.OrderIndex,
	})

	if e != nil {
		return TaskModel{}, e
	}

	return TaskModel(task), nil
}

func (repo *TaskRepository) UpdateTask(ctx context.Context, tx pgx.Tx, params UpdateTaskParams) (TaskModel, error) {
	task, e := repo.db.GetQuery().WithTx(tx).UpdateTask(ctx, db.UpdateTaskParams{
		ID:             params.ID,
		Title:          params.Title,
		Description:    &params.Description,
		UserID:         params.UserID,
		AssignerID:     &params.AssigneeID,
		ReporterID:     &params.ReporterID,
		OrganizationID: params.OrganizationID,
		ProjectID:      params.ProjectID,
		StatusID:       params.StatusID,
		Priority:       params.Priority,
		DueDate:        pgtype.Timestamp{Time: params.DueDate, Valid: true},
		OrderIndex:     &params.OrderIndex,
	})

	if e != nil {
		return TaskModel{}, e
	}

	return TaskModel(task), nil
}

func (repo *TaskRepository) DeleteTask(ctx context.Context, tx pgx.Tx, id string) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteTask(ctx, id)
	if e != nil {
		return e
	}

	return nil
}

func (repo *TaskRepository) GetTask(ctx context.Context, tx pgx.Tx, id string) (TaskModel, error) {
	task, e := repo.db.GetQuery().WithTx(tx).GetTask(ctx, id)
	if e != nil {
		return TaskModel{}, e
	}

	return TaskModel(task), nil
}

func (repo *TaskRepository) GetTasks(ctx context.Context, tx pgx.Tx, organizationID string, projectID string) ([]TaskModel, error) {
	tasks, e := repo.db.GetQuery().WithTx(tx).GetTasks(ctx, db.GetTasksParams{
		OrganizationID: organizationID,
		ProjectID:      projectID,
	})
	if e != nil {
		return []TaskModel{}, e
	}

	results := make([]TaskModel, 0)
	for _, task := range tasks {
		results = append(results, TaskModel(task))
	}

	return results, nil
}

func (repo *TaskRepository) CreateTaskComment(ctx context.Context, tx pgx.Tx, params CreateTaskCommentParams) (TaskCommentModel, error) {
	comment, e := repo.db.GetQuery().WithTx(tx).CreateTaskComment(ctx, db.CreateTaskCommentParams{
		TaskID:  params.TaskID,
		UserID:  params.UserID,
		Comment: params.Comment,
	})
	if e != nil {
		return TaskCommentModel{}, e
	}

	return TaskCommentModel(comment), nil
}

func (repo *TaskRepository) UpdateTaskComment(ctx context.Context, tx pgx.Tx, params UpdateTaskCommentParams) (TaskCommentModel, error) {
	comment, e := repo.db.GetQuery().WithTx(tx).UpdateTaskComment(ctx, db.UpdateTaskCommentParams{
		ID:      params.ID,
		Comment: params.Comment,
	})
	if e != nil {
		return TaskCommentModel{}, e
	}

	return TaskCommentModel(comment), nil
}

func (repo *TaskRepository) DeleteTaskComment(ctx context.Context, tx pgx.Tx, id string) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteTaskComment(ctx, id)
	if e != nil {
		return e
	}

	return nil
}

func (repo *TaskRepository) GetTaskComments(ctx context.Context, tx pgx.Tx, taskID string) ([]TaskCommentModel, error) {
	comments, e := repo.db.GetQuery().WithTx(tx).GetTaskComments(ctx, taskID)
	if e != nil {
		return []TaskCommentModel{}, e
	}

	results := make([]TaskCommentModel, 0)
	for _, comment := range comments {
		results = append(results, TaskCommentModel(comment))
	}

	return results, nil
}

func (repo *TaskRepository) CreateTaskAttachment(ctx context.Context, tx pgx.Tx, params CreateTaskAttachmentParams) (TaskAttachmentModel, error) {
	attachment, e := repo.db.GetQuery().WithTx(tx).CreateTaskAttachment(ctx, db.CreateTaskAttachmentParams{
		TaskID:   params.TaskID,
		UserID:   params.UserID,
		FileName: params.FileName,
		FilePath: params.FilePath,
		FileSize: params.FileSize,
		MimeType: params.MimeType,
	})
	if e != nil {
		return TaskAttachmentModel{}, e
	}

	return TaskAttachmentModel(attachment), nil
}

func (repo *TaskRepository) DeleteTaskAttachment(ctx context.Context, tx pgx.Tx, id string) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteTaskAttachment(ctx, id)
	if e != nil {
		return e
	}

	return nil
}

func (repo *TaskRepository) GetTaskAttachments(ctx context.Context, tx pgx.Tx, taskID string) ([]TaskAttachmentModel, error) {
	attachments, e := repo.db.GetQuery().WithTx(tx).GetTaskAttachments(ctx, taskID)
	if e != nil {
		return []TaskAttachmentModel{}, e
	}

	results := make([]TaskAttachmentModel, 0)
	for _, attachment := range attachments {
		results = append(results, TaskAttachmentModel(attachment))
	}

	return results, nil
}

func (repo *TaskRepository) CreateTaskAssignee(ctx context.Context, tx pgx.Tx, params CreateTaskAssigneeParams) (TaskAssigneeModel, error) {
	assignee, e := repo.db.GetQuery().WithTx(tx).CreateTaskAssignee(ctx, db.CreateTaskAssigneeParams{
		TaskID: params.TaskID,
		UserID: params.UserID,
	})
	if e != nil {
		return TaskAssigneeModel{}, e
	}

	return TaskAssigneeModel(assignee), nil
}

func (repo *TaskRepository) DeleteTaskAssignee(ctx context.Context, tx pgx.Tx, taskID string, userID string) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteTaskAssignee(ctx, db.DeleteTaskAssigneeParams{
		TaskID: taskID,
		UserID: userID,
	})
	if e != nil {
		return e
	}

	return nil
}

func (repo *TaskRepository) GetTaskAssignees(ctx context.Context, tx pgx.Tx, taskID string) ([]TaskAssigneeModel, error) {
	assignees, e := repo.db.GetQuery().WithTx(tx).GetTaskAssignees(ctx, taskID)
	if e != nil {
		return []TaskAssigneeModel{}, e
	}

	results := make([]TaskAssigneeModel, 0)
	for _, assignee := range assignees {
		results = append(results, TaskAssigneeModel(assignee))
	}

	return results, nil
}

func (repo *TaskRepository) CreateCommentAttachment(ctx context.Context, tx pgx.Tx, params CreateCommentAttachmentParams) (CommentAttachmentModel, error) {
	attachment, e := repo.db.GetQuery().WithTx(tx).CreateCommentAttachment(ctx, db.CreateCommentAttachmentParams{
		CommentID: params.CommentID,
		UserID:    params.UserID,
		FileName:  params.FileName,
		FilePath:  params.FilePath,
		FileSize:  params.FileSize,
		MimeType:  params.MimeType,
	})
	if e != nil {
		return CommentAttachmentModel{}, e
	}

	return CommentAttachmentModel(attachment), nil
}

func (repo *TaskRepository) DeleteCommentAttachment(ctx context.Context, tx pgx.Tx, id string) error {
	e := repo.db.GetQuery().WithTx(tx).DeleteCommentAttachment(ctx, id)
	if e != nil {
		return e
	}

	return nil
}

func (repo *TaskRepository) GetCommentAttachments(ctx context.Context, tx pgx.Tx, commentID string) ([]CommentAttachmentModel, error) {
	attachments, e := repo.db.GetQuery().WithTx(tx).GetCommentAttachments(ctx, commentID)
	if e != nil {
		return []CommentAttachmentModel{}, e
	}

	results := make([]CommentAttachmentModel, 0)
	for _, attachment := range attachments {
		results = append(results, CommentAttachmentModel(attachment))
	}

	return results, nil
}
