package handler

import (
	"context"
	"to-do-app/internal/adapters/controller/service/handler/dto"
)

type TaskHandler struct {
	ctx     context.Context
	Usecase TaskUsecase
}

func NewTaskHandler(uc TaskUsecase, ctx context.Context) *TaskHandler {
	return &TaskHandler{
		Usecase: uc,
		ctx:     ctx,
	}
}

func (h *TaskHandler) SetContext(ctx context.Context) {
	h.ctx = ctx
}

func (h *TaskHandler) CreateTask(task dto.Task) (dto.Task, error) {
	newTask := dto.FromRequest(task)
	_, err := h.Usecase.CreateTask(h.ctx, newTask)
	if err != nil {
		return dto.Task{}, err
	}

	return dto.Task{}, nil
}

func (h *TaskHandler) GetAllTasks() ([]dto.Task, error) {
	taskList, err := h.Usecase.GetAllTasks(h.ctx)
	if err != nil {
		return nil, err
	}
	convertedTaskList := []dto.Task{}

	for _, task := range taskList {
		convertedTaskList = append(convertedTaskList, dto.ToResponse(&task))
	}

	return convertedTaskList, nil
}

func (h *TaskHandler) ToggleTask(task dto.Task) error {
	return h.Usecase.ToggleTask(h.ctx, dto.FromRequest(task))
}

func (h *TaskHandler) DeleteTask(task dto.Task) error {
	return h.Usecase.DeleteTask(h.ctx, dto.FromRequest(task))
}
