# Tender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор тендера, присвоенный сервером. | 
**Name** | **string** | Полное название тендера | 
**Description** | **string** | Описание тендера | 
**ServiceType** | [**TenderServiceType**](TenderServiceType.md) |  | 
**Status** | [**TenderStatus**](TenderStatus.md) |  | 
**OrganizationId** | **string** | Уникальный идентификатор организации, присвоенный сервером. | 
**Version** | **int32** | Номер версии посел правок | [default to 1]
**CreatedAt** | **string** | Серверная дата и время в момент, когда пользователь отправил тендер на создание. Передается в формате RFC3339.  | 

## Methods

### NewTender

`func NewTender(id string, name string, description string, serviceType TenderServiceType, status TenderStatus, organizationId string, version int32, createdAt string, ) *Tender`

NewTender instantiates a new Tender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenderWithDefaults

`func NewTenderWithDefaults() *Tender`

NewTenderWithDefaults instantiates a new Tender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Tender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tender) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Tender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tender) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Tender) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Tender) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Tender) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetServiceType

`func (o *Tender) GetServiceType() TenderServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *Tender) GetServiceTypeOk() (*TenderServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *Tender) SetServiceType(v TenderServiceType)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *Tender) GetStatus() TenderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Tender) GetStatusOk() (*TenderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Tender) SetStatus(v TenderStatus)`

SetStatus sets Status field to given value.


### GetOrganizationId

`func (o *Tender) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Tender) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Tender) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetVersion

`func (o *Tender) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Tender) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Tender) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *Tender) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Tender) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Tender) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


