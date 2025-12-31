package roles

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("iam", "roles", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRoleDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRoleRenderer()
		},
	})
}
