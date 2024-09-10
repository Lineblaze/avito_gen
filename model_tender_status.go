/*
Tender Management API

API для управления тендерами и предложениями.   Основные функции API включают управление тендерами (создание, изменение, получение списка) и управление предложениями (создание, изменение, получение списка).

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// TenderStatus Статус тендер
type TenderStatus string

// List of tenderStatus
const (
	CREATED   TenderStatus = "Created"
	PUBLISHED TenderStatus = "Published"
	CLOSED    TenderStatus = "Closed"
)

// All allowed values of TenderStatus enum
var AllowedTenderStatusEnumValues = []TenderStatus{
	"Created",
	"Published",
	"Closed",
}

func (v *TenderStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TenderStatus(value)
	for _, existing := range AllowedTenderStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TenderStatus", value)
}

// NewTenderStatusFromValue returns a pointer to a valid TenderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTenderStatusFromValue(v string) (*TenderStatus, error) {
	ev := TenderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TenderStatus: valid values are %v", v, AllowedTenderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TenderStatus) IsValid() bool {
	for _, existing := range AllowedTenderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to tenderStatus value
func (v TenderStatus) Ptr() *TenderStatus {
	return &v
}

type NullableTenderStatus struct {
	value *TenderStatus
	isSet bool
}

func (v NullableTenderStatus) Get() *TenderStatus {
	return v.value
}

func (v *NullableTenderStatus) Set(val *TenderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableTenderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableTenderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenderStatus(val *TenderStatus) *NullableTenderStatus {
	return &NullableTenderStatus{value: val, isSet: true}
}

func (v NullableTenderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
