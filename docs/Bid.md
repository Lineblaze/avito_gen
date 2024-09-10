# Bid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор предложения, присвоенный сервером. | 
**Name** | **string** | Полное название предложения | 
**Description** | **string** | Описание предложения | 
**Status** | [**BidStatus**](BidStatus.md) |  | 
**TenderId** | **string** | Уникальный идентификатор тендера, присвоенный сервером. | 
**AuthorType** | [**BidAuthorType**](BidAuthorType.md) |  | 
**AuthorId** | **string** | Уникальный идентификатор автора предложения, присвоенный сервером. | 
**Version** | **int32** | Номер версии посел правок | [default to 1]
**CreatedAt** | **string** | Серверная дата и время в момент, когда пользователь отправил предложение на создание. Передается в формате RFC3339.  | 

## Methods

### NewBid

`func NewBid(id string, name string, description string, status BidStatus, tenderId string, authorType BidAuthorType, authorId string, version int32, createdAt string, ) *Bid`

NewBid instantiates a new Bid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidWithDefaults

`func NewBidWithDefaults() *Bid`

NewBidWithDefaults instantiates a new Bid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Bid) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Bid) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Bid) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Bid) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Bid) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Bid) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Bid) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Bid) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Bid) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetStatus

`func (o *Bid) GetStatus() BidStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Bid) GetStatusOk() (*BidStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Bid) SetStatus(v BidStatus)`

SetStatus sets Status field to given value.


### GetTenderId

`func (o *Bid) GetTenderId() string`

GetTenderId returns the TenderId field if non-nil, zero value otherwise.

### GetTenderIdOk

`func (o *Bid) GetTenderIdOk() (*string, bool)`

GetTenderIdOk returns a tuple with the TenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenderId

`func (o *Bid) SetTenderId(v string)`

SetTenderId sets TenderId field to given value.


### GetAuthorType

`func (o *Bid) GetAuthorType() BidAuthorType`

GetAuthorType returns the AuthorType field if non-nil, zero value otherwise.

### GetAuthorTypeOk

`func (o *Bid) GetAuthorTypeOk() (*BidAuthorType, bool)`

GetAuthorTypeOk returns a tuple with the AuthorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorType

`func (o *Bid) SetAuthorType(v BidAuthorType)`

SetAuthorType sets AuthorType field to given value.


### GetAuthorId

`func (o *Bid) GetAuthorId() string`

GetAuthorId returns the AuthorId field if non-nil, zero value otherwise.

### GetAuthorIdOk

`func (o *Bid) GetAuthorIdOk() (*string, bool)`

GetAuthorIdOk returns a tuple with the AuthorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorId

`func (o *Bid) SetAuthorId(v string)`

SetAuthorId sets AuthorId field to given value.


### GetVersion

`func (o *Bid) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Bid) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Bid) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *Bid) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Bid) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Bid) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


