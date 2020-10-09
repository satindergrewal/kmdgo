// Copyright © 2018-2019 Satinderjit Singh.
//
// See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at
// the top-level directory of this distribution for the individual copyright
// holder information and the developer policies on copyright and licensing.
//
// Unless otherwise agreed in a custom licensing agreement, no part of the
// kmdgo software, including this file may be copied, modified, propagated.
// or distributed except according to the terms contained in the LICENSE file
//
// Removal or modification of this copyright notice is prohibited.

package kmdgo

import (
	"encoding/json"
	"errors"
)

// DEXAnonsend type
type DEXAnonsend struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXAnonsend This method can be used by a user to broadcast any message to the p2p network
// without revealing either the DEX_pubkeys involved
// or the contents of the message to the network. The datablob so created will
// be added to the "Data Mempools" of all the nodes with the parameter -dexp2p set to 1 or 2,
// but can only be decrypted by the node whose DEX_pubkey is destpub33.
// The recipient node can also see the DEX_pubkey of the sender.
func (appName AppType) DEXAnonsend(params APIParams) (DEXAnonsend, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_anonsend`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexAnonsend DEXAnonsend

	dexAnonsendJSON := appName.APICall(&query)
	if dexAnonsendJSON == "EMPTY RPC INFO" {
		return dexAnonsend, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexAnonsendJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexAnonsendJSON), &dexAnonsend)
		return dexAnonsend, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexAnonsendJSON), &dexAnonsend)
	return dexAnonsend, nil
}

// DEXBroadcast type
type DEXBroadcast struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXBroadcast method can be used to broadcast any data to the p2p network,
// which will be added to the "Data Mempools" of all the nodes with the parameter -dexp2p set to 1 or 2.
func (appName AppType) DEXBroadcast(params APIParams) (DEXBroadcast, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_broadcast`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexBroadcast DEXBroadcast

	dexBroadcastJSON := appName.APICall(&query)
	if dexBroadcastJSON == "EMPTY RPC INFO" {
		return dexBroadcast, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexBroadcastJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexBroadcastJSON), &dexBroadcast)
		return dexBroadcast, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexBroadcastJSON), &dexBroadcast)
	return dexBroadcast, nil
}

