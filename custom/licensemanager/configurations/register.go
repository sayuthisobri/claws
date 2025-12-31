package configurations

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("license-manager", "configurations", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewConfigurationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewConfigurationRenderer()
		},
	})
}
