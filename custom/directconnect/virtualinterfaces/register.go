package virtualinterfaces

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("directconnect", "virtual-interfaces", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewVirtualInterfaceDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewVirtualInterfaceRenderer()
		},
	})
}
