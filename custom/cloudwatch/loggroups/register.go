package loggroups

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("cloudwatch", "log-groups", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewLogGroupDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewLogGroupRenderer()
		},
	})
}
