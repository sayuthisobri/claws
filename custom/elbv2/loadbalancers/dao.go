package loadbalancers

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"

	appaws "github.com/clawscli/claws/internal/aws"
	"github.com/clawscli/claws/internal/dao"
	apperrors "github.com/clawscli/claws/internal/errors"
)

// LoadBalancerDAO provides data access for ELBv2 Load Balancers
type LoadBalancerDAO struct {
	dao.BaseDAO
	client *elasticloadbalancingv2.Client
}

// NewLoadBalancerDAO creates a new LoadBalancerDAO
func NewLoadBalancerDAO(ctx context.Context) (dao.DAO, error) {
	cfg, err := appaws.NewConfig(ctx)
	if err != nil {
		return nil, apperrors.Wrap(err, "new elbv2/loadbalancers dao")
	}
	return &LoadBalancerDAO{
		BaseDAO: dao.NewBaseDAO("elbv2", "load-balancers"),
		client:  elasticloadbalancingv2.NewFromConfig(cfg),
	}, nil
}

// List returns all load balancers (optionally filtered by ARN or name)
func (d *LoadBalancerDAO) List(ctx context.Context) ([]dao.Resource, error) {
	// Check for filters
	lbArn := dao.GetFilterFromContext(ctx, "LoadBalancerArn")
	lbName := dao.GetFilterFromContext(ctx, "LoadBalancerName")

	// If filtering by ARN or name, no pagination needed
	if lbArn != "" || lbName != "" {
		input := &elasticloadbalancingv2.DescribeLoadBalancersInput{}
		if lbArn != "" {
			input.LoadBalancerArns = []string{lbArn}
		}
		if lbName != "" {
			input.Names = []string{lbName}
		}

		output, err := d.client.DescribeLoadBalancers(ctx, input)
		if err != nil {
			return nil, apperrors.Wrap(err, "list load balancers")
		}

		resources := make([]dao.Resource, 0, len(output.LoadBalancers))
		for _, lb := range output.LoadBalancers {
			resources = append(resources, NewLoadBalancerResource(lb))
		}
		return resources, nil
	}

	// No filter - paginate through all
	loadBalancers, err := appaws.Paginate(ctx, func(token *string) ([]types.LoadBalancer, *string, error) {
		output, err := d.client.DescribeLoadBalancers(ctx, &elasticloadbalancingv2.DescribeLoadBalancersInput{
			Marker: token,
		})
		if err != nil {
			return nil, nil, apperrors.Wrap(err, "list load balancers")
		}
		return output.LoadBalancers, output.NextMarker, nil
	})
	if err != nil {
		return nil, err
	}

	resources := make([]dao.Resource, 0, len(loadBalancers))
	for _, lb := range loadBalancers {
		resources = append(resources, NewLoadBalancerResource(lb))
	}
	return resources, nil
}

// Get returns a specific load balancer
func (d *LoadBalancerDAO) Get(ctx context.Context, id string) (dao.Resource, error) {
	output, err := d.client.DescribeLoadBalancers(ctx, &elasticloadbalancingv2.DescribeLoadBalancersInput{
		LoadBalancerArns: []string{id},
	})
	if err != nil {
		return nil, apperrors.Wrapf(err, "get load balancer %s", id)
	}

	if len(output.LoadBalancers) == 0 {
		return nil, fmt.Errorf("load balancer not found: %s", id)
	}

	return NewLoadBalancerResource(output.LoadBalancers[0]), nil
}

// Delete deletes a load balancer
func (d *LoadBalancerDAO) Delete(ctx context.Context, id string) error {
	_, err := d.client.DeleteLoadBalancer(ctx, &elasticloadbalancingv2.DeleteLoadBalancerInput{
		LoadBalancerArn: &id,
	})
	if err != nil {
		if apperrors.IsNotFound(err) {
			return nil // Already deleted
		}
		if apperrors.IsResourceInUse(err) {
			return apperrors.Wrapf(err, "load balancer %s is in use", id)
		}
		return apperrors.Wrapf(err, "delete load balancer %s", id)
	}
	return nil
}

// LoadBalancerResource wraps an ELBv2 Load Balancer
type LoadBalancerResource struct {
	dao.BaseResource
	Item types.LoadBalancer
}

// NewLoadBalancerResource creates a new LoadBalancerResource
func NewLoadBalancerResource(lb types.LoadBalancer) *LoadBalancerResource {
	name := appaws.Str(lb.LoadBalancerName)
	arn := appaws.Str(lb.LoadBalancerArn)

	return &LoadBalancerResource{
		BaseResource: dao.BaseResource{
			ID:   arn,
			Name: name,
			ARN:  arn,
			Tags: make(map[string]string),
			Data: lb,
		},
		Item: lb,
	}
}

// LoadBalancerName returns the load balancer name
func (r *LoadBalancerResource) LoadBalancerName() string {
	if r.Item.LoadBalancerName != nil {
		return *r.Item.LoadBalancerName
	}
	return ""
}

