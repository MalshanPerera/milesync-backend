package services

import (
	"context"
	repo "jira-for-peasants/internal/repositories"
	"time"
)

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

type TaskService struct {
	taskRepository *repo.TaskRepository
}

func NewTaskService(taskRepo *repo.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, params CreateTaskParams) (repo.TaskModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskModel{}, err
	}

	defer func() {
		err = s.taskRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	task, err := s.taskRepository.CreateTask(ctx, tx, repo.CreateTaskParams{
		Title:          params.Title,
		Description:    params.Description,
		UserID:         params.UserID,
		AssigneeID:     params.AssigneeID,
		ReporterID:     params.ReporterID,
		OrganizationID: params.OrganizationID,
		ProjectID:      params.ProjectID,
		StatusID:       params.StatusID,
		Priority:       params.Priority,
		DueDate:        params.DueDate,
		OrderIndex:     params.OrderIndex,
	})

	if err != nil {
		return repo.TaskModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.TaskModel{}, err
	}

	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, params UpdateTaskParams) (repo.TaskModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskModel{}, err
	}

	defer func() {
		err = s.taskRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	task, err := s.taskRepository.UpdateTask(ctx, tx, repo.UpdateTaskParams{
		ID:             params.ID,
		Title:          params.Title,
		Description:    params.Description,
		UserID:         params.UserID,
		AssigneeID:     params.AssigneeID,
		ReporterID:     params.ReporterID,
		OrganizationID: params.OrganizationID,
		ProjectID:      params.ProjectID,
		StatusID:       params.StatusID,
		Priority:       params.Priority,
		DueDate:        params.DueDate,
		OrderIndex:     params.OrderIndex,
	})

	if err != nil {
		return repo.TaskModel{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return repo.TaskModel{}, err
	}

	return task, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id string) error {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		err = s.taskRepository.RollbackTx(ctx, tx)
		if err != nil {
			return
		}
	}()

	err = s.taskRepository.DeleteTask(ctx, tx, id)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
func (s *TaskService) GetTask(ctx context.Context, id string) (repo.TaskModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	task, err := s.taskRepository.GetTask(ctx, tx, id)
	if err != nil {
		return repo.TaskModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.TaskModel{}, err
	}

	return task, nil
}

func (s *TaskService) GetTasks(ctx context.Context, organizationID, projectID string) ([]repo.TaskModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	tasks, err := s.taskRepository.GetTasks(ctx, tx, organizationID, projectID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *TaskService) CreateTaskComment(ctx context.Context, params CreateTaskCommentParams) (repo.TaskCommentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskCommentModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	comment, err := s.taskRepository.CreateTaskComment(ctx, tx, repo.CreateTaskCommentParams{
		TaskID:  params.TaskID,
		UserID:  params.UserID,
		Comment: params.Comment,
	})
	if err != nil {
		return repo.TaskCommentModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.TaskCommentModel{}, err
	}

	return comment, nil
}

func (s *TaskService) UpdateTaskComment(ctx context.Context, params UpdateTaskCommentParams) (repo.TaskCommentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskCommentModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	comment, err := s.taskRepository.UpdateTaskComment(ctx, tx, repo.UpdateTaskCommentParams{
		ID:      params.ID,
		Comment: params.Comment,
	})
	if err != nil {
		return repo.TaskCommentModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.TaskCommentModel{}, err
	}

	return comment, nil
}

func (s *TaskService) DeleteTaskComment(ctx context.Context, id string) error {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	if err := s.taskRepository.DeleteTaskComment(ctx, tx, id); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *TaskService) GetTaskComments(ctx context.Context, taskID string) ([]repo.TaskCommentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	comments, err := s.taskRepository.GetTaskComments(ctx, tx, taskID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return comments, nil
}

func (s *TaskService) CreateTaskAttachment(ctx context.Context, params CreateTaskAttachmentParams) (repo.TaskAttachmentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskAttachmentModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	attachment, err := s.taskRepository.CreateTaskAttachment(ctx, tx, repo.CreateTaskAttachmentParams{
		TaskID:   params.TaskID,
		UserID:   params.UserID,
		FileName: params.FileName,
		FilePath: params.FilePath,
		FileSize: params.FileSize,
		MimeType: params.MimeType,
	})
	if err != nil {
		return repo.TaskAttachmentModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.TaskAttachmentModel{}, err
	}

	return attachment, nil
}

func (s *TaskService) DeleteTaskAttachment(ctx context.Context, id string) error {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	if err := s.taskRepository.DeleteTaskAttachment(ctx, tx, id); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *TaskService) GetTaskAttachments(ctx context.Context, taskID string) ([]repo.TaskAttachmentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	attachments, err := s.taskRepository.GetTaskAttachments(ctx, tx, taskID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return attachments, nil
}

func (s *TaskService) CreateTaskAssignee(ctx context.Context, params CreateTaskAssigneeParams) (repo.TaskAssigneeModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.TaskAssigneeModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	assignee, err := s.taskRepository.CreateTaskAssignee(ctx, tx, repo.CreateTaskAssigneeParams{
		TaskID: params.TaskID,
		UserID: params.UserID,
	})
	if err != nil {
		return repo.TaskAssigneeModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.TaskAssigneeModel{}, err
	}

	return assignee, nil
}

func (s *TaskService) DeleteTaskAssignee(ctx context.Context, taskID, userID string) error {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	if err := s.taskRepository.DeleteTaskAssignee(ctx, tx, taskID, userID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *TaskService) GetTaskAssignees(ctx context.Context, taskID string) ([]repo.TaskAssigneeModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	assignees, err := s.taskRepository.GetTaskAssignees(ctx, tx, taskID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return assignees, nil
}

func (s *TaskService) CreateCommentAttachment(ctx context.Context, params CreateCommentAttachmentParams) (repo.CommentAttachmentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return repo.CommentAttachmentModel{}, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	attachment, err := s.taskRepository.CreateCommentAttachment(ctx, tx, repo.CreateCommentAttachmentParams{
		CommentID: params.CommentID,
		UserID:    params.UserID,
		FileName:  params.FileName,
		FilePath:  params.FilePath,
		FileSize:  params.FileSize,
		MimeType:  params.MimeType,
	})
	if err != nil {
		return repo.CommentAttachmentModel{}, err
	}

	if err = tx.Commit(ctx); err != nil {
		return repo.CommentAttachmentModel{}, err
	}

	return attachment, nil
}

func (s *TaskService) DeleteCommentAttachment(ctx context.Context, id string) error {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	if err := s.taskRepository.DeleteCommentAttachment(ctx, tx, id); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *TaskService) GetCommentAttachments(ctx context.Context, commentID string) ([]repo.CommentAttachmentModel, error) {
	tx, err := s.taskRepository.BeginTx(ctx)
	if err != nil {
		return nil, err
	}
	defer s.taskRepository.RollbackTx(ctx, tx)

	attachments, err := s.taskRepository.GetCommentAttachments(ctx, tx, commentID)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return attachments, nil
}
