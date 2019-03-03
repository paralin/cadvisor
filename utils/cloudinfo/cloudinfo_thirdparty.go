//+build cloudprovider_thirdparty

package cloudinfo

import (
	info "github.com/google/cadvisor/info/v1"
)

func detectCloudProvider() info.CloudProvider {
	switch {
	case onGCE():
		return info.GCE
	case onAWS():
		return info.AWS
	case onAzure():
		return info.Azure
	case onBaremetal():
		return info.Baremetal
	}
	return info.UnknownProvider
}

func detectInstanceType(cloudProvider info.CloudProvider) info.InstanceType {
	switch cloudProvider {
	case info.GCE:
		return getGceInstanceType()
	case info.AWS:
		return getAwsInstanceType()
	case info.Azure:
		return getAzureInstanceType()
	case info.Baremetal:
		return info.NoInstance
	}
	return info.UnknownInstance
}

func detectInstanceID(cloudProvider info.CloudProvider) info.InstanceID {
	switch cloudProvider {
	case info.GCE:
		return getGceInstanceID()
	case info.AWS:
		return getAwsInstanceID()
	case info.Azure:
		return getAzureInstanceID()
	case info.Baremetal:
		return info.UnNamedInstance
	}
	return info.UnNamedInstance
}
