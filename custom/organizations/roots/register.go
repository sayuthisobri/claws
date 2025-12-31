package roots

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("organizations", "roots", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRootDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRootRenderer()
		},
	})
}
