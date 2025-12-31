package buckets

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("s3vectors", "buckets", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewVectorBucketDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewVectorBucketRenderer()
		},
	})
}
