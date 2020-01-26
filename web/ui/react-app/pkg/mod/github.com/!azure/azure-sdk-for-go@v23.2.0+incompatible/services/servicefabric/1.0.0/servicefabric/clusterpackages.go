package servicefabric

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
	"net/http"
)

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// ClusterPackagesClient is the client for the ClusterPackages methods of the Servicefabric service.
type ClusterPackagesClient struct {
	BaseClient
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewClusterPackagesClient creates an instance of the ClusterPackagesClient client.
func NewClusterPackagesClient(timeout *int32) ClusterPackagesClient {
	return NewClusterPackagesClientWithBaseURI(DefaultBaseURI, timeout)
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// NewClusterPackagesClientWithBaseURI creates an instance of the ClusterPackagesClient client.
func NewClusterPackagesClientWithBaseURI(baseURI string, timeout *int32) ClusterPackagesClient {
	return ClusterPackagesClient{NewWithBaseURI(baseURI, timeout)}
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Register register cluster packages
// Parameters:
// registerClusterPackage - the package of the register cluster
func (client ClusterPackagesClient) Register(ctx context.Context, registerClusterPackage RegisterClusterPackage) (result String, err error) {
	req, err := client.RegisterPreparer(ctx, registerClusterPackage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Register", nil, "Failure preparing request")
		return
	}

	resp, err := client.RegisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Register", resp, "Failure sending request")
		return
	}

	result, err = client.RegisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Register", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RegisterPreparer prepares the Register request.
func (client ClusterPackagesClient) RegisterPreparer(ctx context.Context, registerClusterPackage RegisterClusterPackage) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/Provision"),
		autorest.WithJSON(registerClusterPackage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RegisterSender sends the Register request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPackagesClient) RegisterSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// RegisterResponder handles the response to the Register request. The method always
// closes the http.Response Body.
func (client ClusterPackagesClient) RegisterResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// Unregister unregister cluster packages
// Parameters:
// unregisterClusterPackage - the package of the unregister cluster
func (client ClusterPackagesClient) Unregister(ctx context.Context, unregisterClusterPackage UnregisterClusterPackage) (result String, err error) {
	req, err := client.UnregisterPreparer(ctx, unregisterClusterPackage)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Unregister", nil, "Failure preparing request")
		return
	}

	resp, err := client.UnregisterSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Unregister", resp, "Failure sending request")
		return
	}

	result, err = client.UnregisterResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicefabric.ClusterPackagesClient", "Unregister", resp, "Failure responding to request")
	}

	return
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// UnregisterPreparer prepares the Unregister request.
func (client ClusterPackagesClient) UnregisterPreparer(ctx context.Context, unregisterClusterPackage UnregisterClusterPackage) (*http.Request, error) {
	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if client.Timeout != nil {
		queryParameters["timeout"] = autorest.Encode("query", *client.Timeout)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/$/Unprovision"),
		autorest.WithJSON(unregisterClusterPackage),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// UnregisterSender sends the Unregister request. The method will close the
// http.Response Body if it receives an error.
func (client ClusterPackagesClient) UnregisterSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// Deprecated: Please use package github.com/Azure/azure-sdk-for-go/services/servicefabric/6.2/servicefabric instead
// UnregisterResponder handles the response to the Unregister request. The method always
// closes the http.Response Body.
func (client ClusterPackagesClient) UnregisterResponder(resp *http.Response) (result String, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
