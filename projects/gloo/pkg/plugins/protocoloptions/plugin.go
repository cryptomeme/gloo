package protocoloptions

import (
	envoy_config_cluster_v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	envoy_config_core_v3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"github.com/golang/protobuf/ptypes/wrappers"
	errors "github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
)

var (
	_ plugins.Plugin         = new(plugin)
	_ plugins.UpstreamPlugin = new(plugin)
)

const (
	ExtensionName = "protocol_options"
	MinWindowSize = 65535
	MaxWindowSize = 2147483647
)

type plugin struct{}

func NewPlugin() *plugin {
	return &plugin{}
}

func (p *plugin) Name() string {
	return ExtensionName
}

func (p *plugin) Init(params plugins.InitParams) error {
	return nil
}

func (p *plugin) ProcessUpstream(params plugins.Params, in *v1.Upstream, out *envoy_config_cluster_v3.Cluster) error {

	if in.GetUseHttp2() == nil || !in.GetUseHttp2().GetValue() {
		return nil
	}

	if out.GetHttp2ProtocolOptions() == nil {
		out.Http2ProtocolOptions = &envoy_config_core_v3.Http2ProtocolOptions{}
	}

	// Both these values default to 268435456 if unset.
	sws := in.GetInitialStreamWindowSize()
	if sws != nil {
		if validateWindowSize(sws.GetValue()) {
			out.GetHttp2ProtocolOptions().InitialStreamWindowSize = &wrappers.UInt32Value{Value: sws.GetValue()}
		} else {
			return errors.Errorf("Invalid Initial Steam Window Size: %d", sws.GetValue())
		}
	}

	cws := in.GetInitialConnectionWindowSize()
	if cws != nil {
		if validateWindowSize(cws.GetValue()) {
			out.GetHttp2ProtocolOptions().InitialConnectionWindowSize = &wrappers.UInt32Value{Value: cws.GetValue()}
		} else {
			return errors.Errorf("Invalid Initial Connection Window Size: %d", cws.GetValue())
		}
	}

	return nil
}

func validateWindowSize(size uint32) bool {
	if size < MinWindowSize || size > MaxWindowSize {
		return false
	}
	return true
}
