package activities

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("autoscaling", "activities", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewActivityDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewActivityRenderer()
		},
	})
}
