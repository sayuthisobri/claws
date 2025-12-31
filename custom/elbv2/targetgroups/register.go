package targetgroups

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("elbv2", "target-groups", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewTargetGroupDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewTargetGroupRenderer()
		},
	})
}
