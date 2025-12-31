package buckets

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("s3", "buckets", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewBucketDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewBucketRenderer()
		},
	})
}
