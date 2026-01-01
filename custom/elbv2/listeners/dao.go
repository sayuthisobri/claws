package listeners

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"

	appaws "github.com/clawscli/claws/internal/aws"
	"github.com/clawscli/claws/internal/dao"
	apperrors "github.com/clawscli/claws/internal/errors"
)

// ListenerDAO provides data access for ELBv2 Listeners
type ListenerDAO struct {
	dao.BaseDAO
	client *elasticloadbalancingv2.Client
}

// NewListenerDAO creates a new ListenerDAO
func NewListenerDAO(ctx context.Context) (dao.DAO, error) {
	cfg, err := appaws.NewConfig(ctx)
	if err != nil {
		return nil, apperrors.Wrap(err, "new elbv2/listeners dao")
	}
	return &ListenerDAO{
		BaseDAO: dao.NewBaseDAO("elbv2", "listeners"),
		client:  elasticloadbalancingv2.NewFromConfig(cfg),
	}, nil
}

// List returns all listeners (optionally filtered by load balancer ARN or listener ARN)
func (d *ListenerDAO) List(ctx context.Context) ([]dao.Resource, error) {
	// Check for filters
	lbArn := dao.GetFilterFromContext(ctx, "LoadBalancerArn")
	listenerArn := dao.GetFilterFromContext(ctx, "ListenerArn")

	// If filtering by listener ARN, no pagination needed
	if listenerArn != "" {
		output, err := d.client.DescribeListeners(ctx, &elasticloadbalancingv2.DescribeListenersInput{
			ListenerArns: []string{listenerArn},
		})
		if err != nil {
			return nil, apperrors.Wrap(err, "list listeners")
		}
		resources := make([]dao.Resource, 0, len(output.Listeners))
		for _, listener := range output.Listeners {
			resources = append(resources, NewListenerResource(listener))
		}
		return resources, nil
	}

	// Paginate through results
	listeners, err := appaws.Paginate(ctx, func(token *string) ([]types.Listener, *string, error) {
		input := &elasticloadbalancingv2.DescribeListenersInput{
			Marker: token,
		}
		if lbArn != "" {
			input.LoadBalancerArn = &lbArn
		}
		output, err := d.client.DescribeListeners(ctx, input)
		if err != nil {
			return nil, nil, apperrors.Wrap(err, "list listeners")
		}
		return output.Listeners, output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	resources := make([]dao.Resource, 0, len(listeners))
	for _, listener := range listeners {
		resources = append(resources, NewListenerResource(listener))
	}
	return resources, nil
}

// Get returns a specific listener
func (d *ListenerDAO) Get(ctx context.Context, id string) (dao.Resource, error) {
	output, err := d.client.DescribeListeners(ctx, &elasticloadbalancingv2.DescribeListenersInput{
		ListenerArns: []string{id},
	})
	if err != nil {
		return nil, apperrors.Wrapf(err, "get listener %s", id)
	}

	if len(output.Listeners) == 0 {
		return nil, fmt.Errorf("listener not found: %s", id)
	}

	return NewListenerResource(output.Listeners[0]), nil
}

// Delete deletes a listener
func (d *ListenerDAO) Delete(ctx context.Context, id string) error {
	_, err := d.client.DeleteListener(ctx, &elasticloadbalancingv2.DeleteListenerInput{
		ListenerArn: &id,
	})
	if err != nil {
		if apperrors.IsNotFound(err) {
			return nil // Already deleted
		}
		if apperrors.IsResourceInUse(err) {
			return apperrors.Wrapf(err, "listener %s is in use", id)
		}
		return apperrors.Wrapf(err, "delete listener %s", id)
	}
	return nil
}

// ListenerResource wraps an ELBv2 Listener
type ListenerResource struct {
	dao.BaseResource
	Item types.Listener
}

// NewListenerResource creates a new ListenerResource
func NewListenerResource(listener types.Listener) *ListenerResource {
	arn := appaws.Str(listener.ListenerArn)

	return &ListenerResource{
		BaseResource: dao.BaseResource{
			ID:   arn,
			Name: fmt.Sprintf("%s:%d", strings.ToLower(string(listener.Protocol)), appaws.Int32(listener.Port)),
			ARN:  arn,
			Tags: make(map[string]string),
			Data: listener,
		},
		Item: listener,
	}
}

// ListenerArn returns the listener ARN
func (r *ListenerResource) ListenerArn() string {
	if r.Item.ListenerArn != nil {
		return *r.Item.ListenerArn
	}
	return ""
}

// LoadBalancerArn returns the load balancer ARN
func (r *ListenerResource) LoadBalancerArn() string {
	if r.Item.LoadBalancerArn != nil {
		return *r.Item.LoadBalancerArn
	}
	return ""
}

// Protocol returns the protocol
func (r *ListenerResource) Protocol() string {
	return string(r.Item.Protocol)
}

// Port returns the port
func (r *ListenerResource) Port() int32 {
	if r.Item.Port != nil {
		return *r.Item.Port
	}
	return 0
}

// ProtocolPort returns protocol:port string
func (r *ListenerResource) ProtocolPort() string {
	if r.Item.Port != nil {
		return fmt.Sprintf("%s:%d", r.Protocol(), *r.Item.Port)
	}
	return r.Protocol()
}

// DefaultActions returns the default actions
func (r *ListenerResource) DefaultActions() []types.Action {
	return r.Item.DefaultActions
}

// Certificates returns the certificates
func (r *ListenerResource) Certificates() []types.Certificate {
	return r.Item.Certificates
}

// SslPolicy returns the SSL policy
func (r *ListenerResource) SslPolicy() string {
	if r.Item.SslPolicy != nil {
		return *r.Item.SslPolicy
	}
	return ""
}

// AlpnPolicy returns the ALPN policy
func (r *ListenerResource) AlpnPolicy() []string {
	return r.Item.AlpnPolicy
}

// MutualAuthentication returns the mutual authentication configuration
func (r *ListenerResource) MutualAuthentication() *types.MutualAuthenticationAttributes {
	return r.Item.MutualAuthentication
}
