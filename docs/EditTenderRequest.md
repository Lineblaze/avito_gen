# EditTenderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Полное название тендера | [optional] 
**Description** | Pointer to **string** | Описание тендера | [optional] 
**ServiceType** | Pointer to [**TenderServiceType**](TenderServiceType.md) |  | [optional] 

## Methods

### NewEditTenderRequest

`func NewEditTenderRequest() *EditTenderRequest`

NewEditTenderRequest instantiates a new EditTenderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditTenderRequestWithDefaults

`func NewEditTenderRequestWithDefaults() *EditTenderRequest`

NewEditTenderRequestWithDefaults instantiates a new EditTenderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditTenderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditTenderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditTenderRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditTenderRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditTenderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditTenderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditTenderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditTenderRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServiceType

`func (o *EditTenderRequest) GetServiceType() TenderServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *EditTenderRequest) GetServiceTypeOk() (*TenderServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *EditTenderRequest) SetServiceType(v TenderServiceType)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *EditTenderRequest) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


