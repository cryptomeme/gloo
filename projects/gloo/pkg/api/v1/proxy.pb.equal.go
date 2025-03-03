// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/proxy.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *Proxy) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Proxy)
	if !ok {
		that2, ok := that.(Proxy)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetCompressedSpec(), target.GetCompressedSpec()) != 0 {
		return false
	}

	if len(m.GetListeners()) != len(target.GetListeners()) {
		return false
	}
	for idx, v := range m.GetListeners() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetListeners()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetListeners()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Listener) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Listener)
	if !ok {
		that2, ok := that.(Listener)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if strings.Compare(m.GetBindAddress(), target.GetBindAddress()) != 0 {
		return false
	}

	if m.GetBindPort() != target.GetBindPort() {
		return false
	}

	if len(m.GetSslConfigurations()) != len(target.GetSslConfigurations()) {
		return false
	}
	for idx, v := range m.GetSslConfigurations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSslConfigurations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSslConfigurations()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetUseProxyProto()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUseProxyProto()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUseProxyProto(), target.GetUseProxyProto()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetRouteOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRouteOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRouteOptions(), target.GetRouteOptions()) {
			return false
		}
	}

	switch m.ListenerType.(type) {

	case *Listener_HttpListener:
		if _, ok := target.ListenerType.(*Listener_HttpListener); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpListener()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpListener()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpListener(), target.GetHttpListener()) {
				return false
			}
		}

	case *Listener_TcpListener:
		if _, ok := target.ListenerType.(*Listener_TcpListener); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpListener()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpListener()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpListener(), target.GetTcpListener()) {
				return false
			}
		}

	case *Listener_HybridListener:
		if _, ok := target.ListenerType.(*Listener_HybridListener); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHybridListener()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHybridListener()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHybridListener(), target.GetHybridListener()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ListenerType != target.ListenerType {
			return false
		}
	}

	return true
}

// Equal function
func (m *TcpListener) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpListener)
	if !ok {
		that2, ok := that.(TcpListener)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetTcpHosts()) != len(target.GetTcpHosts()) {
		return false
	}
	for idx, v := range m.GetTcpHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTcpHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if strings.Compare(m.GetStatPrefix(), target.GetStatPrefix()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TcpHost) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpHost)
	if !ok {
		that2, ok := that.(TcpHost)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetDestination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestination(), target.GetDestination()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *HttpListener) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HttpListener)
	if !ok {
		that2, ok := that.(HttpListener)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetVirtualHosts()) != len(target.GetVirtualHosts()) {
		return false
	}
	for idx, v := range m.GetVirtualHosts() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetVirtualHosts()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetVirtualHosts()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if strings.Compare(m.GetStatPrefix(), target.GetStatPrefix()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *HybridListener) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*HybridListener)
	if !ok {
		that2, ok := that.(HybridListener)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMatchedListeners()) != len(target.GetMatchedListeners()) {
		return false
	}
	for idx, v := range m.GetMatchedListeners() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchedListeners()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchedListeners()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *MatchedListener) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MatchedListener)
	if !ok {
		that2, ok := that.(MatchedListener)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetMatcher()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMatcher()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMatcher(), target.GetMatcher()) {
			return false
		}
	}

	switch m.ListenerType.(type) {

	case *MatchedListener_HttpListener:
		if _, ok := target.ListenerType.(*MatchedListener_HttpListener); !ok {
			return false
		}

		if h, ok := interface{}(m.GetHttpListener()).(equality.Equalizer); ok {
			if !h.Equal(target.GetHttpListener()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetHttpListener(), target.GetHttpListener()) {
				return false
			}
		}

	case *MatchedListener_TcpListener:
		if _, ok := target.ListenerType.(*MatchedListener_TcpListener); !ok {
			return false
		}

		if h, ok := interface{}(m.GetTcpListener()).(equality.Equalizer); ok {
			if !h.Equal(target.GetTcpListener()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetTcpListener(), target.GetTcpListener()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.ListenerType != target.ListenerType {
			return false
		}
	}

	return true
}

// Equal function
func (m *Matcher) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Matcher)
	if !ok {
		that2, ok := that.(Matcher)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetSslConfig()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfig()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfig(), target.GetSslConfig()) {
			return false
		}
	}

	if len(m.GetSourcePrefixRanges()) != len(target.GetSourcePrefixRanges()) {
		return false
	}
	for idx, v := range m.GetSourcePrefixRanges() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetSourcePrefixRanges()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *VirtualHost) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*VirtualHost)
	if !ok {
		that2, ok := that.(VirtualHost)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	if len(m.GetDomains()) != len(target.GetDomains()) {
		return false
	}
	for idx, v := range m.GetDomains() {

		if strings.Compare(v, target.GetDomains()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetRoutes()) != len(target.GetRoutes()) {
		return false
	}
	for idx, v := range m.GetRoutes() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetRoutes()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetRoutes()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *Route) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Route)
	if !ok {
		that2, ok := that.(Route)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetMatchers()) != len(target.GetMatchers()) {
		return false
	}
	for idx, v := range m.GetMatchers() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetMatchers()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetMatchers()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	if strings.Compare(m.GetName(), target.GetName()) != 0 {
		return false
	}

	switch m.Action.(type) {

	case *Route_RouteAction:
		if _, ok := target.Action.(*Route_RouteAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRouteAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRouteAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRouteAction(), target.GetRouteAction()) {
				return false
			}
		}

	case *Route_RedirectAction:
		if _, ok := target.Action.(*Route_RedirectAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetRedirectAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetRedirectAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetRedirectAction(), target.GetRedirectAction()) {
				return false
			}
		}

	case *Route_DirectResponseAction:
		if _, ok := target.Action.(*Route_DirectResponseAction); !ok {
			return false
		}

		if h, ok := interface{}(m.GetDirectResponseAction()).(equality.Equalizer); ok {
			if !h.Equal(target.GetDirectResponseAction()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetDirectResponseAction(), target.GetDirectResponseAction()) {
				return false
			}
		}

	case *Route_GraphqlSchemaRef:
		if _, ok := target.Action.(*Route_GraphqlSchemaRef); !ok {
			return false
		}

		if h, ok := interface{}(m.GetGraphqlSchemaRef()).(equality.Equalizer); ok {
			if !h.Equal(target.GetGraphqlSchemaRef()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetGraphqlSchemaRef(), target.GetGraphqlSchemaRef()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Action != target.Action {
			return false
		}
	}

	return true
}

