package profile

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("local", "profile", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewProfileDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewProfileRenderer()
		},
	})
}
