package summary

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("computeoptimizer", "summary", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewSummaryDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewSummaryRenderer()
		},
	})
}
