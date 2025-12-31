package costs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("costexplorer", "costs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewCostDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewCostRenderer()
		},
	})
}
