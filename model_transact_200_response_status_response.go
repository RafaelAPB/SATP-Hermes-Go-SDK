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

// checks if the Transact200ResponseStatusResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transact200ResponseStatusResponse{}

// Transact200ResponseStatusResponse struct for Transact200ResponseStatusResponse
type Transact200ResponseStatusResponse struct {
	// Current status of the SATP session.
	Status string `json:"status"`
	// More detailed status of the SATP session.
	Substatus string `json:"substatus"`
	// Current stage of the SATP session.
	Stage string `json:"stage"`
	// Current step within the stage of the SATP session.
	Step string `json:"step"`
	// The start time of the SATP session.
	StartTime time.Time `json:"startTime"`
}

// NewTransact200ResponseStatusResponse instantiates a new Transact200ResponseStatusResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransact200ResponseStatusResponse(status string, substatus string, stage string, step string, startTime time.Time) *Transact200ResponseStatusResponse {
	this := Transact200ResponseStatusResponse{}
	this.Status = status
	this.Substatus = substatus
	this.Stage = stage
	this.Step = step
	this.StartTime = startTime
	return &this
}

// NewTransact200ResponseStatusResponseWithDefaults instantiates a new Transact200ResponseStatusResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransact200ResponseStatusResponseWithDefaults() *Transact200ResponseStatusResponse {
	this := Transact200ResponseStatusResponse{}
	return &this
}

// GetStatus returns the Status field value
func (o *Transact200ResponseStatusResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transact200ResponseStatusResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transact200ResponseStatusResponse) SetStatus(v string) {
	o.Status = v
}

// GetSubstatus returns the Substatus field value
func (o *Transact200ResponseStatusResponse) GetSubstatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Substatus
}

// GetSubstatusOk returns a tuple with the Substatus field value
// and a boolean to check if the value has been set.
func (o *Transact200ResponseStatusResponse) GetSubstatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Substatus, true
}

// SetSubstatus sets field value
func (o *Transact200ResponseStatusResponse) SetSubstatus(v string) {
	o.Substatus = v
}

// GetStage returns the Stage field value
func (o *Transact200ResponseStatusResponse) GetStage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value
// and a boolean to check if the value has been set.
func (o *Transact200ResponseStatusResponse) GetStageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Stage, true
}

// SetStage sets field value
func (o *Transact200ResponseStatusResponse) SetStage(v string) {
	o.Stage = v
}

// GetStep returns the Step field value
func (o *Transact200ResponseStatusResponse) GetStep() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Step
}

// GetStepOk returns a tuple with the Step field value
// and a boolean to check if the value has been set.
func (o *Transact200ResponseStatusResponse) GetStepOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Step, true
}

// SetStep sets field value
func (o *Transact200ResponseStatusResponse) SetStep(v string) {
	o.Step = v
}

// GetStartTime returns the StartTime field value
func (o *Transact200ResponseStatusResponse) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *Transact200ResponseStatusResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *Transact200ResponseStatusResponse) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o Transact200ResponseStatusResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transact200ResponseStatusResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["substatus"] = o.Substatus
	toSerialize["stage"] = o.Stage
	toSerialize["step"] = o.Step
	toSerialize["startTime"] = o.StartTime
	return toSerialize, nil
}

type NullableTransact200ResponseStatusResponse struct {
	value *Transact200ResponseStatusResponse
	isSet bool
}

func (v NullableTransact200ResponseStatusResponse) Get() *Transact200ResponseStatusResponse {
	return v.value
}

func (v *NullableTransact200ResponseStatusResponse) Set(val *Transact200ResponseStatusResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransact200ResponseStatusResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransact200ResponseStatusResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransact200ResponseStatusResponse(val *Transact200ResponseStatusResponse) *NullableTransact200ResponseStatusResponse {
	return &NullableTransact200ResponseStatusResponse{value: val, isSet: true}
}

func (v NullableTransact200ResponseStatusResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransact200ResponseStatusResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


