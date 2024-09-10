# CreateBidRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Полное название предложения | 
**Description** | **string** | Описание предложения | 
**Status** | [**BidStatus**](BidStatus.md) |  | 
**TenderId** | **string** | Уникальный идентификатор тендера, присвоенный сервером. | 
**OrganizationId** | **string** | Уникальный идентификатор организации, присвоенный сервером. | 
**CreatorUsername** | **string** | Уникальный slug пользователя. | 

## Methods

### NewCreateBidRequest

`func NewCreateBidRequest(name string, description string, status BidStatus, tenderId string, organizationId string, creatorUsername string, ) *CreateBidRequest`

NewCreateBidRequest instantiates a new CreateBidRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBidRequestWithDefaults

`func NewCreateBidRequestWithDefaults() *CreateBidRequest`

NewCreateBidRequestWithDefaults instantiates a new CreateBidRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBidRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBidRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBidRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateBidRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateBidRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateBidRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *CreateBidRequest) GetStatus() BidStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBidRequest) GetStatusOk() (*BidStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBidRequest) SetStatus(v BidStatus)`

SetStatus sets Status field to given value.


### GetTenderId

`func (o *CreateBidRequest) GetTenderId() string`

GetTenderId returns the TenderId field if non-nil, zero value otherwise.

### GetTenderIdOk

`func (o *CreateBidRequest) GetTenderIdOk() (*string, bool)`

GetTenderIdOk returns a tuple with the TenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderId

`func (o *CreateBidRequest) SetTenderId(v string)`

SetTenderId sets TenderId field to given value.


### GetOrganizationId

`func (o *CreateBidRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateBidRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateBidRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCreatorUsername

`func (o *CreateBidRequest) GetCreatorUsername() string`

GetCreatorUsername returns the CreatorUsername field if non-nil, zero value otherwise.

### GetCreatorUsernameOk

`func (o *CreateBidRequest) GetCreatorUsernameOk() (*string, bool)`

GetCreatorUsernameOk returns a tuple with the CreatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUsername

`func (o *CreateBidRequest) SetCreatorUsername(v string)`

SetCreatorUsername sets CreatorUsername field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


