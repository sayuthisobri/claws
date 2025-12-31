package foundationmodels

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock", "foundation-models", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewFoundationModelDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewFoundationModelRenderer()
		},
	})
}
