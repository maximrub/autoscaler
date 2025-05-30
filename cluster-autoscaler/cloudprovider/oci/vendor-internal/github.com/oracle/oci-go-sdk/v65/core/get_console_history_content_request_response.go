// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"fmt"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/oci/vendor-internal/github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetConsoleHistoryContentRequest wrapper for the GetConsoleHistoryContent operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/GetConsoleHistoryContent.go.html to see an example of how to use GetConsoleHistoryContentRequest.
type GetConsoleHistoryContentRequest struct {

	// The OCID of the console history.
	InstanceConsoleHistoryId *string `mandatory:"true" contributesTo:"path" name:"instanceConsoleHistoryId"`

	// Offset of the snapshot data to retrieve.
	Offset *int `mandatory:"false" contributesTo:"query" name:"offset"`

	// Length of the snapshot data to retrieve.
	Length *int `mandatory:"false" contributesTo:"query" name:"length"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConsoleHistoryContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetConsoleHistoryContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetConsoleHistoryContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetConsoleHistoryContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetConsoleHistoryContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetConsoleHistoryContentResponse wrapper for the GetConsoleHistoryContent operation
type GetConsoleHistoryContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The string instance
	Value *string `presentIn:"body" encoding:"plain-text"`

	// The number of bytes remaining in the snapshot.
	OpcBytesRemaining *int `presentIn:"header" name:"opc-bytes-remaining"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConsoleHistoryContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetConsoleHistoryContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
