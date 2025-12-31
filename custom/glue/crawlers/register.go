package crawlers

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("glue", "crawlers", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewCrawlerDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewCrawlerRenderer()
		},
	})
}
