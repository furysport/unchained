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

// BlockResultsResponseResult struct for BlockResultsResponseResult
type BlockResultsResponseResult struct {
	Height string `json:"height"`
	TxsResults []BlockResultsResponseResultTxsResults `json:"txs_results,omitempty"`
	TotalGasUsed *string `json:"total_gas_used,omitempty"`
	BeginBlockEvents []BlockResultsResponseResultEvents `json:"begin_block_events,omitempty"`
	EndBlock []BlockResultsResponseResultEvents `json:"end_block,omitempty"`
	ValidatorUpdates []BlockResultsResponseResultValidatorUpdates `json:"validator_updates,omitempty"`
	ConsensusParamsUpdates NullableConsensusParams `json:"consensus_params_updates,omitempty"`
}

// NewBlockResultsResponseResult instantiates a new BlockResultsResponseResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBlockResultsResponseResult(height string) *BlockResultsResponseResult {
	this := BlockResultsResponseResult{}
	this.Height = height
	return &this
}

// NewBlockResultsResponseResultWithDefaults instantiates a new BlockResultsResponseResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBlockResultsResponseResultWithDefaults() *BlockResultsResponseResult {
	this := BlockResultsResponseResult{}
	return &this
}

// GetHeight returns the Height field value
func (o *BlockResultsResponseResult) GetHeight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *BlockResultsResponseResult) GetHeightOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *BlockResultsResponseResult) SetHeight(v string) {
	o.Height = v
}

// GetTxsResults returns the TxsResults field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockResultsResponseResult) GetTxsResults() []BlockResultsResponseResultTxsResults {
	if o == nil  {
		var ret []BlockResultsResponseResultTxsResults
		return ret
	}
	return o.TxsResults
}

// GetTxsResultsOk returns a tuple with the TxsResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockResultsResponseResult) GetTxsResultsOk() (*[]BlockResultsResponseResultTxsResults, bool) {
	if o == nil || o.TxsResults == nil {
		return nil, false
	}
	return &o.TxsResults, true
}

// HasTxsResults returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasTxsResults() bool {
	if o != nil && o.TxsResults != nil {
		return true
	}

	return false
}

// SetTxsResults gets a reference to the given []BlockResultsResponseResultTxsResults and assigns it to the TxsResults field.
func (o *BlockResultsResponseResult) SetTxsResults(v []BlockResultsResponseResultTxsResults) {
	o.TxsResults = v
}

// GetTotalGasUsed returns the TotalGasUsed field value if set, zero value otherwise.
func (o *BlockResultsResponseResult) GetTotalGasUsed() string {
	if o == nil || o.TotalGasUsed == nil {
		var ret string
		return ret
	}
	return *o.TotalGasUsed
}

// GetTotalGasUsedOk returns a tuple with the TotalGasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BlockResultsResponseResult) GetTotalGasUsedOk() (*string, bool) {
	if o == nil || o.TotalGasUsed == nil {
		return nil, false
	}
	return o.TotalGasUsed, true
}

// HasTotalGasUsed returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasTotalGasUsed() bool {
	if o != nil && o.TotalGasUsed != nil {
		return true
	}

	return false
}

// SetTotalGasUsed gets a reference to the given string and assigns it to the TotalGasUsed field.
func (o *BlockResultsResponseResult) SetTotalGasUsed(v string) {
	o.TotalGasUsed = &v
}

// GetBeginBlockEvents returns the BeginBlockEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockResultsResponseResult) GetBeginBlockEvents() []BlockResultsResponseResultEvents {
	if o == nil  {
		var ret []BlockResultsResponseResultEvents
		return ret
	}
	return o.BeginBlockEvents
}

// GetBeginBlockEventsOk returns a tuple with the BeginBlockEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockResultsResponseResult) GetBeginBlockEventsOk() (*[]BlockResultsResponseResultEvents, bool) {
	if o == nil || o.BeginBlockEvents == nil {
		return nil, false
	}
	return &o.BeginBlockEvents, true
}

// HasBeginBlockEvents returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasBeginBlockEvents() bool {
	if o != nil && o.BeginBlockEvents != nil {
		return true
	}

	return false
}

// SetBeginBlockEvents gets a reference to the given []BlockResultsResponseResultEvents and assigns it to the BeginBlockEvents field.
func (o *BlockResultsResponseResult) SetBeginBlockEvents(v []BlockResultsResponseResultEvents) {
	o.BeginBlockEvents = v
}

// GetEndBlock returns the EndBlock field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockResultsResponseResult) GetEndBlock() []BlockResultsResponseResultEvents {
	if o == nil  {
		var ret []BlockResultsResponseResultEvents
		return ret
	}
	return o.EndBlock
}

// GetEndBlockOk returns a tuple with the EndBlock field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockResultsResponseResult) GetEndBlockOk() (*[]BlockResultsResponseResultEvents, bool) {
	if o == nil || o.EndBlock == nil {
		return nil, false
	}
	return &o.EndBlock, true
}

