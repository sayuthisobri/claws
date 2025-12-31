package rules

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("config", "rules", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRuleDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRuleRenderer()
		},
	})
}
