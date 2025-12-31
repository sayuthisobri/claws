package agents

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock-agent", "agents", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewAgentDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewAgentRenderer()
		},
	})
}
