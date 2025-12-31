package domains

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("opensearch", "domains", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewDomainDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewDomainRenderer()
		},
	})
}
