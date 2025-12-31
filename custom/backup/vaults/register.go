package vaults

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("backup", "vaults", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewVaultDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewVaultRenderer()
		},
	})
}