// LoadBalancerArn returns the load balancer ARN
func (r *LoadBalancerResource) LoadBalancerArn() string {
	if r.Item.LoadBalancerArn != nil {
		return *r.Item.LoadBalancerArn
	}
	return ""
}

// Type returns the load balancer type (application, network, gateway)
func (r *LoadBalancerResource) Type() string {
	return string(r.Item.Type)
}

// Scheme returns the scheme (internet-facing, internal)
func (r *LoadBalancerResource) Scheme() string {
	return string(r.Item.Scheme)
}

// State returns the load balancer state
func (r *LoadBalancerResource) State() string {
	if r.Item.State != nil {
		return string(r.Item.State.Code)
	}
	return ""
}

// StateReason returns the state reason if any
func (r *LoadBalancerResource) StateReason() string {
	if r.Item.State != nil && r.Item.State.Reason != nil {
		return *r.Item.State.Reason
	}
	return ""
}

// DNSName returns the DNS name
func (r *LoadBalancerResource) DNSName() string {
	if r.Item.DNSName != nil {
		return *r.Item.DNSName
	}
	return ""
}

// VpcId returns the VPC ID
func (r *LoadBalancerResource) VpcId() string {
	if r.Item.VpcId != nil {
		return *r.Item.VpcId
	}
	return ""
}

// CreatedTime returns the creation time
func (r *LoadBalancerResource) CreatedTime() time.Time {
	if r.Item.CreatedTime != nil {
		return *r.Item.CreatedTime
	}
	return time.Time{}
}

// IpAddressType returns the IP address type
func (r *LoadBalancerResource) IpAddressType() string {
	return string(r.Item.IpAddressType)
}

// CanonicalHostedZoneId returns the Route53 hosted zone ID
func (r *LoadBalancerResource) CanonicalHostedZoneId() string {
	if r.Item.CanonicalHostedZoneId != nil {
		return *r.Item.CanonicalHostedZoneId
	}
	return ""
}

// AvailabilityZones returns the availability zones
func (r *LoadBalancerResource) AvailabilityZones() []string {
	var zones []string
	for _, az := range r.Item.AvailabilityZones {
		if az.ZoneName != nil {
			zones = append(zones, *az.ZoneName)
		}
	}
	return zones
}

// SecurityGroups returns the security group IDs
func (r *LoadBalancerResource) SecurityGroups() []string {
	return r.Item.SecurityGroups
}

// IPAddresses returns the IP addresses of the load balancer
func (r *LoadBalancerResource) IPAddresses() []string {
	var ips []string

	// First try to get IP addresses from load balancer addresses
	for _, az := range r.Item.AvailabilityZones {
		for _, addr := range az.LoadBalancerAddresses {
			if addr.IpAddress != nil {
				ips = append(ips, *addr.IpAddress)
			}
			if addr.IPv6Address != nil {
				ips = append(ips, *addr.IPv6Address)
			}
			if addr.PrivateIPv4Address != nil {
				ips = append(ips, *addr.PrivateIPv4Address)
			}
		}
	}

	// If no IP addresses found from load balancer addresses, try DNS lookup
	if len(ips) == 0 {
		dnsName := r.DNSName()
		if dnsName != "" {
			// Perform DNS lookup to get IP addresses
			addrs, err := net.LookupIP(dnsName)
			if err == nil && len(addrs) > 0 {
				for _, addr := range addrs {
					ips = append(ips, addr.String())
				}
			}
		}
	}

	return ips
}

// ResourceMap represents the relationship between listeners, target groups, and targets
type ResourceMap struct {
	Listeners []ListenerMap `json:"listeners"`
}

// ListenerMap represents a listener and its associated target groups and targets
type ListenerMap struct {
	Protocol string            `json:"protocol"`
	Port     int32             `json:"port"`
	Actions  []ActionMap       `json:"actions"`
}

// ActionMap represents an action and its associated target groups
type ActionMap struct {
	Type          string            `json:"type"`
	TargetGroups  []TargetGroupMap  `json:"targetGroups"`
	Description   string            `json:"description,omitempty"`
	// Fixed Response Action Details
	FixedResponse *FixedResponseDetails `json:"fixedResponse,omitempty"`
	// Redirect Action Details
	Redirect *RedirectDetails `json:"redirect,omitempty"`
}

// FixedResponseDetails represents fixed response action configuration
type FixedResponseDetails struct {
	StatusCode  string `json:"statusCode"`
	ContentType string `json:"contentType,omitempty"`
	MessageBody string `json:"messageBody,omitempty"`
}

// RedirectDetails represents redirect action configuration
type RedirectDetails struct {
	StatusCode string `json:"statusCode"`
	Host       string `json:"host,omitempty"`
	Path       string `json:"path,omitempty"`
	Port       string `json:"port,omitempty"`
	Protocol   string `json:"protocol,omitempty"`
	Query      string `json:"query,omitempty"`
}