// Equal function
func (m *RouteAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RouteAction)
	if !ok {
		that2, ok := that.(RouteAction)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Destination.(type) {

	case *RouteAction_Single:
		if _, ok := target.Destination.(*RouteAction_Single); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSingle()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSingle()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSingle(), target.GetSingle()) {
				return false
			}
		}

	case *RouteAction_Multi:
		if _, ok := target.Destination.(*RouteAction_Multi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMulti()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMulti()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMulti(), target.GetMulti()) {
				return false
			}
		}

	case *RouteAction_UpstreamGroup:
		if _, ok := target.Destination.(*RouteAction_UpstreamGroup); !ok {
			return false
		}

		if h, ok := interface{}(m.GetUpstreamGroup()).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamGroup()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetUpstreamGroup(), target.GetUpstreamGroup()) {
				return false
			}
		}

	case *RouteAction_ClusterHeader:
		if _, ok := target.Destination.(*RouteAction_ClusterHeader); !ok {
			return false
		}

		if strings.Compare(m.GetClusterHeader(), target.GetClusterHeader()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.Destination != target.Destination {
			return false
		}
	}

	return true
}

// Equal function
func (m *Destination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*Destination)
	if !ok {
		that2, ok := that.(Destination)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetDestinationSpec()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestinationSpec()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestinationSpec(), target.GetDestinationSpec()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetSubset()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSubset()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSubset(), target.GetSubset()) {
			return false
		}
	}

	switch m.DestinationType.(type) {

	case *Destination_Upstream:
		if _, ok := target.DestinationType.(*Destination_Upstream); !ok {
			return false
		}

		if h, ok := interface{}(m.GetUpstream()).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstream()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetUpstream(), target.GetUpstream()) {
				return false
			}
		}

	case *Destination_Kube:
		if _, ok := target.DestinationType.(*Destination_Kube); !ok {
			return false
		}

		if h, ok := interface{}(m.GetKube()).(equality.Equalizer); ok {
			if !h.Equal(target.GetKube()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetKube(), target.GetKube()) {
				return false
			}
		}

	case *Destination_Consul:
		if _, ok := target.DestinationType.(*Destination_Consul); !ok {
			return false
		}

		if h, ok := interface{}(m.GetConsul()).(equality.Equalizer); ok {
			if !h.Equal(target.GetConsul()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetConsul(), target.GetConsul()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.DestinationType != target.DestinationType {
			return false
		}
	}

	return true
}

// Equal function
func (m *KubernetesServiceDestination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*KubernetesServiceDestination)
	if !ok {
		that2, ok := that.(KubernetesServiceDestination)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetRef(), target.GetRef()) {
			return false
		}
	}

	if m.GetPort() != target.GetPort() {
		return false
	}

	return true
}

// Equal function
func (m *ConsulServiceDestination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConsulServiceDestination)
	if !ok {
		that2, ok := that.(ConsulServiceDestination)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetServiceName(), target.GetServiceName()) != 0 {
		return false
	}

	if len(m.GetTags()) != len(target.GetTags()) {
		return false
	}
	for idx, v := range m.GetTags() {

		if strings.Compare(v, target.GetTags()[idx]) != 0 {
			return false
		}

	}

	if len(m.GetDataCenters()) != len(target.GetDataCenters()) {
		return false
	}
	for idx, v := range m.GetDataCenters() {

		if strings.Compare(v, target.GetDataCenters()[idx]) != 0 {
			return false
		}

	}

	return true
}

