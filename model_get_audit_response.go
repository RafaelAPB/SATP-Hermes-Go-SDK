/*
SATP Gateway Client (Business Logic Orchestrator)

SATP is a protocol operating between two gateways that conducts the transfer of a digital asset from one gateway to another. The protocol establishes a secure channel between the endpoints and implements a 2-phase commit to ensure the properties of transfer atomicity, consistency, isolation and durability.  This API defines the gateway client facing API (business logic orchestrator, or BLO), which is named API-Type 1 in the SATP-Core specification.  **Additional Resources**: - [Proposed SATP Charter](https://datatracker.ietf.org/doc/charter-ietf-satp/) - [SATP Core draft](https://datatracker.ietf.org/doc/draft-ietf-satp-core) - [SATP Crash Recovery draft](https://datatracker.ietf.org/doc/draft-belchior-satp-gateway-recovery/) - [SATP Architecture draft](https://datatracker.ietf.org/doc/draft-ietf-satp-architecture/) - [SATP Use-Cases draft](https://datatracker.ietf.org/doc/draft-ramakrishna-sat-use-cases/) - [SATP Data sharing draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-data-sharing) - [SATP View Addresses draft](https://datatracker.ietf.org/doc/draft-ramakrishna-satp-views-addresses)

API version: v0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package generated

import (
	"encoding/json"
	"time"
)

// checks if the GetAuditResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAuditResponse{}

// GetAuditResponse struct for GetAuditResponse
type GetAuditResponse struct {
	// An array of strings representing proofs.
	Proofs []string `json:"proofs,omitempty"`
	// The start datetime of the audit period.
	AuditStartTime *time.Time `json:"auditStartTime,omitempty"`
	// The end datetime of the audit period.
	AuditEndTime *time.Time `json:"auditEndTime,omitempty"`
}

// NewGetAuditResponse instantiates a new GetAuditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAuditResponse() *GetAuditResponse {
	this := GetAuditResponse{}
	return &this
}

// NewGetAuditResponseWithDefaults instantiates a new GetAuditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAuditResponseWithDefaults() *GetAuditResponse {
	this := GetAuditResponse{}
	return &this
}

// GetProofs returns the Proofs field value if set, zero value otherwise.
func (o *GetAuditResponse) GetProofs() []string {
	if o == nil || IsNil(o.Proofs) {
		var ret []string
		return ret
	}
	return o.Proofs
}

// GetProofsOk returns a tuple with the Proofs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuditResponse) GetProofsOk() ([]string, bool) {
	if o == nil || IsNil(o.Proofs) {
		return nil, false
	}
	return o.Proofs, true
}

// HasProofs returns a boolean if a field has been set.
func (o *GetAuditResponse) HasProofs() bool {
	if o != nil && !IsNil(o.Proofs) {
		return true
	}

	return false
}

// SetProofs gets a reference to the given []string and assigns it to the Proofs field.
func (o *GetAuditResponse) SetProofs(v []string) {
	o.Proofs = v
}

// GetAuditStartTime returns the AuditStartTime field value if set, zero value otherwise.
func (o *GetAuditResponse) GetAuditStartTime() time.Time {
	if o == nil || IsNil(o.AuditStartTime) {
		var ret time.Time
		return ret
	}
	return *o.AuditStartTime
}

// GetAuditStartTimeOk returns a tuple with the AuditStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuditResponse) GetAuditStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuditStartTime) {
		return nil, false
	}
	return o.AuditStartTime, true
}

// HasAuditStartTime returns a boolean if a field has been set.
func (o *GetAuditResponse) HasAuditStartTime() bool {
	if o != nil && !IsNil(o.AuditStartTime) {
		return true
	}

	return false
}

// SetAuditStartTime gets a reference to the given time.Time and assigns it to the AuditStartTime field.
func (o *GetAuditResponse) SetAuditStartTime(v time.Time) {
	o.AuditStartTime = &v
}

// GetAuditEndTime returns the AuditEndTime field value if set, zero value otherwise.
func (o *GetAuditResponse) GetAuditEndTime() time.Time {
	if o == nil || IsNil(o.AuditEndTime) {
		var ret time.Time
		return ret
	}
	return *o.AuditEndTime
}

// GetAuditEndTimeOk returns a tuple with the AuditEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAuditResponse) GetAuditEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AuditEndTime) {
		return nil, false
	}
	return o.AuditEndTime, true
}

// HasAuditEndTime returns a boolean if a field has been set.
func (o *GetAuditResponse) HasAuditEndTime() bool {
	if o != nil && !IsNil(o.AuditEndTime) {
		return true
	}

	return false
}

// SetAuditEndTime gets a reference to the given time.Time and assigns it to the AuditEndTime field.
func (o *GetAuditResponse) SetAuditEndTime(v time.Time) {
	o.AuditEndTime = &v
}

func (o GetAuditResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAuditResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Proofs) {
		toSerialize["proofs"] = o.Proofs
	}
	if !IsNil(o.AuditStartTime) {
		toSerialize["auditStartTime"] = o.AuditStartTime
	}
	if !IsNil(o.AuditEndTime) {
		toSerialize["auditEndTime"] = o.AuditEndTime
	}
	return toSerialize, nil
}

type NullableGetAuditResponse struct {
	value *GetAuditResponse
	isSet bool
}

func (v NullableGetAuditResponse) Get() *GetAuditResponse {
	return v.value
}

func (v *NullableGetAuditResponse) Set(val *GetAuditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAuditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAuditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAuditResponse(val *GetAuditResponse) *NullableGetAuditResponse {
	return &NullableGetAuditResponse{value: val, isSet: true}
}

func (v NullableGetAuditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAuditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


