// Code generated by smithy-go-codegen DO NOT EDIT.

package awscryptographyprimitivessmithygenerated

import (
	"context"

	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/primitives/AtomicPrimitives"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/primitives/AwsCryptographyPrimitivesTypes"
	"github.com/aws/aws-cryptographic-material-providers-library/releases/go/primitives/awscryptographyprimitivessmithygeneratedtypes"
	"github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"
)

type Client struct {
	DafnyClient AwsCryptographyPrimitivesTypes.IAwsCryptographicPrimitivesClient
}

func NewClient(clientConfig awscryptographyprimitivessmithygeneratedtypes.CryptoConfig) (*Client, error) {
	var dafnyConfig = CryptoConfig_ToDafny(clientConfig)
	var dafny_response = AtomicPrimitives.Companion_Default___.AtomicPrimitives(dafnyConfig)
	if dafny_response.Is_Failure() {
		panic("Client construction failed. This should never happen")
	}
	var dafnyClient = dafny_response.Extract().(AwsCryptographyPrimitivesTypes.IAwsCryptographicPrimitivesClient)
	client := &Client{dafnyClient}
	return client, nil
}

func (client *Client) GenerateRandomBytes(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GenerateRandomBytesInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GenerateRandomBytesInput = GenerateRandomBytesInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GenerateRandomBytes(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_GenerateRandomBytesOutput_data_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) Digest(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.DigestInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.DigestInput = DigestInput_ToDafny(params)
	var dafny_response = client.DafnyClient.Digest(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_DigestOutput_digest_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) HMac(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.HMacInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.HMacInput = HMacInput_ToDafny(params)
	var dafny_response = client.DafnyClient.HMac(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_HMacOutput_digest_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) HkdfExtract(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.HkdfExtractInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.HkdfExtractInput = HkdfExtractInput_ToDafny(params)
	var dafny_response = client.DafnyClient.HkdfExtract(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_HkdfExtractOutput_prk_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) HkdfExpand(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.HkdfExpandInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.HkdfExpandInput = HkdfExpandInput_ToDafny(params)
	var dafny_response = client.DafnyClient.HkdfExpand(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_HkdfExpandOutput_okm_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) Hkdf(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.HkdfInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.HkdfInput = HkdfInput_ToDafny(params)
	var dafny_response = client.DafnyClient.Hkdf(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_HkdfOutput_okm_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) KdfCounterMode(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.KdfCtrInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.KdfCtrInput = KdfCtrInput_ToDafny(params)
	var dafny_response = client.DafnyClient.KdfCounterMode(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_KdfCtrOutput_okm_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) AesKdfCounterMode(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.AesKdfCtrInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.AesKdfCtrInput = AesKdfCtrInput_ToDafny(params)
	var dafny_response = client.DafnyClient.AesKdfCounterMode(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_AesKdfCtrOutput_okm_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) AESEncrypt(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.AESEncryptInput) (*awscryptographyprimitivessmithygeneratedtypes.AESEncryptOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.AESEncryptInput = AESEncryptInput_ToDafny(params)
	var dafny_response = client.DafnyClient.AESEncrypt(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = AESEncryptOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.AESEncryptOutput))
	return &native_response, nil

}

func (client *Client) AESDecrypt(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.AESDecryptInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.AESDecryptInput = AESDecryptInput_ToDafny(params)
	var dafny_response = client.DafnyClient.AESDecrypt(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_AESDecryptOutput_plaintext_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) GenerateRSAKeyPair(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GenerateRSAKeyPairInput) (*awscryptographyprimitivessmithygeneratedtypes.GenerateRSAKeyPairOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GenerateRSAKeyPairInput = GenerateRSAKeyPairInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GenerateRSAKeyPair(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = GenerateRSAKeyPairOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.GenerateRSAKeyPairOutput))
	return &native_response, nil

}

func (client *Client) GetRSAKeyModulusLength(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GetRSAKeyModulusLengthInput) (*awscryptographyprimitivessmithygeneratedtypes.GetRSAKeyModulusLengthOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GetRSAKeyModulusLengthInput = GetRSAKeyModulusLengthInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GetRSAKeyModulusLength(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = GetRSAKeyModulusLengthOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.GetRSAKeyModulusLengthOutput))
	return &native_response, nil

}

func (client *Client) RSADecrypt(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.RSADecryptInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.RSADecryptInput = RSADecryptInput_ToDafny(params)
	var dafny_response = client.DafnyClient.RSADecrypt(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_RSADecryptOutput_plaintext_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) RSAEncrypt(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.RSAEncryptInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.RSAEncryptInput = RSAEncryptInput_ToDafny(params)
	var dafny_response = client.DafnyClient.RSAEncrypt(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_RSAEncryptOutput_cipherText_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) GenerateECDSASignatureKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GenerateECDSASignatureKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.GenerateECDSASignatureKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GenerateECDSASignatureKeyInput = GenerateECDSASignatureKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GenerateECDSASignatureKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = GenerateECDSASignatureKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.GenerateECDSASignatureKeyOutput))
	return &native_response, nil

}

func (client *Client) ECDSASign(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.ECDSASignInput) ([]byte, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal []byte
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.ECDSASignInput = ECDSASignInput_ToDafny(params)
	var dafny_response = client.DafnyClient.ECDSASign(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal []byte
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_ECDSASignOutput_signature_FromDafny(dafny_response.Dtor_value().(dafny.Sequence))
	return native_response, nil

}

func (client *Client) ECDSAVerify(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.ECDSAVerifyInput) (bool, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		var defaultVal bool
		return defaultVal, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.ECDSAVerifyInput = ECDSAVerifyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.ECDSAVerify(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		var defaultVal bool
		return defaultVal, Error_FromDafny(err)
	}
	var native_response = Aws_cryptography_primitives_ECDSAVerifyOutput_success_FromDafny(dafny_response.Dtor_value().(bool))
	return native_response, nil

}

func (client *Client) GenerateECCKeyPair(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GenerateECCKeyPairInput) (*awscryptographyprimitivessmithygeneratedtypes.GenerateECCKeyPairOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GenerateECCKeyPairInput = GenerateECCKeyPairInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GenerateECCKeyPair(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = GenerateECCKeyPairOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.GenerateECCKeyPairOutput))
	return &native_response, nil

}

func (client *Client) GetPublicKeyFromPrivateKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.GetPublicKeyFromPrivateKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.GetPublicKeyFromPrivateKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.GetPublicKeyFromPrivateKeyInput = GetPublicKeyFromPrivateKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.GetPublicKeyFromPrivateKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = GetPublicKeyFromPrivateKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.GetPublicKeyFromPrivateKeyOutput))
	return &native_response, nil

}

func (client *Client) ValidatePublicKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.ValidatePublicKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.ValidatePublicKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.ValidatePublicKeyInput = ValidatePublicKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.ValidatePublicKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = ValidatePublicKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.ValidatePublicKeyOutput))
	return &native_response, nil

}

