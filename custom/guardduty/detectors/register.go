package detectors

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("guardduty", "detectors", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewDetectorDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewDetectorRenderer()
		},
	})
}
