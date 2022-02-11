/*
Tendermint RPC

Tendermint supports the following RPC protocols:  * URI over HTTP * JSONRPC over HTTP * JSONRPC over websockets  ## Configuration  RPC can be configured by tuning parameters under `[rpc]` table in the `$TMHOME/config/config.toml` file or by using the `--rpc.X` command-line flags.  Default rpc listen address is `tcp://0.0.0.0:26657`. To set another address, set the `laddr` config parameter to desired value. CORS (Cross-Origin Resource Sharing) can be enabled by setting `cors_allowed_origins`, `cors_allowed_methods`, `cors_allowed_headers` config parameters.  ## Arguments  Arguments which expect strings or byte arrays may be passed as quoted strings, like `\"abc\"` or as `0x`-prefixed strings, like `0x616263`.  ## URI/HTTP  A REST like interface.      curl localhost:26657/block?height=5  ## JSONRPC/HTTP  JSONRPC requests can be POST'd to the root RPC endpoint via HTTP.      curl --header \"Content-Type: application/json\" --request POST --data '{\"method\": \"block\", \"params\": [\"5\"], \"id\": 1}' localhost:26657  ## JSONRPC/websockets  JSONRPC requests can be also made via websocket. The websocket endpoint is at `/websocket`, e.g. `localhost:26657/websocket`. Asynchronous RPC functions like event `subscribe` and `unsubscribe` are only available via websockets.  Example using https://github.com/hashrocket/ws:      ws ws://localhost:26657/websocket     > { \"jsonrpc\": \"2.0\", \"method\": \"subscribe\", \"params\": [\"tm.event='NewBlock'\"], \"id\": 1 } 

API version: Master
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BlockIDParts struct for BlockIDParts
type BlockIDParts struct {
	Total int32 `json:"total"`
	Hash string `json:"hash"`
}

// NewBlockIDParts instantiates a new BlockIDParts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockIDParts(total int32, hash string) *BlockIDParts {
	this := BlockIDParts{}
	this.Total = total
	this.Hash = hash
	return &this
}

// NewBlockIDPartsWithDefaults instantiates a new BlockIDParts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockIDPartsWithDefaults() *BlockIDParts {
	this := BlockIDParts{}
	return &this
}

// GetTotal returns the Total field value
func (o *BlockIDParts) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *BlockIDParts) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *BlockIDParts) SetTotal(v int32) {
	o.Total = v
}

// GetHash returns the Hash field value
func (o *BlockIDParts) GetHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *BlockIDParts) GetHashOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *BlockIDParts) SetHash(v string) {
	o.Hash = v
}

func (o BlockIDParts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["hash"] = o.Hash
	}
	return json.Marshal(toSerialize)
}

type NullableBlockIDParts struct {
	value *BlockIDParts
	isSet bool
}

func (v NullableBlockIDParts) Get() *BlockIDParts {
	return v.value
}

func (v *NullableBlockIDParts) Set(val *BlockIDParts) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockIDParts) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockIDParts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockIDParts(val *BlockIDParts) *NullableBlockIDParts {
	return &NullableBlockIDParts{value: val, isSet: true}
}

func (v NullableBlockIDParts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockIDParts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