func (client *Client) DeriveSharedSecret(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.DeriveSharedSecretInput) (*awscryptographyprimitivessmithygeneratedtypes.DeriveSharedSecretOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.DeriveSharedSecretInput = DeriveSharedSecretInput_ToDafny(params)
	var dafny_response = client.DafnyClient.DeriveSharedSecret(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = DeriveSharedSecretOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.DeriveSharedSecretOutput))
	return &native_response, nil

}

func (client *Client) CompressPublicKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.CompressPublicKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.CompressPublicKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.CompressPublicKeyInput = CompressPublicKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.CompressPublicKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = CompressPublicKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.CompressPublicKeyOutput))
	return &native_response, nil

}

func (client *Client) DecompressPublicKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.DecompressPublicKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.DecompressPublicKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.DecompressPublicKeyInput = DecompressPublicKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.DecompressPublicKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = DecompressPublicKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.DecompressPublicKeyOutput))
	return &native_response, nil

}

func (client *Client) ParsePublicKey(ctx context.Context, params awscryptographyprimitivessmithygeneratedtypes.ParsePublicKeyInput) (*awscryptographyprimitivessmithygeneratedtypes.ParsePublicKeyOutput, error) {
	err := params.Validate()
	if err != nil {
		opaqueErr := awscryptographyprimitivessmithygeneratedtypes.OpaqueError{
			ErrObject: err,
		}
		return nil, opaqueErr
	}

	var dafny_request AwsCryptographyPrimitivesTypes.ParsePublicKeyInput = ParsePublicKeyInput_ToDafny(params)
	var dafny_response = client.DafnyClient.ParsePublicKey(dafny_request)

	if dafny_response.Is_Failure() {
		err := dafny_response.Dtor_error().(AwsCryptographyPrimitivesTypes.Error)
		return nil, Error_FromDafny(err)
	}
	var native_response = ParsePublicKeyOutput_FromDafny(dafny_response.Dtor_value().(AwsCryptographyPrimitivesTypes.ParsePublicKeyOutput))
	return &native_response, nil

}
