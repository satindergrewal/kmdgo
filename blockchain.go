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

func (appName AppType) CoinSupply(params APIParams) (CoinSupply, error) {
	if params[0] == nil {
		params[0] = 100
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `coinsupply`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var coinsupply CoinSupply

	coinsupplyJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(coinsupplyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
		return coinsupply, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(coinsupplyJson), &coinsupply)
	return coinsupply, nil
}

type GetBestBlockhash struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) GetBestBlockhash() (GetBestBlockhash, error) {
	query := APIQuery{
		Method: `getbestblockhash`,
		Params: `[]`,
	}

	var getbestblockhash GetBestBlockhash

	getbestblockhashJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getbestblockhashJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
		return getbestblockhash, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getbestblockhashJson), &getbestblockhash)
	return getbestblockhash, nil
}

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

func (appName AppType) GetBlock(params APIParams) (GetBlock, error) {
	if params[1] == nil {
		params[1] = 1
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getblock`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getblock GetBlock

	getblockJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblockJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockJson), &getblock)
		return getblock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockJson), &getblock)
	return getblock, nil
}

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

func (appName AppType) GetBlockchainInfo() (GetBlockchainInfo, error) {
	query := APIQuery{
		Method: `getblockchaininfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getblockchaininfo GetBlockchainInfo

	getblockchaininfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblockchaininfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockchaininfoJson), &getblockchaininfo)
		return getblockchaininfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockchaininfoJson), &getblockchaininfo)
	return getblockchaininfo, nil
}

type GetBlockCount struct {
	Result int64  `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetBlockCount() (GetBlockCount, error) {
	query := APIQuery{
		Method: `getblockcount`,
		Params: `[]`,
	}

	var getblockcount GetBlockCount

	getbestblockhashJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getbestblockhashJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbestblockhashJson), &getblockcount)
		return getblockcount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getbestblockhashJson), &getblockcount)
	return getblockcount, nil
}

type GetBlockHash struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetBlockHash(h int) (GetBlockHash, error) {
	query := APIQuery{
		Method: `getblockhash`,
		Params: `[` + strconv.Itoa(h) + `]`,
	}
	//fmt.Println(query)

	var getblockhash GetBlockHash

	getblockhashJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblockhashJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockhashJson), &getblockhash)
		return getblockhash, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockhashJson), &getblockhash)
	return getblockhash, nil
}

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

func (appName AppType) GetBlockHeader(params APIParams) (GetBlockHeader, error) {
	if params[1] == nil {
		params[1] = true
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getblockheader`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getblockheader GetBlockHeader

	getblockheaderJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getblockheaderJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getblockheaderJson), &getblockheader)
		return getblockheader, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getblockheaderJson), &getblockheader)
	return getblockheader, nil
}

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

func (appName AppType) GetChainTips() (GetChainTips, error) {
	query := APIQuery{
		Method: `getchaintips`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getchaintips GetChainTips

	getchaintipsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getchaintipsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getchaintipsJson), &getchaintips)
		return getchaintips, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getchaintipsJson), &getchaintips)
	return getchaintips, nil
}

type GetDifficulty struct {
	Result float64 `json:"result"`
	Error  Error   `json:"error"`
	ID     string  `json:"id"`
}

func (appName AppType) GetDifficulty() (GetDifficulty, error) {
	query := APIQuery{
		Method: `getdifficulty`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getdifficulty GetDifficulty

	getdifficultyJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getdifficultyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getdifficultyJson), &getdifficulty)
		return getdifficulty, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getdifficultyJson), &getdifficulty)
	return getdifficulty, nil
}

type GetMempoolInfo struct {
	Result struct {
		Size  int `json:"size"`
		Bytes int `json:"bytes"`
		Usage int `json:"usage"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetMempoolInfo() (GetMempoolInfo, error) {
	query := APIQuery{
		Method: `getmempoolinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getmempoolinfo GetMempoolInfo

	getmempoolinfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(getmempoolinfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getmempoolinfoJson), &getmempoolinfo)
		return getmempoolinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getmempoolinfoJson), &getmempoolinfo)
	return getmempoolinfo, nil
}

type GetRawMempoolTrue struct {
	Result map[string]RawMempoolTrue `json:"result"`
	Error  Error                     `json:"error"`
	ID     string                    `json:"id"`
}

type RawMempoolTrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

type GetRawMempoolFalse struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) GetRawMempoolTrue(b bool) (GetRawMempoolTrue, error) {
	query := APIQuery{
		Method: `getrawmempool`,
		Params: `[` + strconv.FormatBool(b) + `]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolTrue

	getrawmempoolJson := appName.APICall(query)
	//fmt.Println(getrawmempoolJson)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
		return getrawmempool, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
	return getrawmempool, nil
}

func (appName AppType) GetRawMempoolFalse(b bool) (GetRawMempoolFalse, error) {
	query := APIQuery{
		Method: `getrawmempool`,
		Params: `[` + strconv.FormatBool(b) + `]`,
	}
	//fmt.Println(query)

	var getrawmempool GetRawMempoolFalse

	getrawmempoolJson := appName.APICall(query)
	//fmt.Println(getrawmempoolJson)

	var result APIResult

	json.Unmarshal([]byte(getrawmempoolJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
		return getrawmempool, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawmempoolJson), &getrawmempool)
	return getrawmempool, nil
}

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

func (appName AppType) GetTxOut(params APIParams) (GetTxOut, error) {
	if params[2] == nil {
		params[2] = false
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `gettxout`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var gettxout GetTxOut

	gettxoutJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutJson), &gettxout)
		return gettxout, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutJson), &gettxout)
	return gettxout, nil
}

type GetTxOutProof struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetTxOutProof(txids string) (GetTxOutProof, error) {
	query := APIQuery{
		Method: `gettxoutproof`,
		Params: `[` + txids + `]`,
	}
	//fmt.Println(query)

	var gettxoutproof GetTxOutProof

	gettxoutproofJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutproofJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutproofJson), &gettxoutproof)
		return gettxoutproof, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutproofJson), &gettxoutproof)
	return gettxoutproof, nil
}

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

func (appName AppType) GetTxOutSetInfo() (GetTxOutSetInfo, error) {
	query := APIQuery{
		Method: `gettxoutsetinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var gettxoutsetinfo GetTxOutSetInfo

	gettxoutsetinfoJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(gettxoutsetinfoJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettxoutsetinfoJson), &gettxoutsetinfo)
		return gettxoutsetinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettxoutsetinfoJson), &gettxoutsetinfo)
	return gettxoutsetinfo, nil
}

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

func (appName AppType) MinerIDs(ht string) (MinerIDs, error) {
	query := APIQuery{
		Method: `minerids`,
		Params: `["` + ht + `"]`,
	}
	//fmt.Println(query)

	var minerids MinerIDs

	mineridsJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(mineridsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(mineridsJson), &minerids)
		return minerids, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(mineridsJson), &minerids)
	return minerids, nil
}

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

func (appName AppType) Notaries(ht string) (Notaries, error) {
	query := APIQuery{
		Method: `notaries`,
		Params: `["` + ht + `"]`,
	}
	//fmt.Println(query)

	var notaries Notaries

	notariesJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(notariesJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(notariesJson), &notaries)
		return notaries, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(notariesJson), &notaries)
	return notaries, nil
}

type VerifyChain struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) VerifyChain(params APIParams) (VerifyChain, error) {
	if params[0] == nil {
		params[0] = 3
	}
	if params[1] == nil {
		params[1] = 288
	}
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `verifychain`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var verifychain VerifyChain

	verifychainJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(verifychainJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifychainJson), &verifychain)
		return verifychain, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(verifychainJson), &verifychain)
	return verifychain, nil
}

type VerifyTxOutProof struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) VerifyTxOutProof(pf string) (VerifyTxOutProof, error) {
	query := APIQuery{
		Method: `verifytxoutproof`,
		Params: `["` + pf + `"]`,
	}
	//fmt.Println(query)

	var verifytxoutproof VerifyTxOutProof

	verifytxoutproofJson := appName.APICall(query)

	var result APIResult

	json.Unmarshal([]byte(verifytxoutproofJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(verifytxoutproofJson), &verifytxoutproof)
		return verifytxoutproof, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(verifytxoutproofJson), &verifytxoutproof)
	return verifytxoutproof, nil
}
