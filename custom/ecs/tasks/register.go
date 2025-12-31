package tasks

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("ecs", "tasks", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewTaskDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewTaskRenderer()
		},
	})
}
