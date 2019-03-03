//+build !cloudprovider_thirdparty

package cloudinfo

import (
	info "github.com/google/cadvisor/info/v1"
)

func detectCloudProvider() info.CloudProvider {
	switch {
	case onBaremetal():
		return info.Baremetal
	}
	return info.UnknownProvider
}

func detectInstanceType(cloudProvider info.CloudProvider) info.InstanceType {
	switch cloudProvider {
	case info.Baremetal:
		return info.NoInstance
	}
	return info.UnknownInstance
}

func detectInstanceID(cloudProvider info.CloudProvider) info.InstanceID {
	switch cloudProvider {
	case info.Baremetal:
		return info.UnNamedInstance
	}
	return info.UnNamedInstance
}
