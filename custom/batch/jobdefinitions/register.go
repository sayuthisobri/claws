package jobdefinitions

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("batch", "job-definitions", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewJobDefinitionDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewJobDefinitionRenderer()
		},
	})
}
