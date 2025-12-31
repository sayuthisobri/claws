package projects

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("codebuild", "projects", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewProjectDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewProjectRenderer()
		},
	})
}