// HasEndBlock returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasEndBlock() bool {
	if o != nil && o.EndBlock != nil {
		return true
	}

	return false
}

// SetEndBlock gets a reference to the given []BlockResultsResponseResultEvents and assigns it to the EndBlock field.
func (o *BlockResultsResponseResult) SetEndBlock(v []BlockResultsResponseResultEvents) {
	o.EndBlock = v
}

// GetValidatorUpdates returns the ValidatorUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockResultsResponseResult) GetValidatorUpdates() []BlockResultsResponseResultValidatorUpdates {
	if o == nil  {
		var ret []BlockResultsResponseResultValidatorUpdates
		return ret
	}
	return o.ValidatorUpdates
}

// GetValidatorUpdatesOk returns a tuple with the ValidatorUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockResultsResponseResult) GetValidatorUpdatesOk() (*[]BlockResultsResponseResultValidatorUpdates, bool) {
	if o == nil || o.ValidatorUpdates == nil {
		return nil, false
	}
	return &o.ValidatorUpdates, true
}

// HasValidatorUpdates returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasValidatorUpdates() bool {
	if o != nil && o.ValidatorUpdates != nil {
		return true
	}

	return false
}

// SetValidatorUpdates gets a reference to the given []BlockResultsResponseResultValidatorUpdates and assigns it to the ValidatorUpdates field.
func (o *BlockResultsResponseResult) SetValidatorUpdates(v []BlockResultsResponseResultValidatorUpdates) {
	o.ValidatorUpdates = v
}

// GetConsensusParamsUpdates returns the ConsensusParamsUpdates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BlockResultsResponseResult) GetConsensusParamsUpdates() ConsensusParams {
	if o == nil || o.ConsensusParamsUpdates.Get() == nil {
		var ret ConsensusParams
		return ret
	}
	return *o.ConsensusParamsUpdates.Get()
}

// GetConsensusParamsUpdatesOk returns a tuple with the ConsensusParamsUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BlockResultsResponseResult) GetConsensusParamsUpdatesOk() (*ConsensusParams, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ConsensusParamsUpdates.Get(), o.ConsensusParamsUpdates.IsSet()
}

// HasConsensusParamsUpdates returns a boolean if a field has been set.
func (o *BlockResultsResponseResult) HasConsensusParamsUpdates() bool {
	if o != nil && o.ConsensusParamsUpdates.IsSet() {
		return true
	}

	return false
}

// SetConsensusParamsUpdates gets a reference to the given NullableConsensusParams and assigns it to the ConsensusParamsUpdates field.
func (o *BlockResultsResponseResult) SetConsensusParamsUpdates(v ConsensusParams) {
	o.ConsensusParamsUpdates.Set(&v)
}
// SetConsensusParamsUpdatesNil sets the value for ConsensusParamsUpdates to be an explicit nil
func (o *BlockResultsResponseResult) SetConsensusParamsUpdatesNil() {
	o.ConsensusParamsUpdates.Set(nil)
}

// UnsetConsensusParamsUpdates ensures that no value is present for ConsensusParamsUpdates, not even an explicit nil
func (o *BlockResultsResponseResult) UnsetConsensusParamsUpdates() {
	o.ConsensusParamsUpdates.Unset()
}

func (o BlockResultsResponseResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["height"] = o.Height
	}
	if o.TxsResults != nil {
		toSerialize["txs_results"] = o.TxsResults
	}
	if o.TotalGasUsed != nil {
		toSerialize["total_gas_used"] = o.TotalGasUsed
	}
	if o.BeginBlockEvents != nil {
		toSerialize["begin_block_events"] = o.BeginBlockEvents
	}
	if o.EndBlock != nil {
		toSerialize["end_block"] = o.EndBlock
	}
	if o.ValidatorUpdates != nil {
		toSerialize["validator_updates"] = o.ValidatorUpdates
	}
	if o.ConsensusParamsUpdates.IsSet() {
		toSerialize["consensus_params_updates"] = o.ConsensusParamsUpdates.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableBlockResultsResponseResult struct {
	value *BlockResultsResponseResult
	isSet bool
}

func (v NullableBlockResultsResponseResult) Get() *BlockResultsResponseResult {
	return v.value
}

func (v *NullableBlockResultsResponseResult) Set(val *BlockResultsResponseResult) {
	v.value = val
	v.isSet = true
}

func (v NullableBlockResultsResponseResult) IsSet() bool {
	return v.isSet
}

func (v *NullableBlockResultsResponseResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBlockResultsResponseResult(val *BlockResultsResponseResult) *NullableBlockResultsResponseResult {
	return &NullableBlockResultsResponseResult{value: val, isSet: true}
}

func (v NullableBlockResultsResponseResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBlockResultsResponseResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

