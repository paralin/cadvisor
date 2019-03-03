// Copyright 2015 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Get information about the cloud provider (if any) cAdvisor is running on.

package cloudinfo

import (
	info "github.com/google/cadvisor/info/v1"
)

type CloudInfo interface {
	GetCloudProvider() info.CloudProvider
	GetInstanceType() info.InstanceType
	GetInstanceID() info.InstanceID
}

type realCloudInfo struct {
	cloudProvider info.CloudProvider
	instanceType  info.InstanceType
	instanceID    info.InstanceID
}

func NewRealCloudInfo() CloudInfo {
	cloudProvider := detectCloudProvider()
	instanceType := detectInstanceType(cloudProvider)
	instanceID := detectInstanceID(cloudProvider)
	return &realCloudInfo{
		cloudProvider: cloudProvider,
		instanceType:  instanceType,
		instanceID:    instanceID,
	}
}

func (self *realCloudInfo) GetCloudProvider() info.CloudProvider {
	return self.cloudProvider
}

func (self *realCloudInfo) GetInstanceType() info.InstanceType {
	return self.instanceType
}

func (self *realCloudInfo) GetInstanceID() info.InstanceID {
	return self.instanceID
}

// TODO: Implement method.
func onBaremetal() bool {
	return false
}
