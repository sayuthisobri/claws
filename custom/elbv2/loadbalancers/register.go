package loadbalancers

import (
	"context"

	"github.com/sayuthisobri/claws/internal/dao"
	"github.com/sayuthisobri/claws/internal/registry"
	"github.com/sayuthisobri/claws/internal/render"
)

func init() {
	registry.Global.RegisterCustom("elbv2", "load-balancers", registry.Entry{
		DAOFactory: func(ctx context.Context) (dao.DAO, error) {
			return NewLoadBalancerDAO(ctx)
		},
		RendererFactory: func() render.Renderer {
			return NewLoadBalancerRenderer()
		},
	})
}
