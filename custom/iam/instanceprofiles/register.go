package instanceprofiles

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("iam", "instance-profiles", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewInstanceProfileDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewInstanceProfileRenderer()
		},
	})
}
