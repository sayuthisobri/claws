package taskexecutions

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("datasync", "task-executions", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewTaskExecutionDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewTaskExecutionRenderer()
		},
	})
}
