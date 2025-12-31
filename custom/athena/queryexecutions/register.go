package queryexecutions

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("athena", "query-executions", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewQueryExecutionDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewQueryExecutionRenderer()
		},
	})
}
