package notebooks

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("sagemaker", "notebooks", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewNotebookDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewNotebookRenderer()
		},
	})
}
