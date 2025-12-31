package budgets

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("budgets", "budgets", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewBudgetDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewBudgetRenderer()
		},
	})
}
