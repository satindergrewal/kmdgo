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
	//"fmt"
	"encoding/json"
	"errors"
	"strconv"
)

// CoinSupply type
type CoinSupply struct {
	Result struct {
		Result string  `json:"result"`
		Coin   string  `json:"coin"`
		Height int     `json:"height"`
		Supply float64 `json:"supply"`
		Zfunds float64 `json:"zfunds"`
		Sprout float64 `json:"sprout"`
		Total  float64 `json:"total"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// CoinSupply method returns the coin supply information for the indicated block height.
// If no height is given, the method defaults to the blockchain's current height.
func (appName AppType) CoinSupply(params APIParams) (CoinSupply, error) {
	if params[0] == nil {
		params[0] = 100
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `coinsupply`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var coinsupply CoinSupply

	coinsupplyJSON := appName.APICall(&query)
	if coinsupplyJSON == "EMPTY RPC INFO" {
		return coinsupply, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(coinsupplyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(coinsupplyJSON), &coinsupply)
		return coinsupply, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(coinsupplyJSON), &coinsupply)
	return coinsupply, nil
}

// GetBestBlockhash type
type GetBestBlockhash struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

// GetBestBlockhash method returns the hash of the best (tip) block in the longest block chain.
func (appName AppType) GetBestBlockhash() (GetBestBlockhash, error) {
	query := APIQuery{
		Method: `getbestblockhash`,
		Params: `[]`,
	}

	var getbestblockhash GetBestBlockhash

	getbestblockhashJSON := appName.APICall(&query)
	if getbestblockhashJSON == "EMPTY RPC INFO" {
		return getbestblockhash, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getbestblockhashJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbestblockhashJSON), &getbestblockhash)
		return getbestblockhash, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getbestblockhashJSON), &getbestblockhash)
	return getbestblockhash, nil
}

// GetBlock type
type GetBlock struct {
	Result struct {
		Hash             string `json:"hash"`
		Confirmations    int    `json:"confirmations"`
		Rawconfirmations int    `json:"rawconfirmations"`
		Size             int    `json:"size"`
		Height           int    `json:"height"`
		Version          int    `json:"version"`
		Merkleroot       string `json:"merkleroot"`
		Segid            int    `json:"segid"`
		Finalsaplingroot string `json:"finalsaplingroot"`
		Tx               []struct {
			Txid         string `json:"txid"`
			Overwintered bool   `json:"overwintered"`
			Version      int    `json:"version"`
			Locktime     int    `json:"locktime"`
			Vin          []struct {
				Coinbase string `json:"coinbase"`
				Sequence int64  `json:"sequence"`
			} `json:"vin"`
			Vout []struct {
				Value        float64 `json:"value"`
				ValueZat     int     `json:"valueZat"`
				N            int     `json:"n"`
				ScriptPubKey struct {
					Asm       string   `json:"asm"`
					Hex       string   `json:"hex"`
					ReqSigs   int      `json:"reqSigs"`
					Type      string   `json:"type"`
					Addresses []string `json:"addresses"`
				} `json:"scriptPubKey"`
			} `json:"vout"`
			Vjoinsplit []interface{} `json:"vjoinsplit"`
		} `json:"tx"`
		Time       int    `json:"time"`
		Nonce      string `json:"nonce"`
		Solution   string `json:"solution"`
		Bits       string `json:"bits"`
		Difficulty int    `json:"difficulty"`
		Chainwork  string `json:"chainwork"`
		Anchor     string `json:"anchor"`
		Blocktype  string `json:"blocktype"`
		ValuePools []struct {
			ID            string  `json:"id"`
			Monitored     bool    `json:"monitored"`
			ChainValue    float64 `json:"chainValue"`
			ChainValueZat int     `json:"chainValueZat"`
			ValueDelta    float64 `json:"valueDelta"`
			ValueDeltaZat int     `json:"valueDeltaZat"`
		} `json:"valuePools"`
		Previousblockhash string `json:"previousblockhash"`
		Nextblockhash     string `json:"nextblockhash"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetBlock method returns the block's relevant state information.
// The verbose input is optional. The default value is true, and it will return
// a json object with information about the indicated block.
// If verbose is false, the command returns a string that is
// serialized hex-encoded data for the indicated block.
func (appName AppType) GetBlock(params APIParams) (GetBlock, error) {
	if params[1] == nil {
		params[1] = 1
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getblock`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getblock GetBlock

	getblockJSON := appName.APICall(&query)
	if getblockJSON == "EMPTY RPC INFO" {
		return getblock, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblockJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockJSON), &getblock)
		return getblock, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblockJSON), &getblock)
	return getblock, nil
}

// GetBlockchainInfo type
type GetBlockchainInfo struct {
	Result struct {
		Chain                string  `json:"chain"`
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		Bestblockhash        string  `json:"bestblockhash"`
		Difficulty           float64 `json:"difficulty"`
		Verificationprogress float64 `json:"verificationprogress"`
		Chainwork            string  `json:"chainwork"`
		Pruned               bool    `json:"pruned"`
		Commitments          int     `json:"commitments"`
		ValuePools           []struct {
			ID            string  `json:"id"`
			Monitored     bool    `json:"monitored"`
			ChainValue    float64 `json:"chainValue"`
			ChainValueZat int64   `json:"chainValueZat"`
		} `json:"valuePools"`
		Softforks []struct {
			ID      string `json:"id"`
			Version int    `json:"version"`
			Enforce struct {
				Status   bool `json:"status"`
				Found    int  `json:"found"`
				Required int  `json:"required"`
				Window   int  `json:"window"`
			} `json:"enforce"`
			Reject struct {
				Status   bool `json:"status"`
				Found    int  `json:"found"`
				Required int  `json:"required"`
				Window   int  `json:"window"`
			} `json:"reject"`
		} `json:"softforks"`
		Upgrades struct {
			FiveBa81B19 struct {
				Name             string `json:"name"`
				Activationheight int    `json:"activationheight"`
				Status           string `json:"status"`
				Info             string `json:"info"`
			} `json:"5ba81b19"`
			Seven6B809Bb struct {
				Name             string `json:"name"`
				Activationheight int    `json:"activationheight"`
				Status           string `json:"status"`
				Info             string `json:"info"`
			} `json:"76b809bb"`
		} `json:"upgrades"`
		Consensus struct {
			Chaintip  string `json:"chaintip"`
			Nextblock string `json:"nextblock"`
		} `json:"consensus"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetBlockchainInfo method returns a json object containing state information about blockchain processing.
func (appName AppType) GetBlockchainInfo() (GetBlockchainInfo, error) {
	query := APIQuery{
		Method: `getblockchaininfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getblockchaininfo GetBlockchainInfo

	getblockchaininfoJSON := appName.APICall(&query)
	if getblockchaininfoJSON == "EMPTY RPC INFO" {
		return getblockchaininfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblockchaininfoJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockchaininfoJSON), &getblockchaininfo)
		return getblockchaininfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblockchaininfoJSON), &getblockchaininfo)
	return getblockchaininfo, nil
}

// GetBlockCount type
type GetBlockCount struct {
	Result int64  `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// GetBlockCount method returns the number of blocks in the best valid block chain.
func (appName AppType) GetBlockCount() (GetBlockCount, error) {
	query := APIQuery{
		Method: `getblockcount`,
		Params: `[]`,
	}

	var getblockcount GetBlockCount

	getbestblockhashJSON := appName.APICall(&query)
	if getbestblockhashJSON == "EMPTY RPC INFO" {
		return getblockcount, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getbestblockhashJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbestblockhashJSON), &getblockcount)
		return getblockcount, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getbestblockhashJSON), &getblockcount)
	return getblockcount, nil
}

// GetBlockHash type
type GetBlockHash struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// GetBlockHash method returns the hash of the indicated block index, according to the best blockchain at the time provided.
func (appName AppType) GetBlockHash(h int) (GetBlockHash, error) {
	query := APIQuery{
		Method: `getblockhash`,
		Params: `[` + strconv.Itoa(h) + `]`,
	}
	//fmt.Println(query)

	var getblockhash GetBlockHash

	getblockhashJSON := appName.APICall(&query)
	if getblockhashJSON == "EMPTY RPC INFO" {
		return getblockhash, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblockhashJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockhashJSON), &getblockhash)
		return getblockhash, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblockhashJSON), &getblockhash)
	return getblockhash, nil
}

