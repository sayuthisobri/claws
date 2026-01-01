package listeners

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/clawscli/claws/internal/dao"
	"github.com/clawscli/claws/internal/render"
)

// ListenerRenderer renders ELBv2 Listeners
// Ensure ListenerRenderer implements render.Navigator
var _ render.Navigator = (*ListenerRenderer)(nil)

type ListenerRenderer struct {
	render.BaseRenderer
}

// NewListenerRenderer creates a new ListenerRenderer
func NewListenerRenderer() render.Renderer {
	return &ListenerRenderer{
		BaseRenderer: render.BaseRenderer{
			Service:  "elbv2",
			Resource: "listeners",
			Cols: []render.Column{
				{
					Name:  "NAME",
					Width: 32,
					Getter: func(r dao.Resource) string {
						return r.GetName()
					},
					Priority: 0,
				},
				{
					Name:  "PROTOCOL",
					Width: 12,
					Getter: func(r dao.Resource) string {
						if rr, ok := r.(*ListenerResource); ok {
							return rr.Protocol()
						}
						return ""
					},
					Priority: 1,
				},
				{
					Name:  "PORT",
					Width: 8,
					Getter: func(r dao.Resource) string {
						if rr, ok := r.(*ListenerResource); ok {
							port := rr.Port()
							if port > 0 {
								return fmt.Sprintf("%d", port)
							}
						}
						return ""
					},
					Priority: 2,
				},
				{
					Name:  "LOAD BALANCER",
					Width: 40,
					Getter: func(r dao.Resource) string {
						if rr, ok := r.(*ListenerResource); ok {
							return rr.LoadBalancerArn()
						}
						return ""
					},
					Priority: 3,
				},
				{
					Name:  "DEFAULT ACTIONS",
					Width: 30,
					Getter: func(r dao.Resource) string {
						if rr, ok := r.(*ListenerResource); ok {
							actions := rr.DefaultActions()
							if len(actions) > 0 {
								var actionTypes []string
								for _, action := range actions {
									actionTypes = append(actionTypes, string(action.Type))
								}
								return strings.Join(actionTypes, ", ")
							}
						}
						return ""
					},
					Priority: 4,
				},
				{
					Name:  "SSL POLICY",
					Width: 20,
					Getter: func(r dao.Resource) string {
						if rr, ok := r.(*ListenerResource); ok {
							return rr.SslPolicy()
						}
						return ""
					},
					Priority: 5,
				},
			},
		},
	}
}

