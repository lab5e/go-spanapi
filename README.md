# Go API client for spanapi

API for device, collection, output and firmware management

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.4.1 busy-janay
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://lab5e.com](https://lab5e.com)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import spanapi "github.com/lab5e/go-spanapi"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), spanapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), spanapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), spanapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), spanapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.lab5e.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BlobsApi* | [**DeleteBlob**](docs/BlobsApi.md#deleteblob) | **Delete** /span/collections/{collectionId}/blobs/{blobId} | Remove a blob stored on a collection
*BlobsApi* | [**ListBlobs**](docs/BlobsApi.md#listblobs) | **Get** /span/collections/{collectionId}/blobs | List the blobs for a collection
*CertificatesApi* | [**CreateCertificate**](docs/CertificatesApi.md#createcertificate) | **Post** /span/collections/{collectionId}/certificates/create | Create certificate
*CertificatesApi* | [**RetrieveCertificateChain**](docs/CertificatesApi.md#retrievecertificatechain) | **Get** /span/collections/{collectionId}/certificates | Get certificate chain
*CertificatesApi* | [**SignCertificate**](docs/CertificatesApi.md#signcertificate) | **Post** /span/collections/{collectionId}/certificates/sign | Sign certificate
*CertificatesApi* | [**VerifyCertificate**](docs/CertificatesApi.md#verifycertificate) | **Post** /span/collections/{collectionId}/certificates/verify | Verify certificate
*CollectionsApi* | [**CreateCollection**](docs/CollectionsApi.md#createcollection) | **Post** /span/collections | Create collection
*CollectionsApi* | [**DeleteCollection**](docs/CollectionsApi.md#deletecollection) | **Delete** /span/collections/{collectionId} | Delete collection
*CollectionsApi* | [**ListCollectionData**](docs/CollectionsApi.md#listcollectiondata) | **Get** /span/collections/{collectionId}/data | Retrieve data from devices
*CollectionsApi* | [**ListCollections**](docs/CollectionsApi.md#listcollections) | **Get** /span/collections | List collections
*CollectionsApi* | [**RetrieveCollection**](docs/CollectionsApi.md#retrievecollection) | **Get** /span/collections/{collectionId} | Retrieve collection
*CollectionsApi* | [**UpdateCollection**](docs/CollectionsApi.md#updatecollection) | **Patch** /span/collections/{collectionId} | Update collection
*DevicesApi* | [**AddDownstreamMessage**](docs/DevicesApi.md#adddownstreammessage) | **Post** /span/collections/{collectionId}/devices/{deviceId}/outbox | Add message to oubox
*DevicesApi* | [**CreateDevice**](docs/DevicesApi.md#createdevice) | **Post** /span/collections/{collectionId}/devices | Create device
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /span/collections/{collectionId}/devices/{deviceId} | Remove device.
*DevicesApi* | [**DeleteDownstreamMessage**](docs/DevicesApi.md#deletedownstreammessage) | **Delete** /span/collections/{collectionId}/devices/{deviceId}/outbox/{messageId} | Delete outgoing message
*DevicesApi* | [**DeviceCertificate**](docs/DevicesApi.md#devicecertificate) | **Get** /span/collections/{collectionId}/devices/{deviceId}/certs | Get issued certificate(s) for device
*DevicesApi* | [**ListDeviceData**](docs/DevicesApi.md#listdevicedata) | **Get** /span/collections/{collectionId}/devices/{deviceId}/data | Retrieve data from device
*DevicesApi* | [**ListDevices**](docs/DevicesApi.md#listdevices) | **Get** /span/collections/{collectionId}/devices | List devices in collection.
*DevicesApi* | [**ListDownstreamMessages**](docs/DevicesApi.md#listdownstreammessages) | **Get** /span/collections/{collectionId}/devices/{deviceId}/outbox | List the messages in the outbox
*DevicesApi* | [**ListUpstreamMessages**](docs/DevicesApi.md#listupstreammessages) | **Get** /span/collections/{collectionId}/devices/{deviceId}/inbox | List incoming messages
*DevicesApi* | [**RetrieveDevice**](docs/DevicesApi.md#retrievedevice) | **Get** /span/collections/{collectionId}/devices/{deviceId} | Retrieve device
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Patch** /span/collections/{existingCollectionId}/devices/{deviceId} | Update device
*FotaApi* | [**ClearFirmwareError**](docs/FotaApi.md#clearfirmwareerror) | **Delete** /span/collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
*FotaApi* | [**CreateFirmware**](docs/FotaApi.md#createfirmware) | **Post** /span/collections/{collectionId}/firmware | Create firmware
*FotaApi* | [**DeleteFirmware**](docs/FotaApi.md#deletefirmware) | **Delete** /span/collections/{collectionId}/firmware/{imageId} | Delete firmware
*FotaApi* | [**FirmwareUsage**](docs/FotaApi.md#firmwareusage) | **Get** /span/collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
*FotaApi* | [**ListFirmware**](docs/FotaApi.md#listfirmware) | **Get** /span/collections/{collectionId}/firmware | List firmware
*FotaApi* | [**RetrieveFirmware**](docs/FotaApi.md#retrievefirmware) | **Get** /span/collections/{collectionId}/firmware/{imageId} | Retrieve firmware
*FotaApi* | [**UpdateFirmware**](docs/FotaApi.md#updatefirmware) | **Patch** /span/collections/{existingCollectionId}/firmware/{imageId} | Update firmware
*GatewaysApi* | [**CreateGateway**](docs/GatewaysApi.md#creategateway) | **Post** /span/collections/{collectionId}/gateways | Create gateway
*GatewaysApi* | [**DeleteGateway**](docs/GatewaysApi.md#deletegateway) | **Delete** /span/collections/{collectionId}/gateways/{gatewayId} | Delete gateway
*GatewaysApi* | [**GatewayCertificates**](docs/GatewaysApi.md#gatewaycertificates) | **Get** /span/collections/{collectionId}/gateways/{gatewayId}/certs | Get issued certificate(s) for gateway
*GatewaysApi* | [**ListGateways**](docs/GatewaysApi.md#listgateways) | **Get** /span/collections/{collectionId}/gateways | List gateways
*GatewaysApi* | [**RetrieveGateway**](docs/GatewaysApi.md#retrievegateway) | **Get** /span/collections/{collectionId}/gateways/{gatewayId} | Retrieve gateway
*GatewaysApi* | [**UpdateGateway**](docs/GatewaysApi.md#updategateway) | **Patch** /span/collections/{existingCollectionId}/gateways/{gatewayId} | Update gateway
*OutputsApi* | [**CreateOutput**](docs/OutputsApi.md#createoutput) | **Post** /span/collections/{collectionId}/outputs | Create output
*OutputsApi* | [**DeleteOutput**](docs/OutputsApi.md#deleteoutput) | **Delete** /span/collections/{collectionId}/outputs/{outputId} | Delete output
*OutputsApi* | [**ListOutputs**](docs/OutputsApi.md#listoutputs) | **Get** /span/collections/{collectionId}/outputs | List outputs
*OutputsApi* | [**Logs**](docs/OutputsApi.md#logs) | **Get** /span/collections/{collectionId}/outputs/{outputId}/logs | Output logs
*OutputsApi* | [**RetrieveOutput**](docs/OutputsApi.md#retrieveoutput) | **Get** /span/collections/{collectionId}/outputs/{outputId} | Retrieve output
*OutputsApi* | [**Status**](docs/OutputsApi.md#status) | **Get** /span/collections/{collectionId}/outputs/{outputId}/status | Output status
*OutputsApi* | [**UpdateOutput**](docs/OutputsApi.md#updateoutput) | **Patch** /span/collections/{existingCollectionId}/outputs/{outputId} | Update output
*SpanApi* | [**GetSystemInfo**](docs/SpanApi.md#getsysteminfo) | **Get** /span/system | System information


## Documentation For Models

 - [AddDownstreamMessageRequest](docs/AddDownstreamMessageRequest.md)
 - [Any](docs/Any.md)
 - [Blob](docs/Blob.md)
 - [CellularIoTConfig](docs/CellularIoTConfig.md)
 - [CellularIoTMetadata](docs/CellularIoTMetadata.md)
 - [CertificateChainResponse](docs/CertificateChainResponse.md)
 - [CertificateInfo](docs/CertificateInfo.md)
 - [ClearFirmwareErrorResponse](docs/ClearFirmwareErrorResponse.md)
 - [CoAPMetadata](docs/CoAPMetadata.md)
 - [Collection](docs/Collection.md)
 - [CollectionFirmware](docs/CollectionFirmware.md)
 - [CreateCertificateRequest](docs/CreateCertificateRequest.md)
 - [CreateCertificateResponse](docs/CreateCertificateResponse.md)
 - [CreateCollectionRequest](docs/CreateCollectionRequest.md)
 - [CreateDeviceRequest](docs/CreateDeviceRequest.md)
 - [CreateFirmwareRequest](docs/CreateFirmwareRequest.md)
 - [CreateOutputRequest](docs/CreateOutputRequest.md)
 - [DeleteDownstreamMessageResponse](docs/DeleteDownstreamMessageResponse.md)
 - [Device](docs/Device.md)
 - [DeviceCertificateResponse](docs/DeviceCertificateResponse.md)
 - [DeviceConfig](docs/DeviceConfig.md)
 - [DeviceMetadata](docs/DeviceMetadata.md)
 - [Firmware](docs/Firmware.md)
 - [FirmwareManagement](docs/FirmwareManagement.md)
 - [FirmwareMetadata](docs/FirmwareMetadata.md)
 - [FirmwareUsageResponse](docs/FirmwareUsageResponse.md)
 - [Gateway](docs/Gateway.md)
 - [GatewayCIoTConfig](docs/GatewayCIoTConfig.md)
 - [GatewayCertificateResponse](docs/GatewayCertificateResponse.md)
 - [GatewayConfig](docs/GatewayConfig.md)
 - [GatewayCustomConfig](docs/GatewayCustomConfig.md)
 - [GatewayDeviceConfig](docs/GatewayDeviceConfig.md)
 - [GatewayDeviceMetadata](docs/GatewayDeviceMetadata.md)
 - [GatewayInetConfig](docs/GatewayInetConfig.md)
 - [GatewayMetadata](docs/GatewayMetadata.md)
 - [GatewayStatus](docs/GatewayStatus.md)
 - [GatewayType](docs/GatewayType.md)
 - [InetMetadata](docs/InetMetadata.md)
 - [InlineObject](docs/InlineObject.md)
 - [InlineObject1](docs/InlineObject1.md)
 - [ListBlobResponse](docs/ListBlobResponse.md)
 - [ListCollectionResponse](docs/ListCollectionResponse.md)
 - [ListDataResponse](docs/ListDataResponse.md)
 - [ListDevicesResponse](docs/ListDevicesResponse.md)
 - [ListDownstreamMessagesResponse](docs/ListDownstreamMessagesResponse.md)
 - [ListFirmwareResponse](docs/ListFirmwareResponse.md)
 - [ListGatewayResponse](docs/ListGatewayResponse.md)
 - [ListOutputResponse](docs/ListOutputResponse.md)
 - [ListUpstreamMessagesResponse](docs/ListUpstreamMessagesResponse.md)
 - [MQTTMetadata](docs/MQTTMetadata.md)
 - [MessageDownstream](docs/MessageDownstream.md)
 - [MessageState](docs/MessageState.md)
 - [MessageTransport](docs/MessageTransport.md)
 - [MessageUpstream](docs/MessageUpstream.md)
 - [NetworkMetadata](docs/NetworkMetadata.md)
 - [NetworkOperator](docs/NetworkOperator.md)
 - [Output](docs/Output.md)
 - [OutputConfig](docs/OutputConfig.md)
 - [OutputDataMessage](docs/OutputDataMessage.md)
 - [OutputLogEntry](docs/OutputLogEntry.md)
 - [OutputLogResponse](docs/OutputLogResponse.md)
 - [OutputMessageType](docs/OutputMessageType.md)
 - [OutputStatusResponse](docs/OutputStatusResponse.md)
 - [OutputType](docs/OutputType.md)
 - [RetrieveBlobResponse](docs/RetrieveBlobResponse.md)
 - [SignCertificateRequest](docs/SignCertificateRequest.md)
 - [SignCertificateResponse](docs/SignCertificateResponse.md)
 - [Status](docs/Status.md)
 - [SystemInfoResponse](docs/SystemInfoResponse.md)
 - [UDPMetadata](docs/UDPMetadata.md)
 - [UpdateCollectionRequest](docs/UpdateCollectionRequest.md)
 - [UpdateDeviceRequest](docs/UpdateDeviceRequest.md)
 - [UpdateFirmwareRequest](docs/UpdateFirmwareRequest.md)
 - [UpdateOutputRequest](docs/UpdateOutputRequest.md)
 - [VerifyCertificateRequest](docs/VerifyCertificateRequest.md)
 - [VerifyCertificateResponse](docs/VerifyCertificateResponse.md)


## Documentation For Authorization



### APIToken

- **Type**: API key
- **API key parameter name**: X-API-Token
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-API-Token and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

dev@lab5e.com

