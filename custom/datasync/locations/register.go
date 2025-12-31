package locations

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("datasync", "locations", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewLocationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewLocationRenderer()
		},
	})
}
