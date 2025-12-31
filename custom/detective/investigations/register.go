package investigations

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("detective", "investigations", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewInvestigationDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewInvestigationRenderer()
		},
	})
}
