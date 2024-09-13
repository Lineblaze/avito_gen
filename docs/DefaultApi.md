# \DefaultApi

All URIs are relative to *http://localhost:8080/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignEmployeeToOrganization**](DefaultApi.md#AssignEmployeeToOrganization) | **Post** /assign | Присвоение пользователя к организации
[**CheckServer**](DefaultApi.md#CheckServer) | **Get** /ping | Проверка доступности сервера
[**CreateBid**](DefaultApi.md#CreateBid) | **Post** /bids/new | Создание нового предложения
[**CreateEmployee**](DefaultApi.md#CreateEmployee) | **Post** /employeee/new | Создание нового пользователя
[**CreateOrganization**](DefaultApi.md#CreateOrganization) | **Post** /organization/new | Создание новой организации
[**CreateTender**](DefaultApi.md#CreateTender) | **Post** /tenders/new | Создание нового тендера
[**EditBid**](DefaultApi.md#EditBid) | **Patch** /bids/{bidId}/edit | Редактирование параметров предложения
[**EditTender**](DefaultApi.md#EditTender) | **Patch** /tenders/{tenderId}/edit | Редактирование тендера
[**GetBidReviews**](DefaultApi.md#GetBidReviews) | **Get** /bids/{tenderId}/reviews | Просмотр отзывов на прошлые предложения
[**GetBidStatus**](DefaultApi.md#GetBidStatus) | **Get** /bids/{bidId}/status | Получение текущего статуса предложения
[**GetBidsForTender**](DefaultApi.md#GetBidsForTender) | **Get** /bids/{tenderId}/list | Получение списка предложений для тендера
[**GetTenderStatus**](DefaultApi.md#GetTenderStatus) | **Get** /tenders/{tenderId}/status | Получение текущего статуса тендера
[**GetTenders**](DefaultApi.md#GetTenders) | **Get** /tenders | Получение списка тендеров
[**GetUserBids**](DefaultApi.md#GetUserBids) | **Get** /bids/my | Получение списка ваших предложений
[**GetUserTenders**](DefaultApi.md#GetUserTenders) | **Get** /tenders/my | Получить тендеры пользователя
[**RollbackBid**](DefaultApi.md#RollbackBid) | **Put** /bids/{bidId}/rollback/{version} | Откат версии предложения
[**RollbackTender**](DefaultApi.md#RollbackTender) | **Put** /tenders/{tenderId}/rollback/{version} | Откат версии тендера
[**SubmitBidDecision**](DefaultApi.md#SubmitBidDecision) | **Put** /bids/{bidId}/submit_decision | Отправка решения по предложению
[**SubmitBidFeedback**](DefaultApi.md#SubmitBidFeedback) | **Put** /bids/{bidId}/feedback | Отправка отзыва по предложению
[**UpdateBidStatus**](DefaultApi.md#UpdateBidStatus) | **Put** /bids/{bidId}/status/{bidId} | Изменение статуса предложения
[**UpdateTenderStatus**](DefaultApi.md#UpdateTenderStatus) | **Put** /tenders/{tenderId}/status | Изменение статуса тендера



## AssignEmployeeToOrganization

> OrganizationResponsible AssignEmployeeToOrganization(ctx).AssignEmployeeToOrganizationRequest(assignEmployeeToOrganizationRequest).Execute()

Присвоение пользователя к организации



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assignEmployeeToOrganizationRequest := *openapiclient.NewAssignEmployeeToOrganizationRequest("874fdc00-8bb4-4423-894e-01a6a3937883", "ea35e012-cc15-4d94-a0ec-d83a96eb5540") // AssignEmployeeToOrganizationRequest | Данные для присвоения.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.AssignEmployeeToOrganization(context.Background()).AssignEmployeeToOrganizationRequest(assignEmployeeToOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.AssignEmployeeToOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignEmployeeToOrganization`: OrganizationResponsible
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.AssignEmployeeToOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignEmployeeToOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignEmployeeToOrganizationRequest** | [**AssignEmployeeToOrganizationRequest**](AssignEmployeeToOrganizationRequest.md) | Данные для присвоения. | 

### Return type

[**OrganizationResponsible**](OrganizationResponsible.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckServer

> string CheckServer(ctx).Execute()

Проверка доступности сервера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CheckServer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CheckServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckServer`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CheckServer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckServerRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBid

> Bid CreateBid(ctx).CreateBidRequest(createBidRequest).Execute()

Создание нового предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createBidRequest := *openapiclient.NewCreateBidRequest("Name_example", "Description_example", openapiclient.bidStatus("Open"), "550e8400-e29b-41d4-a716-446655440000", "550e8400-e29b-41d4-a716-446655440000", "test_user") // CreateBidRequest | Данные нового предложения.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateBid(context.Background()).CreateBidRequest(createBidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBid`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateBid`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBidRequest** | [**CreateBidRequest**](CreateBidRequest.md) | Данные нового предложения. | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateEmployee

> Employee CreateEmployee(ctx).CreateEmployeeRequest(createEmployeeRequest).Execute()

Создание нового пользователя



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createEmployeeRequest := *openapiclient.NewCreateEmployeeRequest("test_user") // CreateEmployeeRequest | Данные нового пользователя.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateEmployee(context.Background()).CreateEmployeeRequest(createEmployeeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateEmployee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmployee`: Employee
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateEmployee`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmployeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmployeeRequest** | [**CreateEmployeeRequest**](CreateEmployeeRequest.md) | Данные нового пользователя. | 

### Return type

[**Employee**](Employee.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrganization

> Organization CreateOrganization(ctx).CreateOrganizationRequest(createOrganizationRequest).Execute()

Создание новой организации



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createOrganizationRequest := *openapiclient.NewCreateOrganizationRequest("test_organization", "test_description", openapiclient.organizationType("IE")) // CreateOrganizationRequest | Данные нового  новой организации.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateOrganization(context.Background()).CreateOrganizationRequest(createOrganizationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateOrganization``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrganization`: Organization
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateOrganization`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrganizationRequest** | [**CreateOrganizationRequest**](CreateOrganizationRequest.md) | Данные нового  новой организации. | 

### Return type

[**Organization**](Organization.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTender

> Tender CreateTender(ctx).CreateTenderRequest(createTenderRequest).Execute()

Создание нового тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createTenderRequest := *openapiclient.NewCreateTenderRequest("Name_example", "Description_example", openapiclient.tenderServiceType("Construction"), openapiclient.tenderStatus("Open"), "550e8400-e29b-41d4-a716-446655440000", "test_user") // CreateTenderRequest | Данные нового тендера.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.CreateTender(context.Background()).CreateTenderRequest(createTenderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateTender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTender`: Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateTender`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTenderRequest** | [**CreateTenderRequest**](CreateTenderRequest.md) | Данные нового тендера. | 

### Return type

[**Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditBid

> Bid EditBid(ctx, bidId).Username(username).EditBidRequest(editBidRequest).Execute()

Редактирование параметров предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    username := "username_example" // string | 
    editBidRequest := *openapiclient.NewEditBidRequest() // EditBidRequest | Перечисление параметров и их новых значений для обновления предложения.  Если значение не передано, оно останется без изменений. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EditBid(context.Background(), bidId).Username(username).EditBidRequest(editBidRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditBid`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EditBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** |  | 
 **editBidRequest** | [**EditBidRequest**](EditBidRequest.md) | Перечисление параметров и их новых значений для обновления предложения.  Если значение не передано, оно останется без изменений.  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EditTender

> Tender EditTender(ctx, tenderId).Username(username).EditTenderRequest(editTenderRequest).Execute()

Редактирование тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    username := "username_example" // string | 
    editTenderRequest := *openapiclient.NewEditTenderRequest() // EditTenderRequest | Перечисление параметров и их новых значений для обновления тендера.  Если значение не передано, оно останется без изменений. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.EditTender(context.Background(), tenderId).Username(username).EditTenderRequest(editTenderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.EditTender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EditTender`: Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.EditTender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditTenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** |  | 
 **editTenderRequest** | [**EditTenderRequest**](EditTenderRequest.md) | Перечисление параметров и их новых значений для обновления тендера.  Если значение не передано, оно останется без изменений.  | 

### Return type

[**Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidReviews

> []BidReview GetBidReviews(ctx, tenderId).AuthorUsername(authorUsername).RequesterUsername(requesterUsername).Limit(limit).Offset(offset).Execute()

Просмотр отзывов на прошлые предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    authorUsername := "authorUsername_example" // string | Имя пользователя автора предложений, отзывы на которые нужно просмотреть.
    requesterUsername := "requesterUsername_example" // string | Имя пользователя, который запрашивает отзывы.
    limit := int32(56) // int32 | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  (optional) (default to 5)
    offset := int32(56) // int32 | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBidReviews(context.Background(), tenderId).AuthorUsername(authorUsername).RequesterUsername(requesterUsername).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBidReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBidReviews`: []BidReview
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBidReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorUsername** | **string** | Имя пользователя автора предложений, отзывы на которые нужно просмотреть. | 
 **requesterUsername** | **string** | Имя пользователя, который запрашивает отзывы. | 
 **limit** | **int32** | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  | [default to 5]
 **offset** | **int32** | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  | [default to 0]

### Return type

[**[]BidReview**](BidReview.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidStatus

> BidStatus GetBidStatus(ctx, bidId).Username(username).Execute()

Получение текущего статуса предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBidStatus(context.Background(), bidId).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBidStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBidStatus`: BidStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBidStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** |  | 

### Return type

[**BidStatus**](BidStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBidsForTender

> []Bid GetBidsForTender(ctx, tenderId).Username(username).Limit(limit).Offset(offset).Execute()

Получение списка предложений для тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    username := "username_example" // string | 
    limit := int32(56) // int32 | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  (optional) (default to 5)
    offset := int32(56) // int32 | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetBidsForTender(context.Background(), tenderId).Username(username).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetBidsForTender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBidsForTender`: []Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetBidsForTender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBidsForTenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** |  | 
 **limit** | **int32** | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  | [default to 5]
 **offset** | **int32** | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  | [default to 0]

### Return type

[**[]Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenderStatus

> TenderStatus GetTenderStatus(ctx, tenderId).Username(username).Execute()

Получение текущего статуса тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTenderStatus(context.Background(), tenderId).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTenderStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenderStatus`: TenderStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTenderStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** |  | 

### Return type

[**TenderStatus**](TenderStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenders

> []Tender GetTenders(ctx).Limit(limit).Offset(offset).ServiceType(serviceType).Execute()

Получение списка тендеров



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  (optional) (default to 5)
    offset := int32(56) // int32 | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  (optional) (default to 0)
    serviceType := []openapiclient.TenderServiceType{openapiclient.tenderServiceType("Construction")} // []TenderServiceType | Возвращенные тендеры должны соответствовать указанным видам услуг.  Если список пустой, фильтры не применяются.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetTenders(context.Background()).Limit(limit).Offset(offset).ServiceType(serviceType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetTenders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenders`: []Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetTenders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTendersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  | [default to 5]
 **offset** | **int32** | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  | [default to 0]
 **serviceType** | [**[]TenderServiceType**](TenderServiceType.md) | Возвращенные тендеры должны соответствовать указанным видам услуг.  Если список пустой, фильтры не применяются.  | 

### Return type

[**[]Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserBids

> []Bid GetUserBids(ctx).Limit(limit).Offset(offset).Username(username).Execute()

Получение списка ваших предложений



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  (optional) (default to 5)
    offset := int32(56) // int32 | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  (optional) (default to 0)
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserBids(context.Background()).Limit(limit).Offset(offset).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserBids``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserBids`: []Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserBids`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserBidsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  | [default to 5]
 **offset** | **int32** | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  | [default to 0]
 **username** | **string** |  | 

### Return type

[**[]Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserTenders

> []Tender GetUserTenders(ctx).Limit(limit).Offset(offset).Username(username).Execute()

Получить тендеры пользователя



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  (optional) (default to 5)
    offset := int32(56) // int32 | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  (optional) (default to 0)
    username := "username_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetUserTenders(context.Background()).Limit(limit).Offset(offset).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetUserTenders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserTenders`: []Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetUserTenders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTendersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Максимальное число возвращаемых объектов. Используется для запросов с пагинацией.  Сервер должен возвращать максимальное допустимое число объектов.  | [default to 5]
 **offset** | **int32** | Какое количество объектов должно быть пропущено с начала. Используется для запросов с пагинацией.  | [default to 0]
 **username** | **string** |  | 

### Return type

[**[]Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackBid

> Bid RollbackBid(ctx, bidId, version).Username(username).Execute()

Откат версии предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    version := int32(56) // int32 | Номер версии, к которой нужно откатить предложение.
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RollbackBid(context.Background(), bidId, version).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RollbackBid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackBid`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RollbackBid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 
**version** | **int32** | Номер версии, к которой нужно откатить предложение. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackBidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **username** | **string** |  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackTender

> Tender RollbackTender(ctx, tenderId, version).Username(username).Execute()

Откат версии тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    version := int32(56) // int32 | Номер версии, к которой нужно откатить тендер.
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.RollbackTender(context.Background(), tenderId, version).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RollbackTender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackTender`: Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RollbackTender`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 
**version** | **int32** | Номер версии, к которой нужно откатить тендер. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackTenderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **username** | **string** |  | 

### Return type

[**Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBidDecision

> Bid SubmitBidDecision(ctx, bidId).Decision(decision).Username(username).Execute()

Отправка решения по предложению



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    decision := openapiclient.bidDecision("Approved") // BidDecision | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubmitBidDecision(context.Background(), bidId).Decision(decision).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitBidDecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitBidDecision`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubmitBidDecision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBidDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **decision** | [**BidDecision**](BidDecision.md) |  | 
 **username** | **string** |  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBidFeedback

> Bid SubmitBidFeedback(ctx, bidId).BidFeedback(bidFeedback).Username(username).Execute()

Отправка отзыва по предложению



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    bidFeedback := "bidFeedback_example" // string | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.SubmitBidFeedback(context.Background(), bidId).BidFeedback(bidFeedback).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.SubmitBidFeedback``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SubmitBidFeedback`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.SubmitBidFeedback`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBidFeedbackRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bidFeedback** | **string** |  | 
 **username** | **string** |  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBidStatus

> Bid UpdateBidStatus(ctx, bidId).Status(status).Username(username).Execute()

Изменение статуса предложения



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bidId := "bidId_example" // string | 
    status := openapiclient.bidStatus("Open") // BidStatus | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateBidStatus(context.Background(), bidId).Status(status).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateBidStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBidStatus`: Bid
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateBidStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bidId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBidStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**BidStatus**](BidStatus.md) |  | 
 **username** | **string** |  | 

### Return type

[**Bid**](Bid.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTenderStatus

> Tender UpdateTenderStatus(ctx, tenderId).Status(status).Username(username).Execute()

Изменение статуса тендера



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenderId := "tenderId_example" // string | 
    status := openapiclient.tenderStatus("Open") // TenderStatus | 
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.UpdateTenderStatus(context.Background(), tenderId).Status(status).Username(username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.UpdateTenderStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTenderStatus`: Tender
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.UpdateTenderStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenderStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | [**TenderStatus**](TenderStatus.md) |  | 
 **username** | **string** |  | 

### Return type

[**Tender**](Tender.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

