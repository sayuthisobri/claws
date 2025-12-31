package anomalies

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("costexplorer", "anomalies", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewAnomalyDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewAnomalyRenderer()
		},
	})
}
