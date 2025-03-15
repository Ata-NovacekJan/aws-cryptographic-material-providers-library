// Code generated by smithy-go-codegen DO NOT EDIT.

package awscryptographykeystoresmithygenerated

import (
	"unicode/utf8"

	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/dynamodb/DynamoDBwrapped"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/dynamodb/comamazonawsdynamodbsmithygenerated"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/kms/KMSwrapped"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/kms/comamazonawskmssmithygenerated"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/mpl/AwsCryptographyKeyStoreTypes"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/mpl/awscryptographykeystoresmithygeneratedtypes"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library/UTF8"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/smithy-dafny-standard-library/Wrappers"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/smithy-go"
	"github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"
)

func CreateKeyInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.CreateKeyInput) AwsCryptographyKeyStoreTypes.CreateKeyInput {

	return func() AwsCryptographyKeyStoreTypes.CreateKeyInput {

		return AwsCryptographyKeyStoreTypes.Companion_CreateKeyInput_.Create_CreateKeyInput_(Aws_cryptography_keyStore_CreateKeyInput_branchKeyIdentifier_ToDafny(nativeInput.BranchKeyIdentifier), Aws_cryptography_keyStore_CreateKeyInput_encryptionContext_ToDafny(nativeInput.EncryptionContext))
	}()

}

func CreateKeyOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.CreateKeyOutput) AwsCryptographyKeyStoreTypes.CreateKeyOutput {

	return func() AwsCryptographyKeyStoreTypes.CreateKeyOutput {

		return AwsCryptographyKeyStoreTypes.Companion_CreateKeyOutput_.Create_CreateKeyOutput_(Aws_cryptography_keyStore_CreateKeyOutput_branchKeyIdentifier_ToDafny(nativeOutput.BranchKeyIdentifier))
	}()

}

func CreateKeyStoreInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.CreateKeyStoreInput) AwsCryptographyKeyStoreTypes.CreateKeyStoreInput {

	return func() AwsCryptographyKeyStoreTypes.CreateKeyStoreInput {

		return AwsCryptographyKeyStoreTypes.Companion_CreateKeyStoreInput_.Create_CreateKeyStoreInput_()
	}()

}

func CreateKeyStoreOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.CreateKeyStoreOutput) AwsCryptographyKeyStoreTypes.CreateKeyStoreOutput {

	return func() AwsCryptographyKeyStoreTypes.CreateKeyStoreOutput {

		return AwsCryptographyKeyStoreTypes.Companion_CreateKeyStoreOutput_.Create_CreateKeyStoreOutput_(Aws_cryptography_keyStore_CreateKeyStoreOutput_tableArn_ToDafny(nativeOutput.TableArn))
	}()

}

func GetActiveBranchKeyInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.GetActiveBranchKeyInput) AwsCryptographyKeyStoreTypes.GetActiveBranchKeyInput {

	return func() AwsCryptographyKeyStoreTypes.GetActiveBranchKeyInput {

		return AwsCryptographyKeyStoreTypes.Companion_GetActiveBranchKeyInput_.Create_GetActiveBranchKeyInput_(Aws_cryptography_keyStore_GetActiveBranchKeyInput_branchKeyIdentifier_ToDafny(nativeInput.BranchKeyIdentifier))
	}()

}

func GetActiveBranchKeyOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.GetActiveBranchKeyOutput) AwsCryptographyKeyStoreTypes.GetActiveBranchKeyOutput {

	return func() AwsCryptographyKeyStoreTypes.GetActiveBranchKeyOutput {

		return AwsCryptographyKeyStoreTypes.Companion_GetActiveBranchKeyOutput_.Create_GetActiveBranchKeyOutput_(Aws_cryptography_keyStore_GetActiveBranchKeyOutput_branchKeyMaterials_ToDafny(nativeOutput.BranchKeyMaterials))
	}()

}

func GetBeaconKeyInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.GetBeaconKeyInput) AwsCryptographyKeyStoreTypes.GetBeaconKeyInput {

	return func() AwsCryptographyKeyStoreTypes.GetBeaconKeyInput {

		return AwsCryptographyKeyStoreTypes.Companion_GetBeaconKeyInput_.Create_GetBeaconKeyInput_(Aws_cryptography_keyStore_GetBeaconKeyInput_branchKeyIdentifier_ToDafny(nativeInput.BranchKeyIdentifier))
	}()

}

func GetBeaconKeyOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.GetBeaconKeyOutput) AwsCryptographyKeyStoreTypes.GetBeaconKeyOutput {

	return func() AwsCryptographyKeyStoreTypes.GetBeaconKeyOutput {

		return AwsCryptographyKeyStoreTypes.Companion_GetBeaconKeyOutput_.Create_GetBeaconKeyOutput_(Aws_cryptography_keyStore_GetBeaconKeyOutput_beaconKeyMaterials_ToDafny(nativeOutput.BeaconKeyMaterials))
	}()

}

func GetBranchKeyVersionInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.GetBranchKeyVersionInput) AwsCryptographyKeyStoreTypes.GetBranchKeyVersionInput {

	return func() AwsCryptographyKeyStoreTypes.GetBranchKeyVersionInput {

		return AwsCryptographyKeyStoreTypes.Companion_GetBranchKeyVersionInput_.Create_GetBranchKeyVersionInput_(Aws_cryptography_keyStore_GetBranchKeyVersionInput_branchKeyIdentifier_ToDafny(nativeInput.BranchKeyIdentifier), Aws_cryptography_keyStore_GetBranchKeyVersionInput_branchKeyVersion_ToDafny(nativeInput.BranchKeyVersion))
	}()

}

func GetBranchKeyVersionOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.GetBranchKeyVersionOutput) AwsCryptographyKeyStoreTypes.GetBranchKeyVersionOutput {

	return func() AwsCryptographyKeyStoreTypes.GetBranchKeyVersionOutput {

		return AwsCryptographyKeyStoreTypes.Companion_GetBranchKeyVersionOutput_.Create_GetBranchKeyVersionOutput_(Aws_cryptography_keyStore_GetBranchKeyVersionOutput_branchKeyMaterials_ToDafny(nativeOutput.BranchKeyMaterials))
	}()

}

func GetKeyStoreInfoOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.GetKeyStoreInfoOutput) AwsCryptographyKeyStoreTypes.GetKeyStoreInfoOutput {

	return func() AwsCryptographyKeyStoreTypes.GetKeyStoreInfoOutput {

		return AwsCryptographyKeyStoreTypes.Companion_GetKeyStoreInfoOutput_.Create_GetKeyStoreInfoOutput_(Aws_cryptography_keyStore_GetKeyStoreInfoOutput_keyStoreId_ToDafny(nativeOutput.KeyStoreId), Aws_cryptography_keyStore_GetKeyStoreInfoOutput_keyStoreName_ToDafny(nativeOutput.KeyStoreName), Aws_cryptography_keyStore_GetKeyStoreInfoOutput_logicalKeyStoreName_ToDafny(nativeOutput.LogicalKeyStoreName), Aws_cryptography_keyStore_GetKeyStoreInfoOutput_grantTokens_ToDafny(nativeOutput.GrantTokens), Aws_cryptography_keyStore_GetKeyStoreInfoOutput_kmsConfiguration_ToDafny(nativeOutput.KmsConfiguration))
	}()

}

func VersionKeyInput_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.VersionKeyInput) AwsCryptographyKeyStoreTypes.VersionKeyInput {

	return func() AwsCryptographyKeyStoreTypes.VersionKeyInput {

		return AwsCryptographyKeyStoreTypes.Companion_VersionKeyInput_.Create_VersionKeyInput_(Aws_cryptography_keyStore_VersionKeyInput_branchKeyIdentifier_ToDafny(nativeInput.BranchKeyIdentifier))
	}()

}

func VersionKeyOutput_ToDafny(nativeOutput awscryptographykeystoresmithygeneratedtypes.VersionKeyOutput) AwsCryptographyKeyStoreTypes.VersionKeyOutput {

	return func() AwsCryptographyKeyStoreTypes.VersionKeyOutput {

		return AwsCryptographyKeyStoreTypes.Companion_VersionKeyOutput_.Create_VersionKeyOutput_()
	}()

}

func KeyStoreException_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.KeyStoreException) AwsCryptographyKeyStoreTypes.Error {
	return func() AwsCryptographyKeyStoreTypes.Error {

		return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_KeyStoreException_(Aws_cryptography_keyStore_KeyStoreException_message_ToDafny(nativeInput.Message))
	}()

}

