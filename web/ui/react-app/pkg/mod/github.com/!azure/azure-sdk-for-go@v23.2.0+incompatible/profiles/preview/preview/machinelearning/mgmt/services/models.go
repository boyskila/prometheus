// +build go1.9

// Copyright 2018 Microsoft Corporation
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

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package services

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/machinelearning/mgmt/2018-03-01-preview/services"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type BaseClient = original.BaseClient
type MachineLearningComputeClient = original.MachineLearningComputeClient
type ComputeType = original.ComputeType

const (
	ComputeTypeAKS            ComputeType = original.ComputeTypeAKS
	ComputeTypeBatchAI        ComputeType = original.ComputeTypeBatchAI
	ComputeTypeDataFactory    ComputeType = original.ComputeTypeDataFactory
	ComputeTypeHDInsight      ComputeType = original.ComputeTypeHDInsight
	ComputeTypeVirtualMachine ComputeType = original.ComputeTypeVirtualMachine
)

type ComputeTypeBasicCompute = original.ComputeTypeBasicCompute

const (
	ComputeTypeAKS1            ComputeTypeBasicCompute = original.ComputeTypeAKS1
	ComputeTypeBatchAI1        ComputeTypeBasicCompute = original.ComputeTypeBatchAI1
	ComputeTypeCompute         ComputeTypeBasicCompute = original.ComputeTypeCompute
	ComputeTypeDataFactory1    ComputeTypeBasicCompute = original.ComputeTypeDataFactory1
	ComputeTypeHDInsight1      ComputeTypeBasicCompute = original.ComputeTypeHDInsight1
	ComputeTypeVirtualMachine1 ComputeTypeBasicCompute = original.ComputeTypeVirtualMachine1
)

type ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecrets

const (
	ComputeTypeBasicComputeSecretsComputeTypeAKS            ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeAKS
	ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeComputeSecrets
	ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine ComputeTypeBasicComputeSecrets = original.ComputeTypeBasicComputeSecretsComputeTypeVirtualMachine
)

type ProvisioningState = original.ProvisioningState

const (
	Canceled  ProvisioningState = original.Canceled
	Creating  ProvisioningState = original.Creating
	Deleting  ProvisioningState = original.Deleting
	Failed    ProvisioningState = original.Failed
	Succeeded ProvisioningState = original.Succeeded
	Unknown   ProvisioningState = original.Unknown
	Updating  ProvisioningState = original.Updating
)

type ResourceIdentityType = original.ResourceIdentityType

const (
	SystemAssigned ResourceIdentityType = original.SystemAssigned
)

type Status = original.Status

const (
	Disabled Status = original.Disabled
	Enabled  Status = original.Enabled
)

type AKS = original.AKS
type AksComputeSecrets = original.AksComputeSecrets
type AKSProperties = original.AKSProperties
type BatchAI = original.BatchAI
type BatchAIProperties = original.BatchAIProperties
type BasicCompute = original.BasicCompute
type Compute = original.Compute
type ComputeResource = original.ComputeResource
type BasicComputeSecrets = original.BasicComputeSecrets
type ComputeSecrets = original.ComputeSecrets
type ComputeSecretsModel = original.ComputeSecretsModel
type DataFactory = original.DataFactory
type ErrorDetail = original.ErrorDetail
type ErrorResponse = original.ErrorResponse
type HDInsight = original.HDInsight
type HDInsightProperties = original.HDInsightProperties
type Identity = original.Identity
type ListWorkspaceKeysResult = original.ListWorkspaceKeysResult
type MachineLearningComputeCreateOrUpdateFuture = original.MachineLearningComputeCreateOrUpdateFuture
type MachineLearningComputeDeleteFuture = original.MachineLearningComputeDeleteFuture
type MachineLearningComputeSystemUpdateFuture = original.MachineLearningComputeSystemUpdateFuture
type MachineLearningServiceError = original.MachineLearningServiceError
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type PaginatedComputeResourcesList = original.PaginatedComputeResourcesList
type PaginatedComputeResourcesListIterator = original.PaginatedComputeResourcesListIterator
type PaginatedComputeResourcesListPage = original.PaginatedComputeResourcesListPage
type Password = original.Password
type PrincipalCredentials = original.PrincipalCredentials
type RegistryListCredentialsResult = original.RegistryListCredentialsResult
type Resource = original.Resource
type ScaleSettings = original.ScaleSettings
type SslConfiguration = original.SslConfiguration
type SystemService = original.SystemService
type VirtualMachine = original.VirtualMachine
type VirtualMachineProperties = original.VirtualMachineProperties
type VirtualMachineSecrets = original.VirtualMachineSecrets
type VirtualMachineSSHCredentials = original.VirtualMachineSSHCredentials
type Workspace = original.Workspace
type WorkspaceListResult = original.WorkspaceListResult
type WorkspaceListResultIterator = original.WorkspaceListResultIterator
type WorkspaceListResultPage = original.WorkspaceListResultPage
type WorkspaceProperties = original.WorkspaceProperties
type WorkspacePropertiesUpdateParameters = original.WorkspacePropertiesUpdateParameters
type WorkspaceUpdateParameters = original.WorkspaceUpdateParameters
type OperationsClient = original.OperationsClient
type WorkspacesClient = original.WorkspacesClient

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func NewMachineLearningComputeClient(subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClient(subscriptionID)
}
func NewMachineLearningComputeClientWithBaseURI(baseURI string, subscriptionID string) MachineLearningComputeClient {
	return original.NewMachineLearningComputeClientWithBaseURI(baseURI, subscriptionID)
}
func PossibleComputeTypeValues() []ComputeType {
	return original.PossibleComputeTypeValues()
}
func PossibleComputeTypeBasicComputeValues() []ComputeTypeBasicCompute {
	return original.PossibleComputeTypeBasicComputeValues()
}
func PossibleComputeTypeBasicComputeSecretsValues() []ComputeTypeBasicComputeSecrets {
	return original.PossibleComputeTypeBasicComputeSecretsValues()
}
func PossibleProvisioningStateValues() []ProvisioningState {
	return original.PossibleProvisioningStateValues()
}
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return original.PossibleResourceIdentityTypeValues()
}
func PossibleStatusValues() []Status {
	return original.PossibleStatusValues()
}
func NewPaginatedComputeResourcesListIterator(page PaginatedComputeResourcesListPage) PaginatedComputeResourcesListIterator {
	return original.NewPaginatedComputeResourcesListIterator(page)
}
func NewPaginatedComputeResourcesListPage(getNextPage func(context.Context, PaginatedComputeResourcesList) (PaginatedComputeResourcesList, error)) PaginatedComputeResourcesListPage {
	return original.NewPaginatedComputeResourcesListPage(getNextPage)
}
func NewWorkspaceListResultIterator(page WorkspaceListResultPage) WorkspaceListResultIterator {
	return original.NewWorkspaceListResultIterator(page)
}
func NewWorkspaceListResultPage(getNextPage func(context.Context, WorkspaceListResult) (WorkspaceListResult, error)) WorkspaceListResultPage {
	return original.NewWorkspaceListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
func NewWorkspacesClient(subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClient(subscriptionID)
}
func NewWorkspacesClientWithBaseURI(baseURI string, subscriptionID string) WorkspacesClient {
	return original.NewWorkspacesClientWithBaseURI(baseURI, subscriptionID)
}
