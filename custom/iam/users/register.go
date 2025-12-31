package users

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("iam", "users", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewUserDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewUserRenderer()
		},
	})
}
