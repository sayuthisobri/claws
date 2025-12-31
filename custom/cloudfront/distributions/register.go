package distributions

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("cloudfront", "distributions", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewDistributionDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewDistributionRenderer()
		},
	})
}
