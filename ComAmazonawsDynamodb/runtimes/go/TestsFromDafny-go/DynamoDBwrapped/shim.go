// Code generated by smithy-go-codegen DO NOT EDIT.

package DynamoDBwrapped

import (
	"context"

	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/dynamodb/ComAmazonawsDynamodbTypes"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/dynamodb/comamazonawsdynamodbsmithygenerated"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library/Wrappers"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"
)

type Shim struct {
	ComAmazonawsDynamodbTypes.IDynamoDBClient
	Client *dynamodb.Client
}

func (shim *Shim) BatchExecuteStatement(input ComAmazonawsDynamodbTypes.BatchExecuteStatementInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.BatchExecuteStatementInput_FromDafny(input)
	var native_response, native_error = shim.Client.BatchExecuteStatement(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.BatchExecuteStatementOutput_ToDafny(*native_response))
}

func (shim *Shim) BatchGetItem(input ComAmazonawsDynamodbTypes.BatchGetItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.BatchGetItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.BatchGetItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.BatchGetItemOutput_ToDafny(*native_response))
}

func (shim *Shim) BatchWriteItem(input ComAmazonawsDynamodbTypes.BatchWriteItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.BatchWriteItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.BatchWriteItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.BatchWriteItemOutput_ToDafny(*native_response))
}

func (shim *Shim) CreateBackup(input ComAmazonawsDynamodbTypes.CreateBackupInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.CreateBackupInput_FromDafny(input)
	var native_response, native_error = shim.Client.CreateBackup(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.CreateBackupOutput_ToDafny(*native_response))
}

func (shim *Shim) CreateGlobalTable(input ComAmazonawsDynamodbTypes.CreateGlobalTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.CreateGlobalTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.CreateGlobalTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.CreateGlobalTableOutput_ToDafny(*native_response))
}

func (shim *Shim) CreateTable(input ComAmazonawsDynamodbTypes.CreateTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.CreateTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.CreateTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.CreateTableOutput_ToDafny(*native_response))
}

func (shim *Shim) DeleteBackup(input ComAmazonawsDynamodbTypes.DeleteBackupInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DeleteBackupInput_FromDafny(input)
	var native_response, native_error = shim.Client.DeleteBackup(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DeleteBackupOutput_ToDafny(*native_response))
}

func (shim *Shim) DeleteItem(input ComAmazonawsDynamodbTypes.DeleteItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DeleteItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.DeleteItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DeleteItemOutput_ToDafny(*native_response))
}

func (shim *Shim) DeleteResourcePolicy(input ComAmazonawsDynamodbTypes.DeleteResourcePolicyInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DeleteResourcePolicyInput_FromDafny(input)
	var native_response, native_error = shim.Client.DeleteResourcePolicy(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DeleteResourcePolicyOutput_ToDafny(*native_response))
}

