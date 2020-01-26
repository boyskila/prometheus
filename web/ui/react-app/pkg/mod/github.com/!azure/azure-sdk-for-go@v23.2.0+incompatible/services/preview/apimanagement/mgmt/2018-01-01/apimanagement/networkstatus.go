package apimanagement

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// NetworkStatusClient is the apiManagement Client
type NetworkStatusClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// NewNetworkStatusClient creates an instance of the NetworkStatusClient client.
func NewNetworkStatusClient(subscriptionID string) NetworkStatusClient {
	return NewNetworkStatusClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// NewNetworkStatusClientWithBaseURI creates an instance of the NetworkStatusClient client.
func NewNetworkStatusClientWithBaseURI(baseURI string, subscriptionID string) NetworkStatusClient {
	return NetworkStatusClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByLocation gets the Connectivity Status to the external resources on which the Api Management service depends
// from inside the Cloud Service. This also returns the DNS Servers as visible to the CloudService.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
// locationName - location in which the API Management service is deployed. This is one of the Azure Regions
// like West US, East US, South Central US.
func (client NetworkStatusClient) ListByLocation(ctx context.Context, resourceGroupName string, serviceName string, locationName string) (result NetworkStatusContract, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkStatusClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}},
		{TargetValue: locationName,
			Constraints: []validation.Constraint{{Target: "locationName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NetworkStatusClient", "ListByLocation", err.Error())
	}

	req, err := client.ListByLocationPreparer(ctx, resourceGroupName, serviceName, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByLocation", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByLocationPreparer prepares the ListByLocation request.
func (client NetworkStatusClient) ListByLocationPreparer(ctx context.Context, resourceGroupName string, serviceName string, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":      autorest.Encode("path", locationName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/locations/{locationName}/networkstatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkStatusClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client NetworkStatusClient) ListByLocationResponder(resp *http.Response) (result NetworkStatusContract, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByService gets the Connectivity Status to the external resources on which the Api Management service depends
// from inside the Cloud Service. This also returns the DNS Servers as visible to the CloudService.
// Parameters:
// resourceGroupName - the name of the resource group.
// serviceName - the name of the API Management service.
func (client NetworkStatusClient) ListByService(ctx context.Context, resourceGroupName string, serviceName string) (result ListNetworkStatusContractByLocation, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/NetworkStatusClient.ListByService")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: serviceName,
			Constraints: []validation.Constraint{{Target: "serviceName", Name: validation.MaxLength, Rule: 50, Chain: nil},
				{Target: "serviceName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "serviceName", Name: validation.Pattern, Rule: `^[a-zA-Z](?:[a-zA-Z0-9-]*[a-zA-Z0-9])?$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("apimanagement.NetworkStatusClient", "ListByService", err.Error())
	}

	req, err := client.ListByServicePreparer(ctx, resourceGroupName, serviceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByService", resp, "Failure sending request")
		return
	}

	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "apimanagement.NetworkStatusClient", "ListByService", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByServicePreparer prepares the ListByService request.
func (client NetworkStatusClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, serviceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"serviceName":       autorest.Encode("path", serviceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-01-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/networkstatus", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client NetworkStatusClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/apimanagement/mgmt/2018-01-01/apimanagement instead.
// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client NetworkStatusClient) ListByServiceResponder(resp *http.Response) (result ListNetworkStatusContractByLocation, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
