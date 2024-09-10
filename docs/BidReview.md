# BidReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Уникальный идентификатор отзыва, присвоенный сервером. | 
**Description** | **string** | Описание предложения | 
**CreatedAt** | **string** | Серверная дата и время в момент, когда пользователь отправил отзыв на предложение. Передается в формате RFC3339.  | 

## Methods

### NewBidReview

`func NewBidReview(id string, description string, createdAt string, ) *BidReview`

NewBidReview instantiates a new BidReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBidReviewWithDefaults

`func NewBidReviewWithDefaults() *BidReview`

NewBidReviewWithDefaults instantiates a new BidReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BidReview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BidReview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BidReview) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *BidReview) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BidReview) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BidReview) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCreatedAt

`func (o *BidReview) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BidReview) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BidReview) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


