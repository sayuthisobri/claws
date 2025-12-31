package inferenceprofiles

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock", "inference-profiles", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewInferenceProfileDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewInferenceProfileRenderer()
		},
	})
}
