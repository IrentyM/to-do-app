package handler

import (
	"context"
	"to-do-app/internal/adapters/controller/service/handler/dto"
	"to-do-app/internal/model"
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
	h.Usecase.CreateTask(h.ctx, newTask)

	return dto.Task{}, nil
}

func (h *TaskHandler) GetAllTasks() ([]*model.Task, error) {
	return h.Usecase.GetAllTasks(h.ctx)
}

func (h *TaskHandler) DeleteTask(task *model.Task) {
	h.Usecase.DeleteTask(h.ctx, task)
}