// TargetGroupMap represents a target group and its targets
type TargetGroupMap struct {
	Name    string      `json:"name"`
	Arn     string      `json:"arn"`
	Targets []TargetMap `json:"targets"`
}

// TargetMap represents a target with its health status
type TargetMap struct {
	ID           string `json:"id"`
	Port         int32  `json:"port"`
	HealthState  string `json:"healthState"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
}

// GetResourceMap returns the complete resource map for a load balancer
func (r *LoadBalancerResource) GetResourceMap(ctx context.Context) (*ResourceMap, error) {
	// Create AWS config and client
	cfg, err := appaws.NewConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS config: %w", err)
	}
	client := elasticloadbalancingv2.NewFromConfig(cfg)

	// Get listeners for this load balancer
	listeners, err := client.DescribeListeners(ctx, &elasticloadbalancingv2.DescribeListenersInput{
		LoadBalancerArn: r.Item.LoadBalancerArn,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get listeners: %w", err)
	}

	resourceMap := &ResourceMap{
		Listeners: make([]ListenerMap, 0, len(listeners.Listeners)),
	}

	for _, listener := range listeners.Listeners {
		listenerMap := ListenerMap{
			Protocol: string(listener.Protocol),
			Port:     *listener.Port,
			Actions:  make([]ActionMap, 0, len(listener.DefaultActions)),
		}

		for _, action := range listener.DefaultActions {
			actionMap := ActionMap{
				Type: string(action.Type),
			}

			switch action.Type {
			case types.ActionTypeEnumForward:
				if action.ForwardConfig != nil && action.ForwardConfig.TargetGroups != nil {
					actionMap.TargetGroups = make([]TargetGroupMap, 0, len(action.ForwardConfig.TargetGroups))

					for _, tg := range action.ForwardConfig.TargetGroups {
						if tg.TargetGroupArn != nil {
							// Get target group details
							tgOutput, err := client.DescribeTargetGroups(ctx, &elasticloadbalancingv2.DescribeTargetGroupsInput{
								TargetGroupArns: []string{*tg.TargetGroupArn},
							})
							if err != nil {
								actionMap.Description = fmt.Sprintf("Error getting target group: %v", err)
								continue
							}

							var tgName string
							if len(tgOutput.TargetGroups) > 0 && tgOutput.TargetGroups[0].TargetGroupName != nil {
								tgName = *tgOutput.TargetGroups[0].TargetGroupName
							}

							// Get targets for this target group
							thOutput, err := client.DescribeTargetHealth(ctx, &elasticloadbalancingv2.DescribeTargetHealthInput{
								TargetGroupArn: tg.TargetGroupArn,
							})

							targets := make([]TargetMap, 0)
							if err == nil && len(thOutput.TargetHealthDescriptions) > 0 {
								for _, th := range thOutput.TargetHealthDescriptions {
									if th.Target != nil && th.Target.Id != nil {
										target := TargetMap{
											ID: *th.Target.Id,
											Port: 0,
											HealthState: string(th.TargetHealth.State),
										}
										
										if th.Target.Port != nil {
											target.Port = *th.Target.Port
										}
										
										if th.Target.AvailabilityZone != nil {
											target.AvailabilityZone = *th.Target.AvailabilityZone
										}
										
										targets = append(targets, target)
									}
								}
							}

							targetGroupMap := TargetGroupMap{
								Name:    tgName,
								Arn:     *tg.TargetGroupArn,
								Targets: targets,
							}
							actionMap.TargetGroups = append(actionMap.TargetGroups, targetGroupMap)
						}
					}
				}
			case types.ActionTypeEnumFixedResponse:
				if action.FixedResponseConfig != nil {
					actionMap.FixedResponse = &FixedResponseDetails{
						StatusCode:  appaws.Str(action.FixedResponseConfig.StatusCode),
						ContentType: appaws.Str(action.FixedResponseConfig.ContentType),
						MessageBody: appaws.Str(action.FixedResponseConfig.MessageBody),
					}
				}
			case types.ActionTypeEnumRedirect:
				if action.RedirectConfig != nil {
					actionMap.Redirect = &RedirectDetails{
						StatusCode: string(action.RedirectConfig.StatusCode),
						Host:       appaws.Str(action.RedirectConfig.Host),
						Path:       appaws.Str(action.RedirectConfig.Path),
						Port:       appaws.Str(action.RedirectConfig.Port),
						Protocol:   appaws.Str(action.RedirectConfig.Protocol),
						Query:      appaws.Str(action.RedirectConfig.Query),
					}
				}
			default:
				actionMap.Description = fmt.Sprintf("Action: %s", action.Type)
			}

			listenerMap.Actions = append(listenerMap.Actions, actionMap)
		}

		resourceMap.Listeners = append(resourceMap.Listeners, listenerMap)
	}

	return resourceMap, nil
}
