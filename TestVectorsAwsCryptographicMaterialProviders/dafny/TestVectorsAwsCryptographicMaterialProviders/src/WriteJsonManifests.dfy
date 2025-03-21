// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

include "JSONHelpers.dfy"
include "VectorsComposition/AllAlgorithmSuites.dfy"
include "TestVectors.dfy"

module {:options "-functionSyntax:4"} WriteJsonManifests {
  import opened JSON.Values
  import AllAlgorithmSuites
  import AllAlgorithmSuites.Types
  import opened Wrappers
  import UTF8
  import SortedSets
  import Seq
  import StandardLibrary
  import Base64

  import TestVectors
  import KeyDescription

  function EncryptionContextKeysToJson(keys: Option<Types.EncryptionContextKeys>)
    : Result<seq<(string, JSON)>, string>
  {
    if keys.Some? then

      var tmp :- Seq.MapWithResult(
                   bytes =>
                     var key :- UTF8.Decode(bytes);
                     Success(String(key)),
                   keys.value);

      Success([("requiredEncryptionContextKeys", Array(tmp))])

    else
      Success([])
  }

  function EncryptionContextToJson(key: string, m: Types.EncryptionContext)
    : Result<seq<(string, JSON)>, string>
  {
    var keys := SortedSets.ComputeSetToOrderedSequence2(m.Keys, (a, b) => a < b);
    var pairsBytes
      :- Seq.MapWithResult(
           k requires k in m.Keys =>
             var key :- UTF8.Decode(k);
             var value :- UTF8.Decode(m[k]);
             Success((key, String(value))),
           keys);
    Success([(key, Object(pairsBytes))])
  }

  function printJson(j: JSON) : (){()} by method {print j, "\n", "\n"; return ();}

  function EncryptTestVectorToJson(
    test: TestVectors.EncryptTestVector
  ): Result<JSON, string>
  {
    var id := AllAlgorithmSuites.ToHex(test.algorithmSuite);
    var description := test.name + " " + id;

    var encryptionContext :- EncryptionContextToJson("encryptionContext", test.encryptionContext);
    var maxPlaintextL
      := if test.maxPlaintextLength.Some? then
           [("maxPlaintextLength",
             Number(Int(test.maxPlaintextLength.value as int)))]
         else
           [];
    var requiredEncryptionContextKeys :- EncryptionContextKeysToJson(test.requiredEncryptionContextKeys);
    var reproducedEc
      :- if
           && !test.NegativeEncryptKeyringVector?
           && test.reproducedEncryptionContext.Some?
         then
           EncryptionContextToJson("reproducedEncryptionContext", test.reproducedEncryptionContext.value)
         else
           Success([]);
    var optionalValues := reproducedEc + maxPlaintextL + requiredEncryptionContextKeys + encryptionContext;

    match test
    case PositiveEncryptKeyringVector(_,_,_,_,_,_,_,_,_,_) =>
      var encrypt :- KeyDescription.ToJson(test.encryptDescription, 3);
      var decrypt :- KeyDescription.ToJson(test.decryptDescription, 3);
      Success(Object([
                       ("type", String("positive-keyring")),
                       ("description", String(description)),
                       ("algorithmSuiteId", String(id)),
                       ("encryptKeyDescription", encrypt),
                       ("decryptKeyDescription", decrypt)
                     ] + optionalValues))
    case PositiveEncryptNegativeDecryptKeyringVector(_,_,_,_,_,_,_,_,_,_,_) =>
      var encrypt :- KeyDescription.ToJson(test.encryptDescription, 3);
      var decrypt :- KeyDescription.ToJson(test.decryptDescription, 3);
      Success(Object([
                       ("type", String("negative-decrypt-keyring")),
                       ("description", String(description)),
                       ("decryptErrorDescription", String(test.decryptErrorDescription)),
                       ("algorithmSuiteId", String(id)),
                       ("encryptKeyDescription", encrypt),
                       ("decryptKeyDescription", decrypt)
                     ] + optionalValues))
    case NegativeEncryptKeyringVector(_,_,_,_,_,_,_,_,_) =>
      var keyDescription :- KeyDescription.ToJson(test.keyDescription, 3);
      Success(Object([
                       ("type", String("negative-encrypt-keyring")),
                       ("description", String(description)),
                       ("errorDescription", String(test.errorDescription)),
                       ("algorithmSuiteId", String(id)),
                       ("keyDescription", keyDescription)
                     ] + optionalValues))
  }

  function OptionalBytes(key: string, secret: Option<Types.Secret>)
    : seq<(string, JSON)>
  {
    if secret.Some? then
      var base64 := Base64.Encode(secret.value);
      [(key, String(base64))]
    else
      []
  }

  function EncryptedDataKey(
    encryptedDataKey: Types.EncryptedDataKey
  )
    : Result<JSON, string>
  {
    var keyProviderId :- UTF8.Decode(encryptedDataKey.keyProviderId);
    Success(Object([
                     ("keyProviderId", String(keyProviderId)),
                     ("keyProviderInfo", String(Base64.Encode(encryptedDataKey.keyProviderInfo))),
                     ("ciphertext", String(Base64.Encode(encryptedDataKey.ciphertext)))
                   ]))
  }

  function DecryptTestVectorToJson(
    test: TestVectors.DecryptTestVector
  ): Result<JSON, string>
  {
    var id := AllAlgorithmSuites.ToHex(test.algorithmSuite);
    var description := test.name + " " + id;

    var encryptionContext :- EncryptionContextToJson("encryptionContext", test.encryptionContext);
    var reproducedEc
      :- if test.reproducedEncryptionContext.Some? then
           EncryptionContextToJson("reproducedEncryptionContext", test.reproducedEncryptionContext.value)
         else
           Success([]);
    var keyDescription :- KeyDescription.ToJson(test.keyDescription, 3);
    var encryptedDataKeys :- Seq.MapWithResult(edk => EncryptedDataKey(edk), test.encryptedDataKeys);

    match test
    case PositiveDecryptKeyringTest(_,_,_,_,_,_,_,_,_) =>
      var plaintextDataKey := OptionalBytes("plaintextDataKey", test.expectedResult.plaintextDataKey);
      var symmetricSigningKey := OptionalBytes("symmetricSigningKey", test.expectedResult.symmetricSigningKey);
      var requiredEncryptionContextKeys :- EncryptionContextKeysToJson(Some(test.expectedResult.requiredEncryptionContextKeys));
      Success(Object([
                       ("type", String("positive-keyring")),
                       ("description", String(description)),
                       ("algorithmSuiteId", String(id)),
                       ("keyDescription", keyDescription),
                       ("encryptedDataKeys", Array(encryptedDataKeys)),
                       ("result", Object(
                        plaintextDataKey + symmetricSigningKey + requiredEncryptionContextKeys
                        ))
                     ] + reproducedEc + encryptionContext))
    case NegativeDecryptKeyringTest(_,_,_,_,_,_,_,_,_) =>
      Success(Object([
                       ("type", String("negative-keyring")),
                       ("description", String(description)),
                       ("errorDescription", String(test.errorDescription)),
                       ("algorithmSuiteId", String(id)),
                       ("keyDescription", keyDescription),
                       ("encryptedDataKeys", Array(encryptedDataKeys))
                     ] + reproducedEc + encryptionContext))
  }
}
