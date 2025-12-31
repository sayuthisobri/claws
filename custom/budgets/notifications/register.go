package notifications

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("budgets", "notifications", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewNotificationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewNotificationRenderer()
		},
	})
}
