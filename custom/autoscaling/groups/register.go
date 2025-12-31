package groups

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("autoscaling", "groups", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewAutoScalingGroupDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewAutoScalingGroupRenderer()
		},
	})
}
