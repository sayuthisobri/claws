package images

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("ec2", "images", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewImageDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewImageRenderer()
		},
	})
}