// GetBlockHeader type
type GetBlockHeader struct {
	Result struct {
		Hash              string  `json:"hash"`
		Confirmations     int     `json:"confirmations"`
		Rawconfirmations  int     `json:"rawconfirmations"`
		Height            int     `json:"height"`
		Version           int     `json:"version"`
		Merkleroot        string  `json:"merkleroot"`
		Finalsaplingroot  string  `json:"finalsaplingroot"`
		Time              int     `json:"time"`
		Nonce             string  `json:"nonce"`
		Solution          string  `json:"solution"`
		Bits              string  `json:"bits"`
		Difficulty        float64 `json:"difficulty"`
		Chainwork         string  `json:"chainwork"`
		Segid             int     `json:"segid"`
		Previousblockhash string  `json:"previousblockhash"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetBlockHeader method returns information about the indicated block.
// The verbose input is optional. If verbose is false, the method returns a string that is serialized,
// hex-encoded data for the indicated blockheader. If verbose is true,
// the method returns a json object with information about the indicated blockheader.
func (appName AppType) GetBlockHeader(params APIParams) (GetBlockHeader, error) {
	if params[1] == nil {
		params[1] = true
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `getblockheader`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getblockheader GetBlockHeader

	getblockheaderJSON := appName.APICall(&query)
	if getblockheaderJSON == "EMPTY RPC INFO" {
		return getblockheader, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getblockheaderJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockheaderJSON), &getblockheader)
		return getblockheader, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getblockheaderJSON), &getblockheader)
	return getblockheader, nil
}

// GetChainTips type
type GetChainTips struct {
	Result []struct {
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Branchlen int    `json:"branchlen"`
		Status    string `json:"status"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetChainTips method returns information about all known tips in the block tree,
// including the main chain and any orphaned branches.
func (appName AppType) GetChainTips() (GetChainTips, error) {
	query := APIQuery{
		Method: `getchaintips`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getchaintips GetChainTips

	getchaintipsJSON := appName.APICall(&query)
	if getchaintipsJSON == "EMPTY RPC INFO" {
		return getchaintips, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getchaintipsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getchaintipsJSON), &getchaintips)
		return getchaintips, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getchaintipsJSON), &getchaintips)
	return getchaintips, nil
}

