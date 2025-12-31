package alarms

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("cloudwatch", "alarms", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewAlarmDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewAlarmRenderer()
		},
	})
}
