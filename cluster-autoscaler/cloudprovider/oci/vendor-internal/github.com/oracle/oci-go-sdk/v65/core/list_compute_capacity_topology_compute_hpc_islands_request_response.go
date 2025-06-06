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

// ListComputeCapacityTopologyComputeHpcIslandsRequest wrapper for the ListComputeCapacityTopologyComputeHpcIslands operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/core/ListComputeCapacityTopologyComputeHpcIslands.go.html to see an example of how to use ListComputeCapacityTopologyComputeHpcIslandsRequest.
type ListComputeCapacityTopologyComputeHpcIslandsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute capacity topology.
	ComputeCapacityTopologyId *string `mandatory:"true" contributesTo:"path" name:"computeCapacityTopologyId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListComputeCapacityTopologyComputeHpcIslandsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListComputeCapacityTopologyComputeHpcIslandsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListComputeCapacityTopologyComputeHpcIslandsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListComputeCapacityTopologyComputeHpcIslandsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListComputeCapacityTopologyComputeHpcIslandsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListComputeCapacityTopologyComputeHpcIslandsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListComputeCapacityTopologyComputeHpcIslandsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListComputeCapacityTopologyComputeHpcIslandsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListComputeCapacityTopologyComputeHpcIslandsResponse wrapper for the ListComputeCapacityTopologyComputeHpcIslands operation
type ListComputeCapacityTopologyComputeHpcIslandsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ComputeHpcIslandCollection instances
	ComputeHpcIslandCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListComputeCapacityTopologyComputeHpcIslandsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListComputeCapacityTopologyComputeHpcIslandsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListComputeCapacityTopologyComputeHpcIslandsSortByEnum Enum with underlying type: string
type ListComputeCapacityTopologyComputeHpcIslandsSortByEnum string

// Set of constants representing the allowable values for ListComputeCapacityTopologyComputeHpcIslandsSortByEnum
const (
	ListComputeCapacityTopologyComputeHpcIslandsSortByTimecreated ListComputeCapacityTopologyComputeHpcIslandsSortByEnum = "TIMECREATED"
	ListComputeCapacityTopologyComputeHpcIslandsSortByDisplayname ListComputeCapacityTopologyComputeHpcIslandsSortByEnum = "DISPLAYNAME"
)

var mappingListComputeCapacityTopologyComputeHpcIslandsSortByEnum = map[string]ListComputeCapacityTopologyComputeHpcIslandsSortByEnum{
	"TIMECREATED": ListComputeCapacityTopologyComputeHpcIslandsSortByTimecreated,
	"DISPLAYNAME": ListComputeCapacityTopologyComputeHpcIslandsSortByDisplayname,
}

var mappingListComputeCapacityTopologyComputeHpcIslandsSortByEnumLowerCase = map[string]ListComputeCapacityTopologyComputeHpcIslandsSortByEnum{
	"timecreated": ListComputeCapacityTopologyComputeHpcIslandsSortByTimecreated,
	"displayname": ListComputeCapacityTopologyComputeHpcIslandsSortByDisplayname,
}

// GetListComputeCapacityTopologyComputeHpcIslandsSortByEnumValues Enumerates the set of values for ListComputeCapacityTopologyComputeHpcIslandsSortByEnum
func GetListComputeCapacityTopologyComputeHpcIslandsSortByEnumValues() []ListComputeCapacityTopologyComputeHpcIslandsSortByEnum {
	values := make([]ListComputeCapacityTopologyComputeHpcIslandsSortByEnum, 0)
	for _, v := range mappingListComputeCapacityTopologyComputeHpcIslandsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeCapacityTopologyComputeHpcIslandsSortByEnumStringValues Enumerates the set of values in String for ListComputeCapacityTopologyComputeHpcIslandsSortByEnum
func GetListComputeCapacityTopologyComputeHpcIslandsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListComputeCapacityTopologyComputeHpcIslandsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeCapacityTopologyComputeHpcIslandsSortByEnum(val string) (ListComputeCapacityTopologyComputeHpcIslandsSortByEnum, bool) {
	enum, ok := mappingListComputeCapacityTopologyComputeHpcIslandsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum Enum with underlying type: string
type ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum string

// Set of constants representing the allowable values for ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum
const (
	ListComputeCapacityTopologyComputeHpcIslandsSortOrderAsc  ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum = "ASC"
	ListComputeCapacityTopologyComputeHpcIslandsSortOrderDesc ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum = "DESC"
)

var mappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum = map[string]ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum{
	"ASC":  ListComputeCapacityTopologyComputeHpcIslandsSortOrderAsc,
	"DESC": ListComputeCapacityTopologyComputeHpcIslandsSortOrderDesc,
}

var mappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumLowerCase = map[string]ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum{
	"asc":  ListComputeCapacityTopologyComputeHpcIslandsSortOrderAsc,
	"desc": ListComputeCapacityTopologyComputeHpcIslandsSortOrderDesc,
}

// GetListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumValues Enumerates the set of values for ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum
func GetListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumValues() []ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum {
	values := make([]ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum, 0)
	for _, v := range mappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumStringValues Enumerates the set of values in String for ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum
func GetListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum(val string) (ListComputeCapacityTopologyComputeHpcIslandsSortOrderEnum, bool) {
	enum, ok := mappingListComputeCapacityTopologyComputeHpcIslandsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
