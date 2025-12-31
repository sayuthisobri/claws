package secrets

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("secretsmanager", "secrets", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewSecretDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewSecretRenderer()
		},
	})
}
