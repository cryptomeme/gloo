// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/aws/filter.proto

package aws

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
func (m *AWSLambdaPerRoute) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSLambdaPerRoute)
	if !ok {
		that2, ok := that.(AWSLambdaPerRoute)
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

	if strings.Compare(m.GetQualifier(), target.GetQualifier()) != 0 {
		return false
	}

	if m.GetAsync() != target.GetAsync() {
		return false
	}

	if h, ok := interface{}(m.GetEmptyBodyOverride()).(equality.Equalizer); ok {
		if !h.Equal(target.GetEmptyBodyOverride()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetEmptyBodyOverride(), target.GetEmptyBodyOverride()) {
			return false
		}
	}

	if m.GetUnwrapAsAlb() != target.GetUnwrapAsAlb() {
		return false
	}

	return true
}

// Equal function
func (m *AWSLambdaProtocolExtension) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSLambdaProtocolExtension)
	if !ok {
		that2, ok := that.(AWSLambdaProtocolExtension)
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

	if strings.Compare(m.GetHost(), target.GetHost()) != 0 {
		return false
	}

	if strings.Compare(m.GetRegion(), target.GetRegion()) != 0 {
		return false
	}

	if strings.Compare(m.GetAccessKey(), target.GetAccessKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetSecretKey(), target.GetSecretKey()) != 0 {
		return false
	}

	if strings.Compare(m.GetSessionToken(), target.GetSessionToken()) != 0 {
		return false
	}

	if strings.Compare(m.GetRoleArn(), target.GetRoleArn()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *AWSLambdaConfig) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSLambdaConfig)
	if !ok {
		that2, ok := that.(AWSLambdaConfig)
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

	if m.GetPropagateOriginalRouting() != target.GetPropagateOriginalRouting() {
		return false
	}

	if h, ok := interface{}(m.GetCredentialRefreshDelay()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCredentialRefreshDelay()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCredentialRefreshDelay(), target.GetCredentialRefreshDelay()) {
			return false
		}
	}

	switch m.CredentialsFetcher.(type) {

	case *AWSLambdaConfig_UseDefaultCredentials:
		if _, ok := target.CredentialsFetcher.(*AWSLambdaConfig_UseDefaultCredentials); !ok {
			return false
		}

		if h, ok := interface{}(m.GetUseDefaultCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetUseDefaultCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetUseDefaultCredentials(), target.GetUseDefaultCredentials()) {
				return false
			}
		}

	case *AWSLambdaConfig_ServiceAccountCredentials_:
		if _, ok := target.CredentialsFetcher.(*AWSLambdaConfig_ServiceAccountCredentials_); !ok {
			return false
		}

		if h, ok := interface{}(m.GetServiceAccountCredentials()).(equality.Equalizer); ok {
			if !h.Equal(target.GetServiceAccountCredentials()) {
				return false
			}
		} else {
			if !proto.Equal(m.GetServiceAccountCredentials(), target.GetServiceAccountCredentials()) {
				return false
			}
		}

	default:
		// m is nil but target is not nil
		if m.CredentialsFetcher != target.CredentialsFetcher {
			return false
		}
	}

	return true
}

// Equal function
func (m *AWSLambdaConfig_ServiceAccountCredentials) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*AWSLambdaConfig_ServiceAccountCredentials)
	if !ok {
		that2, ok := that.(AWSLambdaConfig_ServiceAccountCredentials)
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

	if strings.Compare(m.GetCluster(), target.GetCluster()) != 0 {
		return false
	}

	if strings.Compare(m.GetUri(), target.GetUri()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetTimeout()).(equality.Equalizer); ok {
		if !h.Equal(target.GetTimeout()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetTimeout(), target.GetTimeout()) {
			return false
		}
	}

	return true
}
