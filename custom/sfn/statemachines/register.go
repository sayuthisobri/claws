package statemachines

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("sfn", "state-machines", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewStateMachineDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewStateMachineRenderer()
		},
	})
}
