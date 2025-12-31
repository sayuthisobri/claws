package copyjobs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("backup", "copy-jobs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewCopyJobDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewCopyJobRenderer()
		},
	})
}