func CollectionOfErrors_Input_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.CollectionOfErrors) AwsCryptographyKeyStoreTypes.Error {
	var e []interface{}
	for _, i2 := range nativeInput.ListOfErrors {
		e = append(e, Error_ToDafny(i2))
	}
	return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_CollectionOfErrors_(dafny.SeqOf(e...), dafny.SeqOfChars([]dafny.Char(nativeInput.Message)...))
}
func OpaqueError_Input_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.OpaqueError) AwsCryptographyKeyStoreTypes.Error {
	return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_Opaque_(nativeInput.ErrObject)
}

func Error_ToDafny(err error) AwsCryptographyKeyStoreTypes.Error {
	switch err.(type) {
	// Service Errors
	case awscryptographykeystoresmithygeneratedtypes.KeyStoreException:
		return KeyStoreException_ToDafny(err.(awscryptographykeystoresmithygeneratedtypes.KeyStoreException))

	//DependentErrors
	case *smithy.OperationError:
		if err.(*smithy.OperationError).Service() == "DynamoDB" {
			DynamoDBError := comamazonawsdynamodbsmithygenerated.Error_ToDafny(err)
			return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_ComAmazonawsDynamodb_(DynamoDBError)
		}
		if err.(*smithy.OperationError).Service() == "KMS" {
			KMSError := comamazonawskmssmithygenerated.Error_ToDafny(err)
			return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_ComAmazonawsKms_(KMSError)
		}
		return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_Opaque_(err)

	case smithy.APIError:
		DynamoDBError := comamazonawsdynamodbsmithygenerated.Error_ToDafny(err)
		if !DynamoDBError.Is_OpaqueWithText() {
			return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_ComAmazonawsDynamodb_(DynamoDBError)
		}
		KMSError := comamazonawskmssmithygenerated.Error_ToDafny(err)
		if !KMSError.Is_OpaqueWithText() {
			return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_ComAmazonawsKms_(KMSError)
		}
		return AwsCryptographyKeyStoreTypes.Companion_Error_.Create_Opaque_(err)

	//Unmodelled Errors
	case awscryptographykeystoresmithygeneratedtypes.CollectionOfErrors:
		return CollectionOfErrors_Input_ToDafny(err.(awscryptographykeystoresmithygeneratedtypes.CollectionOfErrors))

	default:
		error, ok := err.(awscryptographykeystoresmithygeneratedtypes.OpaqueError)
		if !ok {
			panic("Error is not an OpaqueError")
		}
		return OpaqueError_Input_ToDafny(error)
	}
}

func KeyStoreConfig_ToDafny(nativeInput awscryptographykeystoresmithygeneratedtypes.KeyStoreConfig) AwsCryptographyKeyStoreTypes.KeyStoreConfig {
	return func() AwsCryptographyKeyStoreTypes.KeyStoreConfig {

		return AwsCryptographyKeyStoreTypes.Companion_KeyStoreConfig_.Create_KeyStoreConfig_(Aws_cryptography_keyStore_KeyStoreConfig_ddbTableName_ToDafny(nativeInput.DdbTableName), Aws_cryptography_keyStore_KeyStoreConfig_kmsConfiguration_ToDafny(nativeInput.KmsConfiguration), Aws_cryptography_keyStore_KeyStoreConfig_logicalKeyStoreName_ToDafny(nativeInput.LogicalKeyStoreName), Aws_cryptography_keyStore_KeyStoreConfig_id_ToDafny(nativeInput.Id), Aws_cryptography_keyStore_KeyStoreConfig_grantTokens_ToDafny(nativeInput.GrantTokens), Aws_cryptography_keyStore_KeyStoreConfig_ddbClient_ToDafny(nativeInput.DdbClient), Aws_cryptography_keyStore_KeyStoreConfig_kmsClient_ToDafny(nativeInput.KmsClient))
	}()

}

func Aws_cryptography_keyStore_CreateKeyInput_branchKeyIdentifier_ToDafny(input *string) Wrappers.Option {
	return func() Wrappers.Option {
		if input == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		return Wrappers.Companion_Option_.Create_Some_(func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(*input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}())
	}()
}

func Aws_cryptography_keyStore_CreateKeyInput_encryptionContext_ToDafny(input map[string]string) Wrappers.Option {
	return func() Wrappers.Option {
		fieldValue := dafny.NewMapBuilder()
		for key, val := range input {
			fieldValue.Add(Aws_cryptography_keyStore_EncryptionContext_key_ToDafny(key), Aws_cryptography_keyStore_EncryptionContext_value_ToDafny(val))
		}
		return Wrappers.Companion_Option_.Create_Some_(fieldValue.ToMap())
	}()
}

