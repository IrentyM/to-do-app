package handler

import (
	"context"
	"to-do-app/internal/adapters/controller/service/handler/dto"
	"to-do-app/internal/model"
)

type TaskHandler struct {
	Usecase TaskUsecase
}

func NewTaskHandler(uc TaskUsecase) *TaskHandler {
	return &TaskHandler{
		Usecase: uc,
	}
}

func (h *TaskHandler) CreateTask(ctx context.Context, task dto.Task) (dto.Task, error) {
	newTask := dto.FromRequest(task)
	h.Usecase.CreateTask(ctx, newTask)
	
	return dto.Task{}, nil
}

func (h *TaskHandler) GetAllTasks(ctx context.Context) ([]*model.Task, error) {
	return h.Usecase.GetAllTasks(ctx)
}

func (h *TaskHandler) DeleteTask(ctx context.Context, task *model.Task) {
	h.Usecase.DeleteTask(ctx, task)
}
