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

// DumpConsensusResponseResultRoundStateLastCommit struct for DumpConsensusResponseResultRoundStateLastCommit
type DumpConsensusResponseResultRoundStateLastCommit struct {
	Votes []string `json:"votes"`
	VotesBitArray string `json:"votes_bit_array"`
	PeerMaj23s map[string]interface{} `json:"peer_maj_23s"`
}

// NewDumpConsensusResponseResultRoundStateLastCommit instantiates a new DumpConsensusResponseResultRoundStateLastCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDumpConsensusResponseResultRoundStateLastCommit(votes []string, votesBitArray string, peerMaj23s map[string]interface{}) *DumpConsensusResponseResultRoundStateLastCommit {
	this := DumpConsensusResponseResultRoundStateLastCommit{}
	this.Votes = votes
	this.VotesBitArray = votesBitArray
	this.PeerMaj23s = peerMaj23s
	return &this
}

// NewDumpConsensusResponseResultRoundStateLastCommitWithDefaults instantiates a new DumpConsensusResponseResultRoundStateLastCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDumpConsensusResponseResultRoundStateLastCommitWithDefaults() *DumpConsensusResponseResultRoundStateLastCommit {
	this := DumpConsensusResponseResultRoundStateLastCommit{}
	return &this
}

// GetVotes returns the Votes field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetVotes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Votes
}

// GetVotesOk returns a tuple with the Votes field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetVotesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Votes, true
}

// SetVotes sets field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) SetVotes(v []string) {
	o.Votes = v
}

// GetVotesBitArray returns the VotesBitArray field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetVotesBitArray() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VotesBitArray
}

// GetVotesBitArrayOk returns a tuple with the VotesBitArray field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetVotesBitArrayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VotesBitArray, true
}

// SetVotesBitArray sets field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) SetVotesBitArray(v string) {
	o.VotesBitArray = v
}

// GetPeerMaj23s returns the PeerMaj23s field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetPeerMaj23s() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.PeerMaj23s
}

// GetPeerMaj23sOk returns a tuple with the PeerMaj23s field value
// and a boolean to check if the value has been set.
func (o *DumpConsensusResponseResultRoundStateLastCommit) GetPeerMaj23sOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PeerMaj23s, true
}

// SetPeerMaj23s sets field value
func (o *DumpConsensusResponseResultRoundStateLastCommit) SetPeerMaj23s(v map[string]interface{}) {
	o.PeerMaj23s = v
}

func (o DumpConsensusResponseResultRoundStateLastCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["votes"] = o.Votes
	}
	if true {
		toSerialize["votes_bit_array"] = o.VotesBitArray
	}
	if true {
		toSerialize["peer_maj_23s"] = o.PeerMaj23s
	}
	return json.Marshal(toSerialize)
}

type NullableDumpConsensusResponseResultRoundStateLastCommit struct {
	value *DumpConsensusResponseResultRoundStateLastCommit
	isSet bool
}

func (v NullableDumpConsensusResponseResultRoundStateLastCommit) Get() *DumpConsensusResponseResultRoundStateLastCommit {
	return v.value
}

func (v *NullableDumpConsensusResponseResultRoundStateLastCommit) Set(val *DumpConsensusResponseResultRoundStateLastCommit) {
	v.value = val
	v.isSet = true
}

func (v NullableDumpConsensusResponseResultRoundStateLastCommit) IsSet() bool {
	return v.isSet
}

func (v *NullableDumpConsensusResponseResultRoundStateLastCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDumpConsensusResponseResultRoundStateLastCommit(val *DumpConsensusResponseResultRoundStateLastCommit) *NullableDumpConsensusResponseResultRoundStateLastCommit {
	return &NullableDumpConsensusResponseResultRoundStateLastCommit{value: val, isSet: true}
}

func (v NullableDumpConsensusResponseResultRoundStateLastCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDumpConsensusResponseResultRoundStateLastCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

