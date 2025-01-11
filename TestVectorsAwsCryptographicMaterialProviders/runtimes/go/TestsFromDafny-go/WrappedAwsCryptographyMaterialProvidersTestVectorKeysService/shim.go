// Code generated by smithy-go-codegen DO NOT EDIT.

package WrappedAwsCryptographyMaterialProvidersTestVectorKeysService

import (
	"context"

	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/mpl/awscryptographymaterialproviderssmithygenerated"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library/Wrappers"
	"github.com/aws/aws-cryptographic-material-providers-library/testvectors/AwsCryptographyMaterialProvidersTestVectorKeysTypes"
	"github.com/aws/aws-cryptographic-material-providers-library/testvectors/awscryptographymaterialproviderstestvectorkeyssmithygenerated"
)

type Shim struct {
	AwsCryptographyMaterialProvidersTestVectorKeysTypes.IKeyVectorsClient
	client *awscryptographymaterialproviderstestvectorkeyssmithygenerated.Client
}

func WrappedKeyVectors(inputConfig AwsCryptographyMaterialProvidersTestVectorKeysTypes.KeyVectorsConfig) Wrappers.Result {
	var nativeConfig = awscryptographymaterialproviderstestvectorkeyssmithygenerated.KeyVectorsConfig_FromDafny(inputConfig)
	var nativeClient, nativeError = awscryptographymaterialproviderstestvectorkeyssmithygenerated.NewClient(nativeConfig)
	if nativeError != nil {
		return Wrappers.Companion_Result_.Create_Failure_(AwsCryptographyMaterialProvidersTestVectorKeysTypes.Companion_Error_.Create_Opaque_(nativeError))
	}
	return Wrappers.Companion_Result_.Create_Success_(&Shim{client: nativeClient})
}

func (shim *Shim) CreateTestVectorKeyring(input AwsCryptographyMaterialProvidersTestVectorKeysTypes.TestVectorKeyringInput) Wrappers.Result {
	var native_request = awscryptographymaterialproviderstestvectorkeyssmithygenerated.TestVectorKeyringInput_FromDafny(input)
	var native_response, native_error = shim.client.CreateTestVectorKeyring(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(awscryptographymaterialproviderssmithygenerated.Keyring_ToDafny(native_response))
}

func (shim *Shim) CreateWrappedTestVectorKeyring(input AwsCryptographyMaterialProvidersTestVectorKeysTypes.TestVectorKeyringInput) Wrappers.Result {
	var native_request = awscryptographymaterialproviderstestvectorkeyssmithygenerated.TestVectorKeyringInput_FromDafny(input)
	var native_response, native_error = shim.client.CreateWrappedTestVectorKeyring(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(awscryptographymaterialproviderssmithygenerated.Keyring_ToDafny(native_response))
}

func (shim *Shim) CreateWrappedTestVectorCmm(input AwsCryptographyMaterialProvidersTestVectorKeysTypes.TestVectorCmmInput) Wrappers.Result {
	var native_request = awscryptographymaterialproviderstestvectorkeyssmithygenerated.TestVectorCmmInput_FromDafny(input)
	var native_response, native_error = shim.client.CreateWrappedTestVectorCmm(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(awscryptographymaterialproviderssmithygenerated.CryptographicMaterialsManager_ToDafny(native_response))
}

func (shim *Shim) GetKeyDescription(input AwsCryptographyMaterialProvidersTestVectorKeysTypes.GetKeyDescriptionInput) Wrappers.Result {
	var native_request = awscryptographymaterialproviderstestvectorkeyssmithygenerated.GetKeyDescriptionInput_FromDafny(input)
	var native_response, native_error = shim.client.GetKeyDescription(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.GetKeyDescriptionOutput_ToDafny(*native_response))
}

func (shim *Shim) SerializeKeyDescription(input AwsCryptographyMaterialProvidersTestVectorKeysTypes.SerializeKeyDescriptionInput) Wrappers.Result {
	var native_request = awscryptographymaterialproviderstestvectorkeyssmithygenerated.SerializeKeyDescriptionInput_FromDafny(input)
	var native_response, native_error = shim.client.SerializeKeyDescription(context.Background(), native_request)
	if native_error != nil {
		return Wrappers.Companion_Result_.Create_Failure_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.Error_ToDafny(native_error))
	}
	return Wrappers.Companion_Result_.Create_Success_(awscryptographymaterialproviderstestvectorkeyssmithygenerated.SerializeKeyDescriptionOutput_ToDafny(*native_response))
}