func (shim *Shim) DeleteTable(input ComAmazonawsDynamodbTypes.DeleteTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DeleteTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.DeleteTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DeleteTableOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeBackup(input ComAmazonawsDynamodbTypes.DescribeBackupInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeBackupInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeBackup(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeBackupOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeContinuousBackups(input ComAmazonawsDynamodbTypes.DescribeContinuousBackupsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeContinuousBackupsInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeContinuousBackups(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeContinuousBackupsOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeContributorInsights(input ComAmazonawsDynamodbTypes.DescribeContributorInsightsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeContributorInsightsInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeContributorInsights(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeContributorInsightsOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeEndpoints(input ComAmazonawsDynamodbTypes.DescribeEndpointsRequest) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeEndpointsInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeEndpoints(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeEndpointsOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeExport(input ComAmazonawsDynamodbTypes.DescribeExportInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeExportInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeExport(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeExportOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeGlobalTable(input ComAmazonawsDynamodbTypes.DescribeGlobalTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeGlobalTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeGlobalTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeGlobalTableOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeGlobalTableSettings(input ComAmazonawsDynamodbTypes.DescribeGlobalTableSettingsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeGlobalTableSettingsInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeGlobalTableSettings(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeGlobalTableSettingsOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeImport(input ComAmazonawsDynamodbTypes.DescribeImportInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeImportInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeImport(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeImportOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeKinesisStreamingDestination(input ComAmazonawsDynamodbTypes.DescribeKinesisStreamingDestinationInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeKinesisStreamingDestinationInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeKinesisStreamingDestination(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeKinesisStreamingDestinationOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeLimits(input ComAmazonawsDynamodbTypes.DescribeLimitsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeLimitsInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeLimits(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeLimitsOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeTable(input ComAmazonawsDynamodbTypes.DescribeTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeTableOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeTableReplicaAutoScaling(input ComAmazonawsDynamodbTypes.DescribeTableReplicaAutoScalingInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeTableReplicaAutoScalingInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeTableReplicaAutoScaling(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeTableReplicaAutoScalingOutput_ToDafny(*native_response))
}

func (shim *Shim) DescribeTimeToLive(input ComAmazonawsDynamodbTypes.DescribeTimeToLiveInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DescribeTimeToLiveInput_FromDafny(input)
	var native_response, native_error = shim.Client.DescribeTimeToLive(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DescribeTimeToLiveOutput_ToDafny(*native_response))
}

func (shim *Shim) DisableKinesisStreamingDestination(input ComAmazonawsDynamodbTypes.DisableKinesisStreamingDestinationInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.DisableKinesisStreamingDestinationInput_FromDafny(input)
	var native_response, native_error = shim.Client.DisableKinesisStreamingDestination(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.DisableKinesisStreamingDestinationOutput_ToDafny(*native_response))
}

func (shim *Shim) EnableKinesisStreamingDestination(input ComAmazonawsDynamodbTypes.EnableKinesisStreamingDestinationInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.EnableKinesisStreamingDestinationInput_FromDafny(input)
	var native_response, native_error = shim.Client.EnableKinesisStreamingDestination(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.EnableKinesisStreamingDestinationOutput_ToDafny(*native_response))
}

func (shim *Shim) ExecuteStatement(input ComAmazonawsDynamodbTypes.ExecuteStatementInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ExecuteStatementInput_FromDafny(input)
	var native_response, native_error = shim.Client.ExecuteStatement(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ExecuteStatementOutput_ToDafny(*native_response))
}

func (shim *Shim) ExecuteTransaction(input ComAmazonawsDynamodbTypes.ExecuteTransactionInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ExecuteTransactionInput_FromDafny(input)
	var native_response, native_error = shim.Client.ExecuteTransaction(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ExecuteTransactionOutput_ToDafny(*native_response))
}

func (shim *Shim) ExportTableToPointInTime(input ComAmazonawsDynamodbTypes.ExportTableToPointInTimeInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ExportTableToPointInTimeInput_FromDafny(input)
	var native_response, native_error = shim.Client.ExportTableToPointInTime(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ExportTableToPointInTimeOutput_ToDafny(*native_response))
}

func (shim *Shim) GetItem(input ComAmazonawsDynamodbTypes.GetItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.GetItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.GetItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.GetItemOutput_ToDafny(*native_response))
}

func (shim *Shim) GetResourcePolicy(input ComAmazonawsDynamodbTypes.GetResourcePolicyInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.GetResourcePolicyInput_FromDafny(input)
	var native_response, native_error = shim.Client.GetResourcePolicy(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.GetResourcePolicyOutput_ToDafny(*native_response))
}

func (shim *Shim) ImportTable(input ComAmazonawsDynamodbTypes.ImportTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ImportTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.ImportTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ImportTableOutput_ToDafny(*native_response))
}

func (shim *Shim) ListBackups(input ComAmazonawsDynamodbTypes.ListBackupsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListBackupsInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListBackups(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListBackupsOutput_ToDafny(*native_response))
}

func (shim *Shim) ListContributorInsights(input ComAmazonawsDynamodbTypes.ListContributorInsightsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListContributorInsightsInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListContributorInsights(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListContributorInsightsOutput_ToDafny(*native_response))
}

func (shim *Shim) ListExports(input ComAmazonawsDynamodbTypes.ListExportsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListExportsInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListExports(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListExportsOutput_ToDafny(*native_response))
}

func (shim *Shim) ListGlobalTables(input ComAmazonawsDynamodbTypes.ListGlobalTablesInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListGlobalTablesInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListGlobalTables(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListGlobalTablesOutput_ToDafny(*native_response))
}

func (shim *Shim) ListImports(input ComAmazonawsDynamodbTypes.ListImportsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListImportsInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListImports(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListImportsOutput_ToDafny(*native_response))
}

func (shim *Shim) ListTables(input ComAmazonawsDynamodbTypes.ListTablesInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListTablesInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListTables(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListTablesOutput_ToDafny(*native_response))
}

func (shim *Shim) ListTagsOfResource(input ComAmazonawsDynamodbTypes.ListTagsOfResourceInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ListTagsOfResourceInput_FromDafny(input)
	var native_response, native_error = shim.Client.ListTagsOfResource(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ListTagsOfResourceOutput_ToDafny(*native_response))
}

func (shim *Shim) PutItem(input ComAmazonawsDynamodbTypes.PutItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.PutItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.PutItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.PutItemOutput_ToDafny(*native_response))
}

func (shim *Shim) PutResourcePolicy(input ComAmazonawsDynamodbTypes.PutResourcePolicyInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.PutResourcePolicyInput_FromDafny(input)
	var native_response, native_error = shim.Client.PutResourcePolicy(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.PutResourcePolicyOutput_ToDafny(*native_response))
}

func (shim *Shim) Query(input ComAmazonawsDynamodbTypes.QueryInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.QueryInput_FromDafny(input)
	var native_response, native_error = shim.Client.Query(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.QueryOutput_ToDafny(*native_response))
}

func (shim *Shim) RestoreTableFromBackup(input ComAmazonawsDynamodbTypes.RestoreTableFromBackupInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.RestoreTableFromBackupInput_FromDafny(input)
	var native_response, native_error = shim.Client.RestoreTableFromBackup(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.RestoreTableFromBackupOutput_ToDafny(*native_response))
}

func (shim *Shim) RestoreTableToPointInTime(input ComAmazonawsDynamodbTypes.RestoreTableToPointInTimeInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.RestoreTableToPointInTimeInput_FromDafny(input)
	var native_response, native_error = shim.Client.RestoreTableToPointInTime(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.RestoreTableToPointInTimeOutput_ToDafny(*native_response))
}

func (shim *Shim) Scan(input ComAmazonawsDynamodbTypes.ScanInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.ScanInput_FromDafny(input)
	var native_response, native_error = shim.Client.Scan(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.ScanOutput_ToDafny(*native_response))
}

func (shim *Shim) TagResource(input ComAmazonawsDynamodbTypes.TagResourceInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.TagResourceInput_FromDafny(input)
	var _, native_error = shim.Client.TagResource(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(dafny.TupleOf())
}

func (shim *Shim) TransactGetItems(input ComAmazonawsDynamodbTypes.TransactGetItemsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.TransactGetItemsInput_FromDafny(input)
	var native_response, native_error = shim.Client.TransactGetItems(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.TransactGetItemsOutput_ToDafny(*native_response))
}

func (shim *Shim) TransactWriteItems(input ComAmazonawsDynamodbTypes.TransactWriteItemsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.TransactWriteItemsInput_FromDafny(input)
	var native_response, native_error = shim.Client.TransactWriteItems(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.TransactWriteItemsOutput_ToDafny(*native_response))
}

func (shim *Shim) UntagResource(input ComAmazonawsDynamodbTypes.UntagResourceInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UntagResourceInput_FromDafny(input)
	var _, native_error = shim.Client.UntagResource(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(dafny.TupleOf())
}

func (shim *Shim) UpdateContinuousBackups(input ComAmazonawsDynamodbTypes.UpdateContinuousBackupsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateContinuousBackupsInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateContinuousBackups(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateContinuousBackupsOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateContributorInsights(input ComAmazonawsDynamodbTypes.UpdateContributorInsightsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateContributorInsightsInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateContributorInsights(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateContributorInsightsOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateGlobalTable(input ComAmazonawsDynamodbTypes.UpdateGlobalTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateGlobalTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateGlobalTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateGlobalTableOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateGlobalTableSettings(input ComAmazonawsDynamodbTypes.UpdateGlobalTableSettingsInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateGlobalTableSettingsInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateGlobalTableSettings(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateGlobalTableSettingsOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateItem(input ComAmazonawsDynamodbTypes.UpdateItemInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateItemInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateItem(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateItemOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateKinesisStreamingDestination(input ComAmazonawsDynamodbTypes.UpdateKinesisStreamingDestinationInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateKinesisStreamingDestinationInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateKinesisStreamingDestination(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateKinesisStreamingDestinationOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateTable(input ComAmazonawsDynamodbTypes.UpdateTableInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateTableInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateTable(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateTableOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateTableReplicaAutoScaling(input ComAmazonawsDynamodbTypes.UpdateTableReplicaAutoScalingInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateTableReplicaAutoScalingInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateTableReplicaAutoScaling(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateTableReplicaAutoScalingOutput_ToDafny(*native_response))
}

func (shim *Shim) UpdateTimeToLive(input ComAmazonawsDynamodbTypes.UpdateTimeToLiveInput) Wrappers.Result {
	var native_request = comamazonawsdynamodbsmithygenerated.UpdateTimeToLiveInput_FromDafny(input)
	var native_response, native_error = shim.Client.UpdateTimeToLive(context.Background(), &native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(comamazonawsdynamodbsmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(comamazonawsdynamodbsmithygenerated.UpdateTimeToLiveOutput_ToDafny(*native_response))
}
