package indexes

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("s3vectors", "indexes", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewVectorIndexDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewVectorIndexRenderer()
		},
	})
}
