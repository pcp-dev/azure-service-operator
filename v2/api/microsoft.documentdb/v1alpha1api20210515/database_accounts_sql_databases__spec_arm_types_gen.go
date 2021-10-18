// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabases_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion DatabaseAccountsSqlDatabasesSpecAPIVersion `json:"apiVersion"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: Cosmos DB database name.
	Name string `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB SQL database.
	Properties SqlDatabaseCreateUpdatePropertiesARM `json:"properties"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type DatabaseAccountsSqlDatabasesSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabases_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (databaseAccountsSqlDatabasesSpecARM DatabaseAccountsSqlDatabases_SpecARM) GetAPIVersion() string {
	return string(databaseAccountsSqlDatabasesSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (databaseAccountsSqlDatabasesSpecARM DatabaseAccountsSqlDatabases_SpecARM) GetName() string {
	return databaseAccountsSqlDatabasesSpecARM.Name
}

// GetType returns the Type of the resource
func (databaseAccountsSqlDatabasesSpecARM DatabaseAccountsSqlDatabases_SpecARM) GetType() string {
	return string(databaseAccountsSqlDatabasesSpecARM.Type)
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSqlDatabasesSpecAPIVersion string

const DatabaseAccountsSqlDatabasesSpecAPIVersion20210515 = DatabaseAccountsSqlDatabasesSpecAPIVersion("2021-05-15")

// +kubebuilder:validation:Enum={"Microsoft.DocumentDB/databaseAccounts/sqlDatabases"}
type DatabaseAccountsSqlDatabasesSpecType string

const DatabaseAccountsSqlDatabasesSpecTypeMicrosoftDocumentDBDatabaseAccountsSqlDatabases = DatabaseAccountsSqlDatabasesSpecType("Microsoft.DocumentDB/databaseAccounts/sqlDatabases")

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseCreateUpdateProperties
type SqlDatabaseCreateUpdatePropertiesARM struct {
	//Options: CreateUpdateOptions are a list of key-value pairs that describe the
	//resource. Supported keys are "If-Match", "If-None-Match", "Session-Token" and
	//"Throughput"
	Options *CreateUpdateOptionsARM `json:"options,omitempty"`

	//Resource: Cosmos DB SQL database resource object
	Resource SqlDatabaseResourceARM `json:"resource"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource
type SqlDatabaseResourceARM struct {
	//Id: Name of the Cosmos DB SQL database
	Id string `json:"id"`
}