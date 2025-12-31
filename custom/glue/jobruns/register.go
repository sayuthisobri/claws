package jobruns

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("glue", "job-runs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewJobRunDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewJobRunRenderer()
		},
	})
}
