package handler

import (
	"context"
	"to-do-app/internal/adapters/controller/service/handler/dto"

	"go.uber.org/zap"
)

type TaskHandler struct {
	ctx     context.Context
	log     *zap.Logger
	Usecase TaskUsecase
}

func NewTaskHandler(uc TaskUsecase, ctx context.Context, logger *zap.Logger) *TaskHandler {
	return &TaskHandler{
		Usecase: uc,
		log:     logger,
		ctx:     ctx,
	}
}

func (h *TaskHandler) CreateTask(task dto.Task) (dto.Task, error) {
	newTask := dto.FromRequest(task)
	_, err := h.Usecase.CreateTask(h.ctx, newTask)
	if err != nil {
		h.log.Error("failed to create task", zap.Error(err))
		return dto.Task{}, err
	}

	return dto.Task{}, nil
}

func (h *TaskHandler) GetAllTasks() ([]dto.Task, error) {
	taskList, err := h.Usecase.GetAllTasks(h.ctx)
	if err != nil {
		h.log.Error("failed to get all tasks", zap.Error(err))
		return nil, err
	}
	convertedTaskList := []dto.Task{}

	for _, task := range taskList {
		convertedTaskList = append(convertedTaskList, dto.ToResponse(&task))
	}

	return convertedTaskList, nil
}

func (h *TaskHandler) ToggleTask(task dto.Task) error {
	err := h.Usecase.ToggleTask(h.ctx, dto.FromRequest(task))
	if err != nil {
		h.log.Error("failed to toggle task", zap.Error(err))
	}
	return nil
}

func (h *TaskHandler) DeleteTask(task dto.Task) error {
	err := h.Usecase.DeleteTask(h.ctx, dto.FromRequest(task))
	if err != nil {
		h.log.Error("failed to delete task", zap.Error(err))
	}

	return nil
}
