package restorejobs

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("backup", "restore-jobs", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewRestoreJobDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewRestoreJobRenderer()
		},
	})
}
