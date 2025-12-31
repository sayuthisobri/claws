package jobqueues

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("batch", "job-queues", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewJobQueueDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewJobQueueRenderer()
		},
	})
}