// GetDifficulty type
type GetDifficulty struct {
	Result float64 `json:"result"`
	Error  Error   `json:"error"`
	ID     string  `json:"id"`
}

// GetDifficulty method returns the proof-of-work difficulty as a multiple of the minimum difficulty.
func (appName AppType) GetDifficulty() (GetDifficulty, error) {
	query := APIQuery{
		Method: `getdifficulty`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getdifficulty GetDifficulty

	getdifficultyJSON := appName.APICall(&query)
	if getdifficultyJSON == "EMPTY RPC INFO" {
		return getdifficulty, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getdifficultyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdifficultyJSON), &getdifficulty)
		return getdifficulty, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getdifficultyJSON), &getdifficulty)
	return getdifficulty, nil
}

// GetMempoolInfo type
type GetMempoolInfo struct {
	Result struct {
		Size  int `json:"size"`
		Bytes int `json:"bytes"`
		Usage int `json:"usage"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetMempoolInfo method returns details on the active state of the transaction memory pool.
func (appName AppType) GetMempoolInfo() (GetMempoolInfo, error) {
	query := APIQuery{
		Method: `getmempoolinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getmempoolinfo GetMempoolInfo

	getmempoolinfoJSON := appName.APICall(&query)
	if getmempoolinfoJSON == "EMPTY RPC INFO" {
		return getmempoolinfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getmempoolinfoJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmempoolinfoJSON), &getmempoolinfo)
		return getmempoolinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getmempoolinfoJSON), &getmempoolinfo)
	return getmempoolinfo, nil
}

// GetRawMempoolTrue type
type GetRawMempoolTrue struct {
	Result map[string]RawMempoolTrue `json:"result"`
	Error  Error                     `json:"error"`
	ID     string                    `json:"id"`
}

