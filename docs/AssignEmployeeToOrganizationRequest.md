# AssignEmployeeToOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationId** | **string** | Уникальный slug организации. | 
**UserId** | **string** | Уникальный slug пользователя. | 

## Methods

### NewAssignEmployeeToOrganizationRequest

`func NewAssignEmployeeToOrganizationRequest(organizationId string, userId string, ) *AssignEmployeeToOrganizationRequest`

NewAssignEmployeeToOrganizationRequest instantiates a new AssignEmployeeToOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignEmployeeToOrganizationRequestWithDefaults

`func NewAssignEmployeeToOrganizationRequestWithDefaults() *AssignEmployeeToOrganizationRequest`

NewAssignEmployeeToOrganizationRequestWithDefaults instantiates a new AssignEmployeeToOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationId

`func (o *AssignEmployeeToOrganizationRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AssignEmployeeToOrganizationRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AssignEmployeeToOrganizationRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetUserId

`func (o *AssignEmployeeToOrganizationRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AssignEmployeeToOrganizationRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AssignEmployeeToOrganizationRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


