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

// checks if the GetChains200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChains200ResponseInner{}

// GetChains200ResponseInner struct for GetChains200ResponseInner
type GetChains200ResponseInner struct {
	// A unique identifier for the blockchain network.
	ChainId string `json:"chainId"`
	// The name of the blockchain network.
	ChainName string `json:"chainName"`
	// The type of blockchain network (e.g., 'evm', 'fabric').
	ChainType string `json:"chainType"`
	// The specific network name within the blockchain (e.g., 'mainnet', 'testnet').
	NetworkName string `json:"networkName"`
}

// NewGetChains200ResponseInner instantiates a new GetChains200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChains200ResponseInner(chainId string, chainName string, chainType string, networkName string) *GetChains200ResponseInner {
	this := GetChains200ResponseInner{}
	this.ChainId = chainId
	this.ChainName = chainName
	this.ChainType = chainType
	this.NetworkName = networkName
	return &this
}

// NewGetChains200ResponseInnerWithDefaults instantiates a new GetChains200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChains200ResponseInnerWithDefaults() *GetChains200ResponseInner {
	this := GetChains200ResponseInner{}
	return &this
}

// GetChainId returns the ChainId field value
func (o *GetChains200ResponseInner) GetChainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainId
}

// GetChainIdOk returns a tuple with the ChainId field value
// and a boolean to check if the value has been set.
func (o *GetChains200ResponseInner) GetChainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainId, true
}

// SetChainId sets field value
func (o *GetChains200ResponseInner) SetChainId(v string) {
	o.ChainId = v
}

// GetChainName returns the ChainName field value
func (o *GetChains200ResponseInner) GetChainName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainName
}

// GetChainNameOk returns a tuple with the ChainName field value
// and a boolean to check if the value has been set.
func (o *GetChains200ResponseInner) GetChainNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainName, true
}

// SetChainName sets field value
func (o *GetChains200ResponseInner) SetChainName(v string) {
	o.ChainName = v
}

// GetChainType returns the ChainType field value
func (o *GetChains200ResponseInner) GetChainType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChainType
}

// GetChainTypeOk returns a tuple with the ChainType field value
// and a boolean to check if the value has been set.
func (o *GetChains200ResponseInner) GetChainTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChainType, true
}

// SetChainType sets field value
func (o *GetChains200ResponseInner) SetChainType(v string) {
	o.ChainType = v
}

// GetNetworkName returns the NetworkName field value
func (o *GetChains200ResponseInner) GetNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkName
}

// GetNetworkNameOk returns a tuple with the NetworkName field value
// and a boolean to check if the value has been set.
func (o *GetChains200ResponseInner) GetNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkName, true
}

// SetNetworkName sets field value
func (o *GetChains200ResponseInner) SetNetworkName(v string) {
	o.NetworkName = v
}

func (o GetChains200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChains200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["chainId"] = o.ChainId
	toSerialize["chainName"] = o.ChainName
	toSerialize["chainType"] = o.ChainType
	toSerialize["networkName"] = o.NetworkName
	return toSerialize, nil
}

type NullableGetChains200ResponseInner struct {
	value *GetChains200ResponseInner
	isSet bool
}

func (v NullableGetChains200ResponseInner) Get() *GetChains200ResponseInner {
	return v.value
}

func (v *NullableGetChains200ResponseInner) Set(val *GetChains200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChains200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChains200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChains200ResponseInner(val *GetChains200ResponseInner) *NullableGetChains200ResponseInner {
	return &NullableGetChains200ResponseInner{value: val, isSet: true}
}

func (v NullableGetChains200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChains200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

