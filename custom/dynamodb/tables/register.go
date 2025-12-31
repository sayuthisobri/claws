package tables

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("dynamodb", "tables", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewTableDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewTableRenderer()
		},
	})
}
