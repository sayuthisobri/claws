package servers

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("transfer", "servers", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewServerDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewServerRenderer()
		},
	})
}
