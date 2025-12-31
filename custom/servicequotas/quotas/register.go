package quotas

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("service-quotas", "quotas", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewQuotaDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewQuotaRenderer()
		},
	})
}