func Aws_cryptography_keyStore_EncryptionContext_key_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return dafny.SeqOf(func() []interface{} {
			if !utf8.ValidString(input) {
				panic("invalid utf8 input provided")
			}
			b := []byte(input)
			f := make([]interface{}, len(b))
			for i, v := range b {
				f[i] = v
			}
			return f
		}()...)
	}()
}

func Aws_cryptography_keyStore_EncryptionContext_value_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return dafny.SeqOf(func() []interface{} {
			if !utf8.ValidString(input) {
				panic("invalid utf8 input provided")
			}
			b := []byte(input)
			f := make([]interface{}, len(b))
			for i, v := range b {
				f[i] = v
			}
			return f
		}()...)
	}()
}

func Aws_cryptography_keyStore_CreateKeyOutput_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_CreateKeyStoreOutput_tableArn_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetActiveBranchKeyInput_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetActiveBranchKeyOutput_branchKeyMaterials_ToDafny(input awscryptographykeystoresmithygeneratedtypes.BranchKeyMaterials) AwsCryptographyKeyStoreTypes.BranchKeyMaterials {
	return func() AwsCryptographyKeyStoreTypes.BranchKeyMaterials {

		return AwsCryptographyKeyStoreTypes.Companion_BranchKeyMaterials_.Create_BranchKeyMaterials_(Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyIdentifier_ToDafny(input.BranchKeyIdentifier), Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyVersion_ToDafny(input.BranchKeyVersion), Aws_cryptography_keyStore_BranchKeyMaterials_encryptionContext_ToDafny(input.EncryptionContext), Aws_cryptography_keyStore_BranchKeyMaterials_branchKey_ToDafny(input.BranchKey))
	}()
}

func Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyVersion_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return dafny.SeqOf(func() []interface{} {
			if !utf8.ValidString(input) {
				panic("invalid utf8 input provided")
			}
			b := []byte(input)
			f := make([]interface{}, len(b))
			for i, v := range b {
				f[i] = v
			}
			return f
		}()...)
	}()
}

func Aws_cryptography_keyStore_BranchKeyMaterials_encryptionContext_ToDafny(input map[string]string) dafny.Map {
	return func() dafny.Map {
		fieldValue := dafny.NewMapBuilder()
		for key, val := range input {
			fieldValue.Add(Aws_cryptography_keyStore_EncryptionContext_key_ToDafny(key), Aws_cryptography_keyStore_EncryptionContext_value_ToDafny(val))
		}
		return fieldValue.ToMap()
	}()
}

func Aws_cryptography_keyStore_BranchKeyMaterials_branchKey_ToDafny(input []byte) dafny.Sequence {
	return func() dafny.Sequence {
		var v []interface{}
		if input == nil {
			return nil
		}
		for _, e := range input {
			v = append(v, e)
		}
		return dafny.SeqOf(v...)
	}()
}

func Aws_cryptography_keyStore_GetBeaconKeyInput_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetBeaconKeyOutput_beaconKeyMaterials_ToDafny(input awscryptographykeystoresmithygeneratedtypes.BeaconKeyMaterials) AwsCryptographyKeyStoreTypes.BeaconKeyMaterials {
	return func() AwsCryptographyKeyStoreTypes.BeaconKeyMaterials {

		return AwsCryptographyKeyStoreTypes.Companion_BeaconKeyMaterials_.Create_BeaconKeyMaterials_(Aws_cryptography_keyStore_BeaconKeyMaterials_beaconKeyIdentifier_ToDafny(input.BeaconKeyIdentifier), Aws_cryptography_keyStore_BeaconKeyMaterials_encryptionContext_ToDafny(input.EncryptionContext), Aws_cryptography_keyStore_BeaconKeyMaterials_beaconKey_ToDafny(input.BeaconKey), Aws_cryptography_keyStore_BeaconKeyMaterials_hmacKeys_ToDafny(input.HmacKeys))
	}()
}

func Aws_cryptography_keyStore_BeaconKeyMaterials_beaconKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_BeaconKeyMaterials_encryptionContext_ToDafny(input map[string]string) dafny.Map {
	return func() dafny.Map {
		fieldValue := dafny.NewMapBuilder()
		for key, val := range input {
			fieldValue.Add(Aws_cryptography_keyStore_EncryptionContext_key_ToDafny(key), Aws_cryptography_keyStore_EncryptionContext_value_ToDafny(val))
		}
		return fieldValue.ToMap()
	}()
}

