package reservedinstances

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("risp", "reserved-instances", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewReservedInstanceDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewReservedInstanceRenderer()
		},
	})
}
