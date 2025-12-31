package grants

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("license-manager", "grants", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewGrantDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewGrantRenderer()
		},
	})
}
