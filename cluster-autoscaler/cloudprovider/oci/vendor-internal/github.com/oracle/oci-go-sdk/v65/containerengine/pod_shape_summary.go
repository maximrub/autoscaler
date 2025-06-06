// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Kubernetes Engine API
//
// API for the Kubernetes Engine service (also known as the Container Engine for Kubernetes service). Use this API to build, deploy,
// and manage cloud-native applications. For more information, see
// Overview of Kubernetes Engine (https://docs.oracle.com/iaas/Content/ContEng/Concepts/contengoverview.htm).
//

package containerengine

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PodShapeSummary Pod shape.
type PodShapeSummary struct {

	// The name of the identifying shape.
	Name *string `mandatory:"true" json:"name"`

	// A short description of the VM's processor (CPU).
	ProcessorDescription *string `mandatory:"false" json:"processorDescription"`

	// Options for OCPU shape.
	OcpuOptions []ShapeOcpuOptions `mandatory:"false" json:"ocpuOptions"`

	// ShapeMemoryOptions.
	MemoryOptions []ShapeMemoryOptions `mandatory:"false" json:"memoryOptions"`

	// ShapeNetworkBandwidthOptions.
	NetworkBandwidthOptions []ShapeNetworkBandwidthOptions `mandatory:"false" json:"networkBandwidthOptions"`
}

func (m PodShapeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PodShapeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