// RawMempoolTrue type
type RawMempoolTrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

// GetRawMempoolFalse type
type GetRawMempoolFalse struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

// GetRawMempoolTrue method returns all transaction ids in the memory pool as a json array of transaction ids.
// The verbose input is optional and is false by default.
// When it is true, the method instead returns a json object with various related data.
func (appName AppType) GetRawMempoolTrue(b bool) (GetRawMempoolTrue, error) {
	query := APIQuery{
		Method: `getrawmempool`,
		Params: `[` + strconv.FormatBool(b) + `]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolTrue

	getrawmempoolJSON := appName.APICall(&query)
	if getrawmempoolJSON == "EMPTY RPC INFO" {
		return getrawmempool, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getrawmempoolJSON)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJSON), &getrawmempool)
		return getrawmempool, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getrawmempoolJSON), &getrawmempool)
	return getrawmempool, nil
}

// GetRawMempoolFalse method returns all transaction ids in the memory pool as a json array of transaction ids.
// The verbose input is optional and is false by default.
// When it is true, the method instead returns a json object with various related data.
func (appName AppType) GetRawMempoolFalse(b bool) (GetRawMempoolFalse, error) {
	query := APIQuery{
		Method: `getrawmempool`,
		Params: `[` + strconv.FormatBool(b) + `]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolFalse

	getrawmempoolJSON := appName.APICall(&query)
	if getrawmempoolJSON == "EMPTY RPC INFO" {
		return getrawmempool, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(getrawmempoolJSON)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJSON), &getrawmempool)
		return getrawmempool, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getrawmempoolJSON), &getrawmempool)
	return getrawmempool, nil
}

