package recommendations

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("computeoptimizer", "recommendations", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRecommendationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRecommendationRenderer()
		},
	})
}