// DEXCancel type
type DEXCancel struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXCancel This method can be used to cancel an order issued by the user's node.
// A node can cancel only the orders that were broadcasted using its current DEX_pubkey.
// Orders that are broadcasted without being authenticated by a pubkey can not be canceled.
func (appName AppType) DEXCancel(params APIParams) (DEXCancel, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_cancel`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexCancel DEXCancel

	dexCancelJSON := appName.APICall(&query)
	if dexCancelJSON == "EMPTY RPC INFO" {
		return dexCancel, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexCancelJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexCancelJSON), &dexCancel)
		return dexCancel, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexCancelJSON), &dexCancel)
	return dexCancel, nil
}

// DEXGet type
type DEXGet struct {
	Result struct {
		Timestamp    int    `json:"timestamp"`
		ID           int    `json:"id"`
		Hash         string `json:"hash"`
		TagA         string `json:"tagA"`
		TagB         string `json:"tagB"`
		Pubkey       string `json:"pubkey"`
		Payload      string `json:"payload"`
		Hex          int    `json:"hex"`
		Decrypted    string `json:"decrypted"`
		Decryptedhex int    `json:"decryptedhex"`
		Senderpub    string `json:"senderpub"`
		AmountA      string `json:"amountA"`
		AmountB      string `json:"amountB"`
		Priority     int    `json:"priority"`
		Recvtime     int    `json:"recvtime"`
		Cancelled    int    `json:"cancelled"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXGet method returns an order's data by its id.
func (appName AppType) DEXGet(params APIParams) (DEXGet, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_get`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexGet DEXGet

	dexGetJSON := appName.APICall(&query)
	if dexGetJSON == "EMPTY RPC INFO" {
		return dexGet, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexGetJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexGetJSON), &dexGet)
		return dexGet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexGetJSON), &dexGet)
	return dexGet, nil
}

// DEXList type
type DEXList struct {
	Result struct {
		Result  string `json:"result"`
		Matches []struct {
			Timestamp    int    `json:"timestamp"`
			ID           int64  `json:"id"`
			Hash         string `json:"hash"`
			TagA         string `json:"tagA"`
			TagB         string `json:"tagB"`
			Pubkey       string `json:"pubkey"`
			Payload      string `json:"payload"`
			Hex          int    `json:"hex"`
			Decrypted    string `json:"decrypted"`
			Decryptedhex int    `json:"decryptedhex"`
			Senderpub    string `json:"senderpub"`
			Error        string `json:"error"`
			AmountA      string `json:"amountA"`
			AmountB      string `json:"amountB"`
			Priority     int    `json:"priority"`
			Recvtime     int    `json:"recvtime"`
			Cancelled    int    `json:"cancelled"`
		} `json:"matches"`
		TagA   string `json:"tagA"`
		TagB   string `json:"tagB"`
		Pubkey string `json:"pubkey"`
		N      int    `json:"n"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXList method can be used to filter and list data from the "Data Mempool" of the node.
// Each specified filter narrows the list down to the datablobs that match it exactly.
// If a filter is specified as "" or 0, it matches all the values a datablob might have for the filter.
func (appName AppType) DEXList(params APIParams) (DEXList, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_list`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexList DEXList

	dexListJSON := appName.APICall(&query)
	if dexListJSON == "EMPTY RPC INFO" {
		return dexList, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexListJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexListJSON), &dexList)
		return dexList, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexListJSON), &dexList)
	return dexList, nil
}

// DEXOrderbook type
type DEXOrderbook struct {
	Result struct {
		Asks []struct {
			Price      string `json:"price"`
			Baseamount string `json:"baseamount"`
			Relamount  string `json:"relamount"`
			Priority   int    `json:"priority"`
			Pubkey     string `json:"pubkey"`
			Timestamp  int    `json:"timestamp"`
			Hash       string `json:"hash"`
			ID         int64  `json:"id"`
		} `json:"asks"`
		Bids []struct {
			Price      string `json:"price"`
			Baseamount string `json:"baseamount"`
			Relamount  string `json:"relamount"`
			Priority   int    `json:"priority"`
			Pubkey     string `json:"pubkey"`
			Timestamp  int    `json:"timestamp"`
			Hash       string `json:"hash"`
			ID         int64  `json:"id"`
		} `json:"bids"`
		Base string `json:"base"`
		Rel  string `json:"rel"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXOrderbook method interprets the datablobs as orders for AtomicDEX
// and displays relevant data for each order that matches the filters applied through the parameters.
func (appName AppType) DEXOrderbook(params APIParams) (DEXOrderbook, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_orderbook`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexOrderbook DEXOrderbook

	dexOrderbookJSON := appName.APICall(&query)
	if dexOrderbookJSON == "EMPTY RPC INFO" {
		return dexOrderbook, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexOrderbookJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexOrderbookJSON), &dexOrderbook)
		return dexOrderbook, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexOrderbookJSON), &dexOrderbook)
	return dexOrderbook, nil
}

// DEXPublish type
type DEXPublish struct {
	Result struct {
		Fname       string `json:"fname"`
		ID          int    `json:"id"`
		Senderpub   string `json:"senderpub"`
		Filesize    int    `json:"filesize"`
		Fragments   int    `json:"fragments"`
		Numlocators int    `json:"numlocators"`
		Filehash    string `json:"filehash"`
		Checkhash   string `json:"checkhash"`
		Result      string `json:"result"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXPublish method allows user to publish file to p2p network.
// By default it looks for the file in followin directories based on different OS:
// Linux: $HOME directory path = $HOME/dexp2p = example: /home/satinder
// Mac: $HOME directory path = $HOME/dexp2p = example: /Users/satinder
// Windows: %AppData% directory path = %AppData%\dexp2p = example: C:\Users\satinder\AppData\dexp2p
func (appName AppType) DEXPublish(params APIParams) (DEXPublish, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_publish`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexPublish DEXPublish

	dexPublishJSON := appName.APICall(&query)
	if dexPublishJSON == "EMPTY RPC INFO" {
		return dexPublish, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexPublishJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexPublishJSON), &dexPublish)
		return dexPublish, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexPublishJSON), &dexPublish)
	return dexPublish, nil
}

// DEXSetPubKey type
type DEXSetPubKey struct {
	Result struct {
		Result            string `json:"result"`
		PublishablePubkey string `json:"publishable_pubkey"`
		Secpkey           string `json:"secpkey"`
		Recvaddr          string `json:"recvaddr"`
		RecvZaddr         string `json:"recvZaddr"`
		Handle            string `json:"handle"`
		Txpowbits         int    `json:"txpowbits"`
		Vip               int    `json:"vip"`
		Cmdpriority       int    `json:"cmdpriority"`
		Perfstats         string `json:"perfstats"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXSetPubKey sets the public key for blockchain in case it is not started with already a pubkey with paramater `-pubkey=` at it's daemon start.
// This method will only set the public key once. Once set, and in case this pubkey needs to be changed, the daemon must be restarted.
// It is a pubkey value which is also displayed in output to method "validateaddress".
func (appName AppType) DEXSetPubKey(params APIParams) (DEXSetPubKey, error) {

	paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_setpubkey`,
		Params: string(paramsJSON),
	}
	// fmt.Println(query)

	var dexSetPubKey DEXSetPubKey

	dexSetPubKeyJSON := appName.APICall(&query)
	if dexSetPubKeyJSON == "EMPTY RPC INFO" {
		return dexSetPubKey, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexSetPubKeyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexSetPubKeyJSON), &dexSetPubKey)
		return dexSetPubKey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexSetPubKeyJSON), &dexSetPubKey)
	return dexSetPubKey, nil
}

// DEXStats type
type DEXStats struct {
	Result struct {
		Result            string `json:"result"`
		PublishablePubkey string `json:"publishable_pubkey"`
		Secpkey           string `json:"secpkey"`
		Recvaddr          string `json:"recvaddr"`
		RecvZaddr         string `json:"recvZaddr"`
		Handle            string `json:"handle"`
		Txpowbits         int    `json:"txpowbits"`
		Vip               int    `json:"vip"`
		Cmdpriority       int    `json:"cmdpriority"`
		Progress          int    `json:"progress"`
		Perfstats         string `json:"perfstats"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// DEXStats method gives info and stats related to the p2p data layer.
func (appName AppType) DEXStats() (DEXStats, error) {

	// paramsJSON, _ := json.Marshal(params)
	// fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `DEX_stats`,
		Params: `[]`,
	}
	// fmt.Println(query)

	var dexStats DEXStats

	dexStatsJSON := appName.APICall(&query)
	if dexStatsJSON == "EMPTY RPC INFO" {
		return dexStats, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dexStatsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dexStatsJSON), &dexStats)
		return dexStats, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dexStatsJSON), &dexStats)
	return dexStats, nil
}
