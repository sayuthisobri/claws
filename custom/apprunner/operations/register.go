package operations

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("apprunner", "operations", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewOperationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewOperationRenderer()
		},
	})
}