func Aws_cryptography_keyStore_BeaconKeyMaterials_beaconKey_ToDafny(input []byte) Wrappers.Option {
	return func() Wrappers.Option {
		var v []interface{}
		if input == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		for _, e := range input {
			v = append(v, e)
		}
		return Wrappers.Companion_Option_.Create_Some_(dafny.SeqOf(v...))
	}()
}

func Aws_cryptography_keyStore_BeaconKeyMaterials_hmacKeys_ToDafny(input map[string][]byte) Wrappers.Option {
	return func() Wrappers.Option {
		fieldValue := dafny.NewMapBuilder()
		for key, val := range input {
			fieldValue.Add(Aws_cryptography_keyStore_HmacKeyMap_key_ToDafny(key), Aws_cryptography_keyStore_HmacKeyMap_value_ToDafny(val))
		}
		return Wrappers.Companion_Option_.Create_Some_(fieldValue.ToMap())
	}()
}

func Aws_cryptography_keyStore_HmacKeyMap_key_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_HmacKeyMap_value_ToDafny(input []byte) dafny.Sequence {
	return func() dafny.Sequence {
		var v []interface{}
		if input == nil {
			return nil
		}
		for _, e := range input {
			v = append(v, e)
		}
		return dafny.SeqOf(v...)
	}()
}

func Aws_cryptography_keyStore_GetBranchKeyVersionInput_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetBranchKeyVersionInput_branchKeyVersion_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetBranchKeyVersionOutput_branchKeyMaterials_ToDafny(input awscryptographykeystoresmithygeneratedtypes.BranchKeyMaterials) AwsCryptographyKeyStoreTypes.BranchKeyMaterials {
	return func() AwsCryptographyKeyStoreTypes.BranchKeyMaterials {

		return AwsCryptographyKeyStoreTypes.Companion_BranchKeyMaterials_.Create_BranchKeyMaterials_(Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyIdentifier_ToDafny(input.BranchKeyIdentifier), Aws_cryptography_keyStore_BranchKeyMaterials_branchKeyVersion_ToDafny(input.BranchKeyVersion), Aws_cryptography_keyStore_BranchKeyMaterials_encryptionContext_ToDafny(input.EncryptionContext), Aws_cryptography_keyStore_BranchKeyMaterials_branchKey_ToDafny(input.BranchKey))
	}()
}

func Aws_cryptography_keyStore_GetKeyStoreInfoOutput_keyStoreId_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetKeyStoreInfoOutput_keyStoreName_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetKeyStoreInfoOutput_logicalKeyStoreName_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetKeyStoreInfoOutput_grantTokens_ToDafny(input []string) dafny.Sequence {
	return func() dafny.Sequence {

		var fieldValue []interface{} = make([]interface{}, 0)
		for _, val := range input {
			element := Aws_cryptography_keyStore_GrantTokenList_member_ToDafny(val)
			fieldValue = append(fieldValue, element)
		}
		return dafny.SeqOf(fieldValue...)
	}()
}

func Aws_cryptography_keyStore_GrantTokenList_member_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_GetKeyStoreInfoOutput_kmsConfiguration_ToDafny(input awscryptographykeystoresmithygeneratedtypes.KMSConfiguration) AwsCryptographyKeyStoreTypes.KMSConfiguration {
	return func() AwsCryptographyKeyStoreTypes.KMSConfiguration {

		switch input.(type) {
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsKeyArn:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_kmsKeyArn_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsKeyArn).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_kmsKeyArn_(inputToConversion.UnwrapOr(nil).(dafny.Sequence))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsMRKeyArn:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_kmsMRKeyArn_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsMRKeyArn).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_kmsMRKeyArn_(inputToConversion.UnwrapOr(nil).(dafny.Sequence))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberdiscovery:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_discovery_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberdiscovery).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_discovery_(inputToConversion.UnwrapOr(nil).(AwsCryptographyKeyStoreTypes.Discovery))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMembermrDiscovery:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_mrDiscovery_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMembermrDiscovery).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_mrDiscovery_(inputToConversion.UnwrapOr(nil).(AwsCryptographyKeyStoreTypes.MRDiscovery))

		default:
			panic("Unhandled union type")
		}
	}()
}

func Aws_cryptography_keyStore_KMSConfiguration_kmsKeyArn_ToDafny(input string) Wrappers.Option {
	return func() Wrappers.Option {

		return Wrappers.Companion_Option_.Create_Some_(func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}())
	}()
}

