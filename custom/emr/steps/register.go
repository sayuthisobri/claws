package steps

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("emr", "steps", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewStepDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewStepRenderer()
		},
	})
}
