package knowledgebases

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("bedrock-agent", "knowledge-bases", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewKnowledgeBaseDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewKnowledgeBaseRenderer()
		},
	})
}