func Aws_cryptography_keyStore_KMSConfiguration_kmsMRKeyArn_ToDafny(input string) Wrappers.Option {
	return func() Wrappers.Option {

		return Wrappers.Companion_Option_.Create_Some_(func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}())
	}()
}

func Aws_cryptography_keyStore_KMSConfiguration_discovery_ToDafny(input awscryptographykeystoresmithygeneratedtypes.Discovery) Wrappers.Option {
	return func() Wrappers.Option {

		return Wrappers.Companion_Option_.Create_Some_(AwsCryptographyKeyStoreTypes.Companion_Discovery_.Create_Discovery_())
	}()
}

func Aws_cryptography_keyStore_KMSConfiguration_mrDiscovery_ToDafny(input awscryptographykeystoresmithygeneratedtypes.MRDiscovery) Wrappers.Option {
	return func() Wrappers.Option {

		return Wrappers.Companion_Option_.Create_Some_(AwsCryptographyKeyStoreTypes.Companion_MRDiscovery_.Create_MRDiscovery_(Aws_cryptography_keyStore_MRDiscovery_region_ToDafny(input.Region)))
	}()
}

func Aws_cryptography_keyStore_MRDiscovery_region_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_VersionKeyInput_branchKeyIdentifier_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_KeyStoreException_message_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_ddbTableName_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_kmsConfiguration_ToDafny(input awscryptographykeystoresmithygeneratedtypes.KMSConfiguration) AwsCryptographyKeyStoreTypes.KMSConfiguration {
	return func() AwsCryptographyKeyStoreTypes.KMSConfiguration {

		switch input.(type) {
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsKeyArn:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_kmsKeyArn_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsKeyArn).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_kmsKeyArn_(inputToConversion.UnwrapOr(nil).(dafny.Sequence))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsMRKeyArn:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_kmsMRKeyArn_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberkmsMRKeyArn).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_kmsMRKeyArn_(inputToConversion.UnwrapOr(nil).(dafny.Sequence))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberdiscovery:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_discovery_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMemberdiscovery).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_discovery_(inputToConversion.UnwrapOr(nil).(AwsCryptographyKeyStoreTypes.Discovery))
		case *awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMembermrDiscovery:
			var inputToConversion = Aws_cryptography_keyStore_KMSConfiguration_mrDiscovery_ToDafny(input.(*awscryptographykeystoresmithygeneratedtypes.KMSConfigurationMembermrDiscovery).Value)
			return AwsCryptographyKeyStoreTypes.CompanionStruct_KMSConfiguration_{}.Create_mrDiscovery_(inputToConversion.UnwrapOr(nil).(AwsCryptographyKeyStoreTypes.MRDiscovery))

		default:
			panic("Unhandled union type")
		}
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_logicalKeyStoreName_ToDafny(input string) dafny.Sequence {
	return func() dafny.Sequence {

		return func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}()
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_id_ToDafny(input *string) Wrappers.Option {
	return func() Wrappers.Option {
		if input == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		return Wrappers.Companion_Option_.Create_Some_(func() dafny.Sequence {
			res, err := UTF8.DecodeFromNativeGoByteArray([]byte(*input))
			if err != nil {
				panic("invalid utf8 input provided")
			}
			return res
		}())
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_grantTokens_ToDafny(input []string) Wrappers.Option {
	return func() Wrappers.Option {
		if input == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		var fieldValue []interface{} = make([]interface{}, 0)
		for _, val := range input {
			element := Aws_cryptography_keyStore_GrantTokenList_member_ToDafny(val)
			fieldValue = append(fieldValue, element)
		}
		return Wrappers.Companion_Option_.Create_Some_(dafny.SeqOf(fieldValue...))
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_ddbClient_ToDafny(input *dynamodb.Client) Wrappers.Option {
	return func() Wrappers.Option {
		if (input) == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		return Wrappers.Companion_Option_.Create_Some_(&DynamoDBwrapped.Shim{Client: input})
	}()
}

func Aws_cryptography_keyStore_KeyStoreConfig_kmsClient_ToDafny(input *kms.Client) Wrappers.Option {
	return func() Wrappers.Option {
		if (input) == nil {
			return Wrappers.Companion_Option_.Create_None_()
		}
		return Wrappers.Companion_Option_.Create_Some_(&KMSwrapped.Shim{Client: input})
	}()
}
