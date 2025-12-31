package resources

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("cloudformation", "resources", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewResourceDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewResourceRenderer()
		},
	})
}
