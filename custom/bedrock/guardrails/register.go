package guardrails

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock", "guardrails", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewGuardrailDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewGuardrailRenderer()
		},
	})
}
