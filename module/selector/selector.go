package selector

import (
	"github.com/micro/go-micro/client/selector"
	"github.com/micro/go-micro/registry"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// Built in random hashed node selector
type firstNodeSelector struct {
	opts selector.Options
}

func (n *firstNodeSelector) Init(opts ...selector.Option) error {
	for _, o := range opts {
		o(&n.opts)
	}
	return nil
}

func (n *firstNodeSelector) Options() selector.Options {
	return n.opts
}

func (n *firstNodeSelector) Select(service string, opts ...selector.SelectOption) (selector.Next, error) {
	services, err := n.opts.Registry.GetService(service)

	if err != nil {
		return nil, err
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	var sopts selector.SelectOptions
	for _, opt := range opts {
		opt(&sopts)
	}

	for _, filter := range sopts.Filters {
		services = filter(services)
	}

	if len(services) == 0 {
		return nil, selector.ErrNotFound
	}

	if len(services[0].Nodes) == 0 {
		return nil, selector.ErrNotFound
	}
	//TODO 游戏重连需要连接上次的服务器
	newNode := &registry.Node{
		Id:       services[0].Nodes[0].Id,
		Address:  services[0].Nodes[0].Address,
		Metadata: services[0].Nodes[0].Metadata,
	}
	return func() (*registry.Node, error) {
		return newNode, nil
	}, nil
}

func (n *firstNodeSelector) Mark(service string, node *registry.Node, err error) {
	return
}

func (n *firstNodeSelector) Reset(service string) {
	return
}

func (n *firstNodeSelector) Close() error {
	return nil
}

func (n *firstNodeSelector) String() string {
	return "first"
}

// Return a new first node selector
func FirstNodeSelector(opts ...selector.Option) selector.Selector {
	var sopts selector.Options
	for _, opt := range opts {
		opt(&sopts)
	}
	if sopts.Registry == nil {
		sopts.Registry = registry.DefaultRegistry
	}
	return &firstNodeSelector{sopts}
}
