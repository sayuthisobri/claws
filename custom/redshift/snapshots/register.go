package snapshots

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("redshift", "snapshots", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewSnapshotDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewSnapshotRenderer()
		},
	})
}
