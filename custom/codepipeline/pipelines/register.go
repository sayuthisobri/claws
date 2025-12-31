package pipelines

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("codepipeline", "pipelines", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewPipelineDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewPipelineRenderer()
		},
	})
}