// RenderDetail renders detailed listener information
func (r *ListenerRenderer) RenderDetail(resource dao.Resource) string {
	rr, ok := resource.(*ListenerResource)
	if !ok {
		return ""
	}

	d := render.NewDetailBuilder()

	d.Title("Listener", rr.ListenerArn())

	// Basic Info
	d.Section("Basic Information")
	d.Field("Listener ARN", rr.ListenerArn())
	d.Field("Load Balancer ARN", rr.LoadBalancerArn())
	d.Field("Protocol", rr.Protocol())
	d.Field("Port", fmt.Sprintf("%d", rr.Port()))
	d.Field("Protocol:Port", rr.ProtocolPort())

	// Default Actions
	actions := rr.DefaultActions()
	if len(actions) > 0 {
		d.Section("Default Actions")
		for i, action := range actions {
			d.Field(fmt.Sprintf("Action %d Type", i+1), string(action.Type))
			if action.TargetGroupArn != nil {
				d.Field("Target Group ARN", *action.TargetGroupArn)
			}
			if action.Order != nil {
				d.Field("Order", fmt.Sprintf("%d", *action.Order))
			}
			if action.Type == types.ActionTypeEnumForward {
				if action.ForwardConfig != nil && action.ForwardConfig.TargetGroups != nil {
					for j, tg := range action.ForwardConfig.TargetGroups {
						if tg.TargetGroupArn != nil {
							d.Field(fmt.Sprintf("Target Group %d", j+1), *tg.TargetGroupArn)
							if tg.Weight != nil {
								d.Field(fmt.Sprintf("Weight %d", j+1), fmt.Sprintf("%d", *tg.Weight))
							}
						}
					}
				}
			}
		}
	}

	// Certificates
	certs := rr.Certificates()
	if len(certs) > 0 {
		d.Section("Certificates")
		for i, cert := range certs {
			if cert.CertificateArn != nil {
				d.Field(fmt.Sprintf("Certificate %d ARN", i+1), *cert.CertificateArn)
			}
			if cert.IsDefault != nil {
				d.Field(fmt.Sprintf("Certificate %d Is Default", i+1), fmt.Sprintf("%t", *cert.IsDefault))
			}
		}
	}

	// SSL Policy
	if sslPolicy := rr.SslPolicy(); sslPolicy != "" {
		d.Section("SSL Configuration")
		d.Field("SSL Policy", sslPolicy)
	}

	// ALPN Policy
	alpnPolicy := rr.AlpnPolicy()
	if len(alpnPolicy) > 0 {
		d.Section("ALPN Configuration")
		d.Field("ALPN Policy", strings.Join(alpnPolicy, ", "))
	}

	// Mutual Authentication
	if ma := rr.MutualAuthentication(); ma != nil {
		d.Section("Mutual Authentication")
		if ma.Mode != nil {
			d.Field("Mode", string(*ma.Mode))
		}
		if ma.TrustStoreArn != nil {
			d.Field("Trust Store ARN", *ma.TrustStoreArn)
		}
		if ma.IgnoreClientCertificateExpiry != nil {
			d.Field("Ignore Client Certificate Expiry", fmt.Sprintf("%t", *ma.IgnoreClientCertificateExpiry))
		}
	}

	return d.String()
}

// RenderSummary returns summary fields for the header panel
func (r *ListenerRenderer) RenderSummary(resource dao.Resource) []render.SummaryField {
	rr, ok := resource.(*ListenerResource)
	if !ok {
		return nil
	}

	return []render.SummaryField{
		{Label: "Protocol", Value: rr.Protocol()},
		{Label: "Port", Value: fmt.Sprintf("%d", rr.Port())},
		{Label: "Load Balancer", Value: rr.LoadBalancerArn()},
		{Label: "Default Actions", Value: fmt.Sprintf("%d actions", len(rr.DefaultActions()))},
	}
}

// Navigations returns available navigation options
func (r *ListenerRenderer) Navigations(resource dao.Resource) []render.Navigation {
	rr, ok := resource.(*ListenerResource)
	if !ok {
		return nil
	}

	navs := []render.Navigation{
		{
			Key:         "l",
			Label:       "Load Balancer",
			Service:     "elbv2",
			Resource:    "load-balancers",
			FilterField: "LoadBalancerArn",
			FilterValue: rr.LoadBalancerArn(),
		},
	}

	// Target groups navigation for forward actions
	actions := rr.DefaultActions()
	for _, action := range actions {
		if action.TargetGroupArn != nil && *action.TargetGroupArn != "" {
			navs = append(navs, render.Navigation{
				Key:         "t",
				Label:       "Target Group",
				Service:     "elbv2",
				Resource:    "target-groups",
				FilterField: "TargetGroupArn",
				FilterValue: *action.TargetGroupArn,
			})
			break // Only add one target group navigation
		}
	}

	// Certificates navigation for HTTPS/TLS listeners
	certs := rr.Certificates()
	for _, cert := range certs {
		if cert.CertificateArn != nil && *cert.CertificateArn != "" {
			navs = append(navs, render.Navigation{
				Key:         "C",
				Label:       "Certificate",
				Service:     "acm",
				Resource:    "certificates",
				FilterField: "CertificateArn",
				FilterValue: *cert.CertificateArn,
			})
			break // Only add one certificate navigation
		}
	}

	return navs
}
