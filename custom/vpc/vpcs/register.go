package vpcs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("vpc", "vpcs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewVPCDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewVPCRenderer()
		},
	})
}
