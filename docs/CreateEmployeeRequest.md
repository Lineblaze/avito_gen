# CreateEmployeeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Уникальный slug пользователя. | 
**Firstname** | Pointer to **string** | Имя пользователя. | [optional] 
**Lastname** | Pointer to **string** | Фамилия пользователя. | [optional] 

## Methods

### NewCreateEmployeeRequest

`func NewCreateEmployeeRequest(username string, ) *CreateEmployeeRequest`

NewCreateEmployeeRequest instantiates a new CreateEmployeeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEmployeeRequestWithDefaults

`func NewCreateEmployeeRequestWithDefaults() *CreateEmployeeRequest`

NewCreateEmployeeRequestWithDefaults instantiates a new CreateEmployeeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CreateEmployeeRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateEmployeeRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateEmployeeRequest) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetFirstname

`func (o *CreateEmployeeRequest) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *CreateEmployeeRequest) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *CreateEmployeeRequest) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *CreateEmployeeRequest) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *CreateEmployeeRequest) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *CreateEmployeeRequest) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *CreateEmployeeRequest) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *CreateEmployeeRequest) HasLastname() bool`

HasLastname returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


