package models

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("sagemaker", "models", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewModelDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewModelRenderer()
		},
	})
}
