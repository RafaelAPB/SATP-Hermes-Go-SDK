/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
)

// checks if the Pause200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pause200Response{}

// Pause200Response Response schema for a pause request, returning the status of the SATP session.
type Pause200Response struct {
	StatusResponse Transact200ResponseStatusResponse `json:"statusResponse"`
}

// NewPause200Response instantiates a new Pause200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPause200Response(statusResponse Transact200ResponseStatusResponse) *Pause200Response {
	this := Pause200Response{}
	this.StatusResponse = statusResponse
	return &this
}

// NewPause200ResponseWithDefaults instantiates a new Pause200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPause200ResponseWithDefaults() *Pause200Response {
	this := Pause200Response{}
	return &this
}

// GetStatusResponse returns the StatusResponse field value
func (o *Pause200Response) GetStatusResponse() Transact200ResponseStatusResponse {
	if o == nil {
		var ret Transact200ResponseStatusResponse
		return ret
	}

	return o.StatusResponse
}

// GetStatusResponseOk returns a tuple with the StatusResponse field value
// and a boolean to check if the value has been set.
func (o *Pause200Response) GetStatusResponseOk() (*Transact200ResponseStatusResponse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusResponse, true
}

// SetStatusResponse sets field value
func (o *Pause200Response) SetStatusResponse(v Transact200ResponseStatusResponse) {
	o.StatusResponse = v
}

func (o Pause200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pause200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["statusResponse"] = o.StatusResponse
	return toSerialize, nil
}

type NullablePause200Response struct {
	value *Pause200Response
	isSet bool
}

func (v NullablePause200Response) Get() *Pause200Response {
	return v.value
}

func (v *NullablePause200Response) Set(val *Pause200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePause200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePause200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePause200Response(val *Pause200Response) *NullablePause200Response {
	return &NullablePause200Response{value: val, isSet: true}
}

func (v NullablePause200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePause200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

