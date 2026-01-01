package listeners

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/clawscli/claws/internal/registry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListenerResource(t *testing.T) {
	// Test NewListenerResource
	listener := types.Listener{
		ListenerArn:     aws.String("arn:aws:elasticloadbalancing:us-east-1:123456789012:listener/app/my-load-balancer/50dc6c495c0c9188/f2f7dc8efc522ab2"),
		LoadBalancerArn: aws.String("arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188"),
		Protocol:        types.ProtocolEnumHttp,
		Port:            aws.Int32(80),
		DefaultActions:  []types.Action{{Type: types.ActionTypeEnumForward}},
		Certificates:    []types.Certificate{{CertificateArn: aws.String("arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012")}},
		SslPolicy:       aws.String("ELBSecurityPolicy-2016-08"),
		AlpnPolicy:      []string{"HTTP1Only"},
	}

	resource := NewListenerResource(listener)
	require.NotNil(t, resource)

	// Test getters
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-east-1:123456789012:listener/app/my-load-balancer/50dc6c495c0c9188/f2f7dc8efc522ab2", resource.ListenerArn())
	assert.Equal(t, "arn:aws:elasticloadbalancing:us-east-1:123456789012:loadbalancer/app/my-load-balancer/50dc6c495c0c9188", resource.LoadBalancerArn())
	assert.Equal(t, "HTTP", resource.Protocol())
	assert.Equal(t, int32(80), resource.Port())
	assert.Equal(t, "HTTP:80", resource.ProtocolPort())
	assert.Equal(t, 1, len(resource.DefaultActions()))
	assert.Equal(t, 1, len(resource.Certificates()))
	assert.Equal(t, "ELBSecurityPolicy-2016-08", resource.SslPolicy())
	assert.Equal(t, []string{"HTTP1Only"}, resource.AlpnPolicy())
}

func TestListenerDAO(t *testing.T) {
	// Test that the DAO can be created
	ctx := context.Background()
	dao, err := NewListenerDAO(ctx)
	require.NoError(t, err)
	require.NotNil(t, dao)
}

func TestListenerRenderer(t *testing.T) {
	// Test that the renderer can be created
	renderer := NewListenerRenderer()
	require.NotNil(t, renderer)
}

func TestListenerRegistry(t *testing.T) {
	// Test that the listener resource is registered
	entry, exists := registry.Global.Get("elbv2", "listeners")
	require.True(t, exists)
	require.NotNil(t, entry)

	// Test that the DAO factory works
	ctx := context.Background()
	dao, err := entry.DAOFactory(ctx)
	require.NoError(t, err)
	require.NotNil(t, dao)

	// Test that the renderer factory works
	renderer := entry.RendererFactory()
	require.NotNil(t, renderer)
}
