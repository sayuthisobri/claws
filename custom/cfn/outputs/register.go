package outputs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("cloudformation", "outputs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewOutputDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewOutputRenderer()
		},
	})
}
