package graphqlapis

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("appsync", "graphql-apis", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewGraphQLApiDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewGraphQLApiRenderer()
		},
	})
}
