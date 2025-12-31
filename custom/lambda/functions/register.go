package functions

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("lambda", "functions", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewFunctionDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewFunctionRenderer()
		},
	})
}