// Equal function
func (m *UpstreamGroup) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamGroup)
	if !ok {
		that2, ok := that.(UpstreamGroup)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetDestinations()) != len(target.GetDestinations()) {
		return false
	}
	for idx, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDestinations()[idx]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(equality.Equalizer); ok {
		if !h.Equal(target.GetNamespacedStatuses()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetNamespacedStatuses(), target.GetNamespacedStatuses()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetMetadata()).(equality.Equalizer); ok {
		if !h.Equal(target.GetMetadata()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetMetadata(), target.GetMetadata()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *MultiDestination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*MultiDestination)
	if !ok {
		that2, ok := that.(MultiDestination)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetDestinations()) != len(target.GetDestinations()) {
		return false
	}
	for idx, v := range m.GetDestinations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetDestinations()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetDestinations()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *WeightedDestination) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WeightedDestination)
	if !ok {
		that2, ok := that.(WeightedDestination)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetDestination()).(equality.Equalizer); ok {
		if !h.Equal(target.GetDestination()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetDestination(), target.GetDestination()) {
			return false
		}
	}

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	if h, ok := interface{}(m.GetOptions()).(equality.Equalizer); ok {
		if !h.Equal(target.GetOptions()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetOptions(), target.GetOptions()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *RedirectAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*RedirectAction)
	if !ok {
		that2, ok := that.(RedirectAction)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetHostRedirect(), target.GetHostRedirect()) != 0 {
		return false
	}

	if m.GetResponseCode() != target.GetResponseCode() {
		return false
	}

	if m.GetHttpsRedirect() != target.GetHttpsRedirect() {
		return false
	}

	if m.GetStripQuery() != target.GetStripQuery() {
		return false
	}

	switch m.PathRewriteSpecifier.(type) {

	case *RedirectAction_PathRedirect:
		if _, ok := target.PathRewriteSpecifier.(*RedirectAction_PathRedirect); !ok {
			return false
		}

		if strings.Compare(m.GetPathRedirect(), target.GetPathRedirect()) != 0 {
			return false
		}

	case *RedirectAction_PrefixRewrite:
		if _, ok := target.PathRewriteSpecifier.(*RedirectAction_PrefixRewrite); !ok {
			return false
		}

		if strings.Compare(m.GetPrefixRewrite(), target.GetPrefixRewrite()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.PathRewriteSpecifier != target.PathRewriteSpecifier {
			return false
		}
	}

	return true
}

// Equal function
func (m *DirectResponseAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DirectResponseAction)
	if !ok {
		that2, ok := that.(DirectResponseAction)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetStatus() != target.GetStatus() {
		return false
	}

	if strings.Compare(m.GetBody(), target.GetBody()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *TcpHost_TcpAction) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*TcpHost_TcpAction)
	if !ok {
		that2, ok := that.(TcpHost_TcpAction)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.Destination.(type) {

	case *TcpHost_TcpAction_Single:
		if _, ok := target.Destination.(*TcpHost_TcpAction_Single); !ok {
			return false
		}

		if h, ok := interface{}(m.GetSingle()).(equality.Equalizer); ok {
			if !h.Equal(target.GetSingle()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetSingle(), target.GetSingle()) {
				return false
			}
		}

	case *TcpHost_TcpAction_Multi:
		if _, ok := target.Destination.(*TcpHost_TcpAction_Multi); !ok {
			return false
		}

		if h, ok := interface{}(m.GetMulti()).(equality.Equalizer); ok {
			if !h.Equal(target.GetMulti()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetMulti(), target.GetMulti()) {
				return false
			}
		}

	case *TcpHost_TcpAction_UpstreamGroup:
		if _, ok := target.Destination.(*TcpHost_TcpAction_UpstreamGroup); !ok {
			return false
		}

		if h, ok := interface{}(m.GetUpstreamGroup()).(equality.Equalizer); ok {
			if !h.Equal(target.GetUpstreamGroup()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetUpstreamGroup(), target.GetUpstreamGroup()) {
				return false
			}
		}

	case *TcpHost_TcpAction_ForwardSniClusterName:
		if _, ok := target.Destination.(*TcpHost_TcpAction_ForwardSniClusterName); !ok {
			return false
		}

		if h, ok := interface{}(m.GetForwardSniClusterName()).(equality.Equalizer); ok {
			if !h.Equal(target.GetForwardSniClusterName()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetForwardSniClusterName(), target.GetForwardSniClusterName()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.Destination != target.Destination {
			return false
		}
	}

	return true
}
