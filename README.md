# Go API client for spanapi

API for device, collection, output and firmware management

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 4.1.16 spooky-devante
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://lab5e.com/span](https://lab5e.com/span)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./spanapi"
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
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.lab5e.com/span*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CollectionsApi* | [**BroadcastMessage**](docs/CollectionsApi.md#broadcastmessage) | **Post** /collections/{collectionId}/to | Broadcast message
*CollectionsApi* | [**CreateCollection**](docs/CollectionsApi.md#createcollection) | **Post** /collections | Create collection
*CollectionsApi* | [**DeleteCollection**](docs/CollectionsApi.md#deletecollection) | **Delete** /collections/{collectionId} | Delete collection
*CollectionsApi* | [**ListCollectionData**](docs/CollectionsApi.md#listcollectiondata) | **Get** /collections/{collectionId}/data | Get payloads
*CollectionsApi* | [**ListCollections**](docs/CollectionsApi.md#listcollections) | **Get** /collections | List collections
*CollectionsApi* | [**RetrieveCollection**](docs/CollectionsApi.md#retrievecollection) | **Get** /collections/{collectionId} | Retrieve collection
*CollectionsApi* | [**UpdateCollection**](docs/CollectionsApi.md#updatecollection) | **Patch** /collections/{collectionId} | Update collection
*DevicesApi* | [**CreateDevice**](docs/DevicesApi.md#createdevice) | **Post** /collections/{collectionId}/devices | Create device
*DevicesApi* | [**DeleteDevice**](docs/DevicesApi.md#deletedevice) | **Delete** /collections/{collectionId}/devices/{deviceId} | Remove device.
*DevicesApi* | [**ListDeviceData**](docs/DevicesApi.md#listdevicedata) | **Get** /collections/{collectionId}/devices/{deviceId}/data | Get payloads
*DevicesApi* | [**ListDevices**](docs/DevicesApi.md#listdevices) | **Get** /collections/{collectionId}/devices | List devices in collection.
*DevicesApi* | [**RetrieveDevice**](docs/DevicesApi.md#retrievedevice) | **Get** /collections/{collectionId}/devices/{deviceId} | Retrieve device
*DevicesApi* | [**SendMessage**](docs/DevicesApi.md#sendmessage) | **Post** /collections/{collectionId}/devices/{deviceId}/to | Send message to a device.
*DevicesApi* | [**UpdateDevice**](docs/DevicesApi.md#updatedevice) | **Patch** /collections/{existingCollectionId}/devices/{deviceId} | Update device. The device can be moved from one collection to another by setting the collection ID field to the new collection. You must have administrative access to both collections.
*FotaApi* | [**ClearFirmwareError**](docs/FotaApi.md#clearfirmwareerror) | **Delete** /collections/{collectionId}/devices/{deviceId}/fwerror | Clear FOTA error
*FotaApi* | [**CreateFirmware**](docs/FotaApi.md#createfirmware) | **Post** /collections/{collectionId}/firmware | Create firmware
*FotaApi* | [**DeleteFirmware**](docs/FotaApi.md#deletefirmware) | **Delete** /collections/{collectionId}/firmware/{imageId} | Delete firmware
*FotaApi* | [**FirmwareUsage**](docs/FotaApi.md#firmwareusage) | **Get** /collections/{collectionId}/firmware/{imageId}/usage | Firmware usage
*FotaApi* | [**ListFirmware**](docs/FotaApi.md#listfirmware) | **Get** /collections/{collectionId}/firmware | List firmware
*FotaApi* | [**RetrieveFirmware**](docs/FotaApi.md#retrievefirmware) | **Get** /collections/{collectionId}/firmware/{imageId} | Retrieve firmware
*FotaApi* | [**UpdateFirmware**](docs/FotaApi.md#updatefirmware) | **Patch** /collections/{collectionId}/firmware/{imageId} | Update firmware. Only the version and tags fields can be updated. The other fields will be ignored..
*OutputsApi* | [**CreateOutput**](docs/OutputsApi.md#createoutput) | **Post** /collections/{collectionId}/outputs | Create output
*OutputsApi* | [**DeleteOutput**](docs/OutputsApi.md#deleteoutput) | **Delete** /collections/{collectionId}/outputs/{outputId} | Delete output
*OutputsApi* | [**ListOutputs**](docs/OutputsApi.md#listoutputs) | **Get** /collections/{collectionId}/outputs | List outputs
*OutputsApi* | [**Logs**](docs/OutputsApi.md#logs) | **Get** /collections/{collectionId}/outputs/{outputId}/logs | Output logs
*OutputsApi* | [**RetrieveOutput**](docs/OutputsApi.md#retrieveoutput) | **Get** /collections/{collectionId}/outputs/{outputId} | Retrieve output
*OutputsApi* | [**Status**](docs/OutputsApi.md#status) | **Get** /collections/{collectionId}/outputs/{outputId}/status | Output status
*OutputsApi* | [**UpdateOutput**](docs/OutputsApi.md#updateoutput) | **Patch** /collections/{collectionId}/outputs/{outputId} | Update output
*SystemApi* | [**GetSystemInfo**](docs/SystemApi.md#getsysteminfo) | **Get** /system | System information


## Documentation For Models

 - [BroadcastMessageRequest](docs/BroadcastMessageRequest.md)
 - [ClearFirmwareErrorResponse](docs/ClearFirmwareErrorResponse.md)
 - [CoAPMetadata](docs/CoAPMetadata.md)
 - [Collection](docs/Collection.md)
 - [CollectionFirmware](docs/CollectionFirmware.md)
 - [CollectionFirmwareFirmwareManagement](docs/CollectionFirmwareFirmwareManagement.md)
 - [CreateFirmwareRequest](docs/CreateFirmwareRequest.md)
 - [Device](docs/Device.md)
 - [DeviceMetadata](docs/DeviceMetadata.md)
 - [FieldMask](docs/FieldMask.md)
 - [Firmware](docs/Firmware.md)
 - [FirmwareMetadata](docs/FirmwareMetadata.md)
 - [FirmwareUsageResponse](docs/FirmwareUsageResponse.md)
 - [ListCollectionResponse](docs/ListCollectionResponse.md)
 - [ListDataResponse](docs/ListDataResponse.md)
 - [ListDevicesResponse](docs/ListDevicesResponse.md)
 - [ListFirmwareResponse](docs/ListFirmwareResponse.md)
 - [ListOutputResponse](docs/ListOutputResponse.md)
 - [MessageSendResult](docs/MessageSendResult.md)
 - [MultiSendMessageResponse](docs/MultiSendMessageResponse.md)
 - [NetworkMetadata](docs/NetworkMetadata.md)
 - [NetworkOperator](docs/NetworkOperator.md)
 - [Output](docs/Output.md)
 - [OutputConfig](docs/OutputConfig.md)
 - [OutputDataMessage](docs/OutputDataMessage.md)
 - [OutputDataMessageOutputMessageType](docs/OutputDataMessageOutputMessageType.md)
 - [OutputLogEntry](docs/OutputLogEntry.md)
 - [OutputLogResponse](docs/OutputLogResponse.md)
 - [OutputStatusResponse](docs/OutputStatusResponse.md)
 - [OutputType](docs/OutputType.md)
 - [ProtobufAny](docs/ProtobufAny.md)
 - [RpcStatus](docs/RpcStatus.md)
 - [SendMessageRequest](docs/SendMessageRequest.md)
 - [SendMessageResponse](docs/SendMessageResponse.md)
 - [SystemInfoResponse](docs/SystemInfoResponse.md)
 - [UDPMetadata](docs/UDPMetadata.md)
 - [UpdateDeviceRequest](docs/UpdateDeviceRequest.md)


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

