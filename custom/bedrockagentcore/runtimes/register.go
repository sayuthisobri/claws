package runtimes

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock-agentcore", "runtimes", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRuntimeDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRuntimeRenderer()
		},
	})
}
