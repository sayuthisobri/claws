package accounts

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("organizations", "accounts", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewAccountDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewAccountRenderer()
		},
	})
}
