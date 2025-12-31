package monitors

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("costexplorer", "monitors", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewMonitorDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewMonitorRenderer()
		},
	})
}
