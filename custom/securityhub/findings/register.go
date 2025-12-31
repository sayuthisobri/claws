package findings

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("securityhub", "findings", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewFindingDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewFindingRenderer()
		},
	})
}
