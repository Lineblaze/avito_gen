# EditBidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Полное название предложения | [optional] 
**Description** | Pointer to **string** | Описание предложения | [optional] 

## Methods

### NewEditBidRequest

`func NewEditBidRequest() *EditBidRequest`

NewEditBidRequest instantiates a new EditBidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditBidRequestWithDefaults

`func NewEditBidRequestWithDefaults() *EditBidRequest`

NewEditBidRequestWithDefaults instantiates a new EditBidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EditBidRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EditBidRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EditBidRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EditBidRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *EditBidRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditBidRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditBidRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditBidRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


