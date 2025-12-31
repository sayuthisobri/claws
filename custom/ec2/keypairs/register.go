package keypairs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("ec2", "key-pairs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewKeyPairDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewKeyPairRenderer()
		},
	})
}