// GetTxOut type
type GetTxOut struct {
	Result struct {
		Bestblock        string  `json:"bestblock"`
		Confirmations    int     `json:"confirmations"`
		Rawconfirmations int     `json:"rawconfirmations"`
		Value            float64 `json:"value"`
		ScriptPubKey     struct {
			Asm       string   `json:"asm"`
			Hex       string   `json:"hex"`
			ReqSigs   int      `json:"reqSigs"`
			Type      string   `json:"type"`
			Addresses []string `json:"addresses"`
		} `json:"scriptPubKey"`
		Version  int  `json:"version"`
		Coinbase bool `json:"coinbase"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetTxOut method returns details about an unspent transaction output.
func (appName AppType) GetTxOut(params APIParams) (GetTxOut, error) {
	if params[2] == nil {
		params[2] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `gettxout`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var gettxout GetTxOut

	gettxoutJSON := appName.APICall(&query)
	if gettxoutJSON == "EMPTY RPC INFO" {
		return gettxout, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(gettxoutJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutJSON), &gettxout)
		return gettxout, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gettxoutJSON), &gettxout)
	return gettxout, nil
}

// GetTxOutProof type
type GetTxOutProof struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// GetTxOutProof method returns a hex-encoded proof showing that the indicated transaction was included in a block.
func (appName AppType) GetTxOutProof(txids string) (GetTxOutProof, error) {
	query := APIQuery{
		Method: `gettxoutproof`,
		Params: `[` + txids + `]`,
	}
	//fmt.Println(query)

	var gettxoutproof GetTxOutProof

	gettxoutproofJSON := appName.APICall(&query)
	if gettxoutproofJSON == "EMPTY RPC INFO" {
		return gettxoutproof, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(gettxoutproofJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutproofJSON), &gettxoutproof)
		return gettxoutproof, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gettxoutproofJSON), &gettxoutproof)
	return gettxoutproof, nil
}

// GetTxOutSetInfo type
type GetTxOutSetInfo struct {
	Result struct {
		Height          int     `json:"height"`
		Bestblock       string  `json:"bestblock"`
		Transactions    int     `json:"transactions"`
		Txouts          int     `json:"txouts"`
		BytesSerialized int     `json:"bytes_serialized"`
		HashSerialized  string  `json:"hash_serialized"`
		TotalAmount     float64 `json:"total_amount"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// GetTxOutSetInfo method returns statistics about the unspent transaction output set.
func (appName AppType) GetTxOutSetInfo() (GetTxOutSetInfo, error) {
	query := APIQuery{
		Method: `gettxoutsetinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var gettxoutsetinfo GetTxOutSetInfo

	gettxoutsetinfoJSON := appName.APICall(&query)
	if gettxoutsetinfoJSON == "EMPTY RPC INFO" {
		return gettxoutsetinfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(gettxoutsetinfoJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutsetinfoJSON), &gettxoutsetinfo)
		return gettxoutsetinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gettxoutsetinfoJSON), &gettxoutsetinfo)
	return gettxoutsetinfo, nil
}

// MinerIDs type
type MinerIDs struct {
	Result struct {
		Mined []struct {
			Notaryid   int    `json:"notaryid,omitempty"`
			KMDaddress string `json:"KMDaddress,omitempty"`
			Pubkey     string `json:"pubkey"`
			Blocks     int    `json:"blocks"`
		} `json:"mined"`
		Numnotaries int `json:"numnotaries"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// MinerIDs method returns information about the notary nodes and external miners at a specific block height.
// The response will calculate results according to the 2000 blocks proceeding the indicated "height" block.
func (appName AppType) MinerIDs(ht string) (MinerIDs, error) {
	query := APIQuery{
		Method: `minerids`,
		Params: `["` + ht + `"]`,
	}
	//fmt.Println(query)

	var minerids MinerIDs

	mineridsJSON := appName.APICall(&query)
	if mineridsJSON == "EMPTY RPC INFO" {
		return minerids, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(mineridsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(mineridsJSON), &minerids)
		return minerids, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(mineridsJSON), &minerids)
	return minerids, nil
}

// Notaries type
type Notaries struct {
	Result struct {
		Notaries []struct {
			Pubkey     string `json:"pubkey"`
			BTCaddress string `json:"BTCaddress"`
			KMDaddress string `json:"KMDaddress"`
		} `json:"notaries"`
		Numnotaries int `json:"numnotaries"`
		Height      int `json:"height"`
		Timestamp   int `json:"timestamp"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

// Notaries method returns the public key, BTC address, and KMD address for each Komodo notary node.
// Either or both of the height and timestamp parameters will suffice
func (appName AppType) Notaries(ht string) (Notaries, error) {
	query := APIQuery{
		Method: `notaries`,
		Params: `["` + ht + `"]`,
	}
	//fmt.Println(query)

	var notaries Notaries

	notariesJSON := appName.APICall(&query)
	if notariesJSON == "EMPTY RPC INFO" {
		return notaries, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(notariesJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(notariesJSON), &notaries)
		return notaries, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(notariesJSON), &notaries)
	return notaries, nil
}

// VerifyChain type
type VerifyChain struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

// VerifyChain method verifies the coin daemon's blockchain database.
func (appName AppType) VerifyChain(params APIParams) (VerifyChain, error) {
	if params[0] == nil {
		params[0] = 3
	}
	if params[1] == nil {
		params[1] = 288
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := APIQuery{
		Method: `verifychain`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var verifychain VerifyChain

	verifychainJSON := appName.APICall(&query)
	if verifychainJSON == "EMPTY RPC INFO" {
		return verifychain, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(verifychainJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifychainJSON), &verifychain)
		return verifychain, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(verifychainJSON), &verifychain)
	return verifychain, nil
}

// VerifyTxOutProof type
type VerifyTxOutProof struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

// VerifyTxOutProof method verifies that a proof points to a transaction in a block.
// It returns the transaction to which the proof is committed,
// or it will throw an RPC error if the block is not in the current best chain
func (appName AppType) VerifyTxOutProof(pf string) (VerifyTxOutProof, error) {
	query := APIQuery{
		Method: `verifytxoutproof`,
		Params: `["` + pf + `"]`,
	}
	//fmt.Println(query)

	var verifytxoutproof VerifyTxOutProof

	verifytxoutproofJSON := appName.APICall(&query)
	if verifytxoutproofJSON == "EMPTY RPC INFO" {
		return verifytxoutproof, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(verifytxoutproofJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifytxoutproofJSON), &verifytxoutproof)
		return verifytxoutproof, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(verifytxoutproofJSON), &verifytxoutproof)
	return verifytxoutproof, nil
}
