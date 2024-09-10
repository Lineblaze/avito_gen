# CreateTenderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Полное название тендера | 
**Description** | **string** | Описание тендера | 
**ServiceType** | [**TenderServiceType**](TenderServiceType.md) |  | 
**Status** | [**TenderStatus**](TenderStatus.md) |  | 
**OrganizationId** | **string** | Уникальный идентификатор организации, присвоенный сервером. | 
**CreatorUsername** | **string** | Уникальный slug пользователя. | 

## Methods

### NewCreateTenderRequest

`func NewCreateTenderRequest(name string, description string, serviceType TenderServiceType, status TenderStatus, organizationId string, creatorUsername string, ) *CreateTenderRequest`

NewCreateTenderRequest instantiates a new CreateTenderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenderRequestWithDefaults

`func NewCreateTenderRequestWithDefaults() *CreateTenderRequest`

NewCreateTenderRequestWithDefaults instantiates a new CreateTenderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTenderRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenderRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenderRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTenderRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenderRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenderRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceType

`func (o *CreateTenderRequest) GetServiceType() TenderServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateTenderRequest) GetServiceTypeOk() (*TenderServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateTenderRequest) SetServiceType(v TenderServiceType)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *CreateTenderRequest) GetStatus() TenderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTenderRequest) GetStatusOk() (*TenderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTenderRequest) SetStatus(v TenderStatus)`

SetStatus sets Status field to given value.


### GetOrganizationId

`func (o *CreateTenderRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *CreateTenderRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *CreateTenderRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCreatorUsername

`func (o *CreateTenderRequest) GetCreatorUsername() string`

GetCreatorUsername returns the CreatorUsername field if non-nil, zero value otherwise.

### GetCreatorUsernameOk

`func (o *CreateTenderRequest) GetCreatorUsernameOk() (*string, bool)`

GetCreatorUsernameOk returns a tuple with the CreatorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorUsername

`func (o *CreateTenderRequest) SetCreatorUsername(v string)`

SetCreatorUsername sets CreatorUsername field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


