// Copyright © 2018-2020 Satinderjit Singh.
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

type ListUnspent struct {
	Result []struct {
		Txid          string  `json:"txid"`
		Vout          int     `json:"vout"`
		Generated     bool    `json:"generated"`
		Address       string  `json:"address"`
		ScriptPubKey  string  `json:"scriptPubKey"`
		Amount        float64 `json:"amount"`
		Interest      float64 `json:"interest"`
		Confirmations int     `json:"confirmations"`
		Spendable     bool    `json:"spendable"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListUnspent(params APIParams) (ListUnspent, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = 9999999
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `listunspent`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var listunspent ListUnspent

	listunspentJSON := appName.APICall(query)
	if listunspentJSON == "EMPTY RPC INFO" {
		return listunspent, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listunspentJSON)

	var result APIResult

	json.Unmarshal([]byte(listunspentJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listunspentJSON), &listunspent)
		return listunspent, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listunspentJSON), &listunspent)
	return listunspent, nil
}

type AddMultiSigAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) AddMultiSigAddress(params APIParams) (AddMultiSigAddress, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `addmultisigaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var addmultisigaddress AddMultiSigAddress

	addmultisigaddressJSON := appName.APICall(query)
	if addmultisigaddressJSON == "EMPTY RPC INFO" {
		return addmultisigaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(addmultisigaddressJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(addmultisigaddressJSON), &addmultisigaddress)
		return addmultisigaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(addmultisigaddressJSON), &addmultisigaddress)
	return addmultisigaddress, nil
}

type BackupWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) BackupWallet(bkpwlt string) (BackupWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `backupwallet`,
		Params: `["` + bkpwlt + `"]`,
	}
	//fmt.Println(query)

	var backupwallet BackupWallet

	backupwalletJSON := appName.APICall(query)
	if backupwalletJSON == "EMPTY RPC INFO" {
		return backupwallet, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(backupwalletJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(backupwalletJSON), &backupwallet)
		return backupwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(backupwalletJSON), &backupwallet)
	return backupwallet, nil
}

type DumpPrivKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) DumpPrivKey(taddr string) (DumpPrivKey, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `dumpprivkey`,
		Params: `["` + taddr + `"]`,
	}
	//fmt.Println(query)

	var dumpprivkey DumpPrivKey

	dumpprivkeyJSON := appName.APICall(query)
	if dumpprivkeyJSON == "EMPTY RPC INFO" {
		return dumpprivkey, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dumpprivkeyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dumpprivkeyJSON), &dumpprivkey)
		return dumpprivkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dumpprivkeyJSON), &dumpprivkey)
	return dumpprivkey, nil
}

type DumpWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) DumpWallet(dmpwlt string) (DumpWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `dumpwallet`,
		Params: `["` + dmpwlt + `"]`,
	}
	//fmt.Println(query)

	var dumpwallet DumpWallet

	dumpwalletJSON := appName.APICall(query)
	if dumpwalletJSON == "EMPTY RPC INFO" {
		return dumpwallet, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(dumpwalletJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dumpwalletJSON), &dumpwallet)
		return dumpwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(dumpwalletJSON), &dumpwallet)
	return dumpwallet, nil
}

type EncryptWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) EncryptWallet(phrase string) (EncryptWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `encryptwallet`,
		Params: `["` + phrase + `"]`,
	}
	//fmt.Println(query)

	var encryptwallet EncryptWallet

	encryptwalletJSON := appName.APICall(query)
	if encryptwalletJSON == "EMPTY RPC INFO" {
		return encryptwallet, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(encryptwalletJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(encryptwalletJSON), &encryptwallet)
		return encryptwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(encryptwalletJSON), &encryptwallet)
	return encryptwallet, nil
}

type GetAccount struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetAccount(taddr string) (GetAccount, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getaccount`,
		Params: `["` + taddr + `"]`,
	}
	//fmt.Println(query)

	var getaccount GetAccount

	getaccountJSON := appName.APICall(query)
	if getaccountJSON == "EMPTY RPC INFO" {
		return getaccount, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getaccountJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaccountJSON), &getaccount)
		return getaccount, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getaccountJSON), &getaccount)
	return getaccount, nil
}

type GetAccountAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetAccountAddress(actname string) (GetAccountAddress, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getaccountaddress`,
		Params: `["` + actname + `"]`,
	}
	//fmt.Println(query)

	var getaccountaddress GetAccountAddress

	getaccountaddressJSON := appName.APICall(query)
	if getaccountaddressJSON == "EMPTY RPC INFO" {
		return getaccountaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getaccountaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaccountaddressJSON), &getaccountaddress)
		return getaccountaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getaccountaddressJSON), &getaccountaddress)
	return getaccountaddress, nil
}

type GetAddressesByAccount struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) GetAddressesByAccount(actname string) (GetAddressesByAccount, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getaddressesbyaccount`,
		Params: `["` + actname + `"]`,
	}
	//fmt.Println(query)

	var getaddressesbyaccount GetAddressesByAccount

	getaddressesbyaccountJSON := appName.APICall(query)
	if getaddressesbyaccountJSON == "EMPTY RPC INFO" {
		return getaddressesbyaccount, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getaddressesbyaccountJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaddressesbyaccountJSON), &getaddressesbyaccount)
		return getaddressesbyaccount, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getaddressesbyaccountJSON), &getaddressesbyaccount)
	return getaddressesbyaccount, nil
}

type GetBalance struct {
	Result int64  `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetBalance(params APIParams) (GetBalance, error) {
	if params[0] == nil {
		params[0] = "*"
	}
	if params[1] == nil {
		params[1] = 1
	}

	if params[2] == nil {
		params[2] = false
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `getbalance`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getbalance GetBalance

	getbalanceJSON := appName.APICall(query)
	if getbalanceJSON == "EMPTY RPC INFO" {
		return getbalance, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getbalanceJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbalanceJSON), &getbalance)
		return getbalance, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getbalanceJSON), &getbalance)
	return getbalance, nil
}

type GetNewAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetNewAddress() (GetNewAddress, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getnewaddress`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getnewaddress GetNewAddress

	getnewaddressJSON := appName.APICall(query)
	if getnewaddressJSON == "EMPTY RPC INFO" {
		return getnewaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getnewaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnewaddressJSON), &getnewaddress)
		return getnewaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getnewaddressJSON), &getnewaddress)
	return getnewaddress, nil
}

type GetRawChangeAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetRawChangeAddress() (GetRawChangeAddress, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getrawchangeaddress`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getrawchangeaddress GetRawChangeAddress

	getrawchangeaddressJSON := appName.APICall(query)
	if getrawchangeaddressJSON == "EMPTY RPC INFO" {
		return getrawchangeaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getrawchangeaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawchangeaddressJSON), &getrawchangeaddress)
		return getrawchangeaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getrawchangeaddressJSON), &getrawchangeaddress)
	return getrawchangeaddress, nil
}

type GetReceivedByAccount struct {
	Result int64  `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetReceivedByAccount(params APIParams) (GetReceivedByAccount, error) {

	// In call cases set account to blank. ACCOUNTS feature is DEPRICATED. It Should not be used anymore.
	if params[0] == nil {
		params[0] = ``
	} else {
		params[0] = ``
	}

	if params[1] == nil {
		params[1] = 1
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `getreceivedbyaccount`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getreceivedbyaccount GetReceivedByAccount

	getreceivedbyaccountJSON := appName.APICall(query)
	if getreceivedbyaccountJSON == "EMPTY RPC INFO" {
		return getreceivedbyaccount, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getreceivedbyaccountJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getreceivedbyaccountJSON), &getreceivedbyaccount)
		return getreceivedbyaccount, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getreceivedbyaccountJSON), &getreceivedbyaccount)
	return getreceivedbyaccount, nil
}

type GetReceivedByAddress struct {
	Result int64  `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetReceivedByAddress(params APIParams) (GetReceivedByAddress, error) {
	if params[1] == nil {
		params[1] = 1
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `getreceivedbyaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var getreceivedbyaddress GetReceivedByAddress

	getreceivedbyaddressJSON := appName.APICall(query)
	if getreceivedbyaddressJSON == "EMPTY RPC INFO" {
		return getreceivedbyaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getreceivedbyaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getreceivedbyaddressJSON), &getreceivedbyaddress)
		return getreceivedbyaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getreceivedbyaddressJSON), &getreceivedbyaddress)
	return getreceivedbyaddress, nil
}

type GetTransaction struct {
	Result struct {
		Amount           float64       `json:"amount"`
		Fee              float64       `json:"fee"`
		Rawconfirmations int           `json:"rawconfirmations"`
		Confirmations    int           `json:"confirmations"`
		Blockhash        string        `json:"blockhash"`
		Blockindex       int           `json:"blockindex"`
		Blocktime        int           `json:"blocktime"`
		Expiryheight     int           `json:"expiryheight"`
		Txid             string        `json:"txid"`
		Walletconflicts  []interface{} `json:"walletconflicts"`
		Time             int           `json:"time"`
		Timereceived     int           `json:"timereceived"`
		Vjoinsplit       []interface{} `json:"vjoinsplit"`
		Details          []struct {
			InvolvesWatchonly bool    `json:"involvesWatchonly"`
			Account           string  `json:"account"`
			Address           string  `json:"address"`
			Category          string  `json:"category"`
			Amount            float64 `json:"amount"`
			Vout              int     `json:"vout"`
			Fee               float64 `json:"fee,omitempty"`
			Size              int     `json:"size"`
		} `json:"details"`
		Hex string `json:"hex"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetTransaction(params APIParams) (GetTransaction, error) {
	if params[1] == nil {
		params[1] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `gettransaction`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var gettransaction GetTransaction

	gettransactionJSON := appName.APICall(query)
	if gettransactionJSON == "EMPTY RPC INFO" {
		return gettransaction, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(gettransactionJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettransactionJSON), &gettransaction)
		return gettransaction, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(gettransactionJSON), &gettransaction)
	return gettransaction, nil
}

type GetUnconfirmedBalance struct {
	Result float64 `json:"result"`
	Error  Error   `json:"error"`
	ID     string  `json:"id"`
}

func (appName AppType) GetUnconfirmedBalance() (GetUnconfirmedBalance, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getunconfirmedbalance`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getunconfirmedbalance GetUnconfirmedBalance

	getunconfirmedbalanceJSON := appName.APICall(query)
	if getunconfirmedbalanceJSON == "EMPTY RPC INFO" {
		return getunconfirmedbalance, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getunconfirmedbalanceJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getunconfirmedbalanceJSON), &getunconfirmedbalance)
		return getunconfirmedbalance, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getunconfirmedbalanceJSON), &getunconfirmedbalance)
	return getunconfirmedbalance, nil
}

type GetWalletInfo struct {
	Result struct {
		Walletversion      int     `json:"walletversion"`
		Balance            float64 `json:"balance"`
		UnconfirmedBalance float64 `json:"unconfirmed_balance"`
		ImmatureBalance    float64 `json:"immature_balance"`
		Txcount            int     `json:"txcount"`
		Keypoololdest      int     `json:"keypoololdest"`
		Keypoolsize        int     `json:"keypoolsize"`
		UnlockedUntil      int     `json:"unlocked_until"`
		Paytxfee           float64 `json:"paytxfee"`
		Seedfp             string  `json:"seedfp"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) GetWalletInfo() (GetWalletInfo, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `getwalletinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getwalletinfo GetWalletInfo

	getwalletinfoJSON := appName.APICall(query)
	if getwalletinfoJSON == "EMPTY RPC INFO" {
		return getwalletinfo, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(getwalletinfoJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getwalletinfoJSON), &getwalletinfo)
		return getwalletinfo, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(getwalletinfoJSON), &getwalletinfo)
	return getwalletinfo, nil
}

type ImportAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ImportAddress(params APIParams) (ImportAddress, error) {
	if params[0] == nil {
		params[0] = "*"
	}
	if params[1] == nil {
		params[1] = ``
	}

	if params[2] == nil {
		params[2] = true
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `importaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var importaddress ImportAddress

	importaddressJSON := appName.APICall(query)
	if importaddressJSON == "EMPTY RPC INFO" {
		return importaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(importaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importaddressJSON), &importaddress)
		return importaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(importaddressJSON), &importaddress)
	return importaddress, nil
}

type ImportPrivKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ImportPrivKey(params APIParams) (ImportPrivKey, error) {
	if params[0] == nil {
		params[0] = "*"
	}
	if params[1] == nil {
		params[1] = ``
	}

	if params[2] == nil {
		params[2] = true
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `importprivkey`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var importprivkey ImportPrivKey

	importprivkeyJSON := appName.APICall(query)
	if importprivkeyJSON == "EMPTY RPC INFO" {
		return importprivkey, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(importprivkeyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importprivkeyJSON), &importprivkey)
		return importprivkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(importprivkeyJSON), &importprivkey)
	return importprivkey, nil
}

type ImportWallet struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) ImportWallet(wltpth string) (ImportWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `importwallet`,
		Params: `["` + wltpth + `"]`,
	}
	//fmt.Println(query)

	var importwallet ImportWallet

	importwalletJSON := appName.APICall(query)
	if importwalletJSON == "EMPTY RPC INFO" {
		return importwallet, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(importwalletJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importwalletJSON), &importwallet)
		return importwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(importwalletJSON), &importwallet)
	return importwallet, nil
}

type KeyPoolRefill struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) KeyPoolRefill(params APIParams) (KeyPoolRefill, error) {
	if params[0] == nil {
		params[0] = 100
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `keypoolrefill`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var keypoolrefill KeyPoolRefill

	keypoolrefillJSON := appName.APICall(query)
	if keypoolrefillJSON == "EMPTY RPC INFO" {
		return keypoolrefill, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(keypoolrefillJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(keypoolrefillJSON), &keypoolrefill)
		return keypoolrefill, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(keypoolrefillJSON), &keypoolrefill)
	return keypoolrefill, nil
}

type ListAccounts struct {
	Result map[string]float64 `json:"result"`
	Error  Error              `json:"error"`
	ID     string             `json:"id"`
}

func (appName AppType) ListAccounts(params APIParams) (ListAccounts, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `listaccounts`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var listaccounts ListAccounts

	listaccountsJSON := appName.APICall(query)
	if listaccountsJSON == "EMPTY RPC INFO" {
		return listaccounts, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listaccountsJSON)

	var result APIResult

	json.Unmarshal([]byte(listaccountsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listaccountsJSON), &listaccounts)
		return listaccounts, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listaccountsJSON), &listaccounts)
	return listaccounts, nil
}

type ListAddressGroupings struct {
	Result [][][]interface{} `json:"result"`
	Error  Error             `json:"error"`
	ID     string            `json:"id"`
}

func (appName AppType) ListAddressGroupings() (ListAddressGroupings, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `listaddressgroupings`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var listaddressgroupings ListAddressGroupings

	listaddressgroupingsJSON := appName.APICall(query)
	if listaddressgroupingsJSON == "EMPTY RPC INFO" {
		return listaddressgroupings, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listaddressgroupingsJSON)

	var result APIResult

	json.Unmarshal([]byte(listaddressgroupingsJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listaddressgroupingsJSON), &listaddressgroupings)
		return listaddressgroupings, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listaddressgroupingsJSON), &listaddressgroupings)
	return listaddressgroupings, nil
}

type ListLockUnspent struct {
	Result []struct {
		Txid string `json:"txid"`
		Vout int    `json:"vout"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListLockUnspent() (ListLockUnspent, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `listlockunspent`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var listlockunspent ListLockUnspent

	listlockunspentJSON := appName.APICall(query)
	if listlockunspentJSON == "EMPTY RPC INFO" {
		return listlockunspent, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listlockunspentJSON)

	var result APIResult

	json.Unmarshal([]byte(listlockunspentJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listlockunspentJSON), &listlockunspent)
		return listlockunspent, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listlockunspentJSON), &listlockunspent)
	return listlockunspent, nil
}

type ListReceivedByAddress struct {
	Result []struct {
		Address          string        `json:"address"`
		Account          string        `json:"account"`
		Amount           float64       `json:"amount"`
		Rawconfirmations int           `json:"rawconfirmations"`
		Confirmations    int           `json:"confirmations"`
		Txids            []interface{} `json:"txids"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListReceivedByAddress(params APIParams) (ListReceivedByAddress, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = false
	}

	if params[2] == nil {
		params[2] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `listreceivedbyaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var listreceivedbyaddress ListReceivedByAddress

	listreceivedbyaddressJSON := appName.APICall(query)
	if listreceivedbyaddressJSON == "EMPTY RPC INFO" {
		return listreceivedbyaddress, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listreceivedbyaddressJSON)

	var result APIResult

	json.Unmarshal([]byte(listreceivedbyaddressJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listreceivedbyaddressJSON), &listreceivedbyaddress)
		return listreceivedbyaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listreceivedbyaddressJSON), &listreceivedbyaddress)
	return listreceivedbyaddress, nil
}

type ListSinceBlock struct {
	Result struct {
		Transactions []struct {
			Account         string        `json:"account"`
			Address         string        `json:"address"`
			Category        string        `json:"category"`
			Amount          float64       `json:"amount"`
			Vout            int           `json:"vout"`
			Confirmations   int           `json:"confirmations"`
			Generated       bool          `json:"generated"`
			Blockhash       string        `json:"blockhash"`
			Blockindex      int           `json:"blockindex"`
			Blocktime       int           `json:"blocktime"`
			Expiryheight    int           `json:"expiryheight"`
			Txid            string        `json:"txid"`
			Walletconflicts []interface{} `json:"walletconflicts"`
			Time            int           `json:"time"`
			Timereceived    int           `json:"timereceived"`
			Vjoinsplit      []interface{} `json:"vjoinsplit"`
			Size            int           `json:"size"`
		} `json:"transactions"`
		Lastblock string `json:"lastblock"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListSinceBlock(params APIParams) (ListSinceBlock, error) {

	//fmt.Println(len(params))

	if len(params) == 1 {
		if params[0] == nil {
			params[0] = "false"
		}
	}

	if len(params) == 3 {
		if params[2] == nil {
			params[2] = false
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `listsinceblock`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var listsinceblock ListSinceBlock

	listsinceblockJSON := appName.APICall(query)
	if listsinceblockJSON == "EMPTY RPC INFO" {
		return listsinceblock, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(listsinceblockJSON)

	var result APIResult

	json.Unmarshal([]byte(listsinceblockJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listsinceblockJSON), &listsinceblock)
		return listsinceblock, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listsinceblockJSON), &listsinceblock)
	return listsinceblock, nil
}

type ListTransactions struct {
	Result []struct {
		InvolvesWatchonly bool          `json:"involvesWatchonly"`
		Account           string        `json:"account"`
		Address           string        `json:"address"`
		Category          string        `json:"category"`
		Amount            float64       `json:"amount"`
		Vout              int           `json:"vout"`
		Fee               float64       `json:"fee,omitempty"`
		Rawconfirmations  int           `json:"rawconfirmations"`
		Confirmations     int           `json:"confirmations"`
		Blockhash         string        `json:"blockhash"`
		Blockindex        int           `json:"blockindex"`
		Blocktime         int           `json:"blocktime"`
		Expiryheight      int           `json:"expiryheight"`
		Txid              string        `json:"txid"`
		Walletconflicts   []interface{} `json:"walletconflicts"`
		Time              int           `json:"time"`
		Timereceived      int           `json:"timereceived"`
		Vjoinsplit        []interface{} `json:"vjoinsplit"`
		Size              int           `json:"size"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ListTransactions(params APIParams) (ListTransactions, error) {
	if params[0] == nil {
		params[0] = "*"
	}
	if params[1] == nil {
		params[1] = 10
	}

	if params[2] == nil {
		params[2] = 0
	}
	if params[3] == nil {
		params[3] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `listtransactions`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var listtransactions ListTransactions

	listtransactionsJSON := appName.APICall(query)
	if listtransactionsJSON == "EMPTY RPC INFO" {
		return listtransactions, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(listtransactionsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listtransactionsJSON), &listtransactions)
		return listtransactions, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(listtransactionsJSON), &listtransactions)
	return listtransactions, nil
}

type LockUnspent struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) LockUnspent(params APIParams) (LockUnspent, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `lockunspent`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var lockunspent LockUnspent

	lockunspentJSON := appName.APICall(query)
	if lockunspentJSON == "EMPTY RPC INFO" {
		return lockunspent, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(lockunspentJSON)

	var result APIResult

	json.Unmarshal([]byte(lockunspentJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(lockunspentJSON), &lockunspent)
		return lockunspent, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(lockunspentJSON), &lockunspent)
	return lockunspent, nil
}

type SendFrom struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SendFrom(params APIParams) (SendFrom, error) {
	if params[3] == nil {
		params[3] = 1
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `sendfrom`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var sendfrom SendFrom

	sendfromJSON := appName.APICall(query)
	if sendfromJSON == "EMPTY RPC INFO" {
		return sendfrom, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(sendfromJSON)

	var result APIResult

	json.Unmarshal([]byte(sendfromJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendfromJSON), &sendfrom)
		return sendfrom, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(sendfromJSON), &sendfrom)
	return sendfrom, nil
}

type SendMany struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SendMany(params APIParams) (SendMany, error) {
	if len(params) > 2 {
		if params[2] == nil {
			params[2] = 1
		}
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `sendmany`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var sendmany SendMany

	sendmanyJSON := appName.APICall(query)
	if sendmanyJSON == "EMPTY RPC INFO" {
		return sendmany, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(sendmanyJSON)

	var result APIResult

	json.Unmarshal([]byte(sendmanyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendmanyJSON), &sendmany)
		return sendmany, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(sendmanyJSON), &sendmany)
	return sendmany, nil
}

type SendToAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SendToAddress(params APIParams) (SendToAddress, error) {
	if len(params) == 5 {
		if params[4] == nil {
			params[4] = false
		}
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `sendtoaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var sendtoaddress SendToAddress

	sendtoaddressJSON := appName.APICall(query)
	if sendtoaddressJSON == "EMPTY RPC INFO" {
		return sendtoaddress, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(sendtoaddressJSON)

	var result APIResult

	json.Unmarshal([]byte(sendtoaddressJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendtoaddressJSON), &sendtoaddress)
		return sendtoaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(sendtoaddressJSON), &sendtoaddress)
	return sendtoaddress, nil
}

type SetPubKey struct {
	Result struct {
		Pubkey   string `json:"pubkey"`
		Ismine   bool   `json:"ismine"`
		RAddress string `json:"R-address"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) SetPubKey(pubkey string) (SetPubKey, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `setpubkey`,
		Params: `["` + pubkey + `"]`,
	}
	//fmt.Println(query)

	var setpubkey SetPubKey

	setpubkeyJSON := appName.APICall(query)
	if setpubkeyJSON == "EMPTY RPC INFO" {
		return setpubkey, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(setpubkeyJSON)

	var result APIResult

	json.Unmarshal([]byte(setpubkeyJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(setpubkeyJSON), &setpubkey)
		return setpubkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(setpubkeyJSON), &setpubkey)
	return setpubkey, nil
}

type SetTxFee struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SetTxFee(amount float64) (SetTxFee, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `settxfee`,
		Params: `[` + strconv.FormatFloat(amount, 'f', 8, 64) + `]`,
	}
	//fmt.Println(query)

	var settxfee SetTxFee

	settxfeeJSON := appName.APICall(query)
	if settxfeeJSON == "EMPTY RPC INFO" {
		return settxfee, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(settxfeeJSON)

	var result APIResult

	json.Unmarshal([]byte(settxfeeJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(settxfeeJSON), &settxfee)
		return settxfee, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(settxfeeJSON), &settxfee)
	return settxfee, nil
}

type SignMessage struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SignMessage(params APIParams) (SignMessage, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `signmessage`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var signmessage SignMessage

	signmessageJSON := appName.APICall(query)
	if signmessageJSON == "EMPTY RPC INFO" {
		return signmessage, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(signmessageJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(signmessageJSON), &signmessage)
		return signmessage, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(signmessageJSON), &signmessage)
	return signmessage, nil
}

type WalletLock struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletLock() (WalletLock, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `walletlock`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var walletlock WalletLock

	walletlockJSON := appName.APICall(query)
	if walletlockJSON == "EMPTY RPC INFO" {
		return walletlock, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(walletlockJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletlockJSON), &walletlock)
		return walletlock, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(walletlockJSON), &walletlock)
	return walletlock, nil
}

type WalletPassPhrase struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletPassPhrase(params APIParams) (WalletPassPhrase, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `walletpassphrase`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var walletpassphrase WalletPassPhrase

	walletpassphraseJSON := appName.APICall(query)
	if walletpassphraseJSON == "EMPTY RPC INFO" {
		return walletpassphrase, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(walletpassphraseJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletpassphraseJSON), &walletpassphrase)
		return walletpassphrase, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(walletpassphraseJSON), &walletpassphrase)
	return walletpassphrase, nil
}

type WalletPassPhrasechangeChange struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletPassPhrasechangeChange(params APIParams) (WalletPassPhrasechangeChange, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `walletpassphrasechange`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var walletpassphrasechange WalletPassPhrasechangeChange

	walletpassphrasechangeJSON := appName.APICall(query)
	if walletpassphrasechangeJSON == "EMPTY RPC INFO" {
		return walletpassphrasechange, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(walletpassphrasechangeJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletpassphrasechangeJSON), &walletpassphrasechange)
		return walletpassphrasechange, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(walletpassphrasechangeJSON), &walletpassphrasechange)
	return walletpassphrasechange, nil
}

type ZExportKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportKey(zaddr string) (ZExportKey, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_exportkey`,
		Params: `["` + zaddr + `"]`,
	}
	//fmt.Println(query)

	var zexportkey ZExportKey

	zexportkeyJSON := appName.APICall(query)
	if zexportkeyJSON == "EMPTY RPC INFO" {
		return zexportkey, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(zexportkeyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zexportkeyJSON), &zexportkey)
		return zexportkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zexportkeyJSON), &zexportkey)
	return zexportkey, nil
}

type ZExportViewingKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportViewingKey(zaddr string) (ZExportViewingKey, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_exportviewingkey`,
		Params: `["` + zaddr + `"]`,
	}
	//fmt.Println(query)

	var zexportviewingkey ZExportViewingKey

	zexportviewingkeyJSON := appName.APICall(query)
	if zexportviewingkeyJSON == "EMPTY RPC INFO" {
		return zexportviewingkey, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(zexportviewingkeyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zexportviewingkeyJSON), &zexportviewingkey)
		return zexportviewingkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zexportviewingkeyJSON), &zexportviewingkey)
	return zexportviewingkey, nil
}

type ZExportWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportWallet(wltfile string) (ZExportWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_exportwallet`,
		Params: `["` + wltfile + `"]`,
	}
	//fmt.Println(query)

	var zExportwallet ZExportWallet

	zExportwalletJSON := appName.APICall(query)
	if zExportwalletJSON == "EMPTY RPC INFO" {
		return zExportwallet, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zExportwalletJSON)

	var result APIResult

	json.Unmarshal([]byte(zExportwalletJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zExportwalletJSON), &zExportwallet)
		return zExportwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zExportwalletJSON), &zExportwallet)
	return zExportwallet, nil
}

type ZGetBalance struct {
	Result float64 `json:"result"`
	Error  Error   `json:"error"`
	ID     string  `json:"id"`
}

func (appName AppType) ZGetBalance(params APIParams) (ZGetBalance, error) {
	if params[1] == nil {
		params[1] = 1
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_getbalance`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zGetbalance ZGetBalance

	zGetbalanceJSON := appName.APICall(query)
	if zGetbalanceJSON == "EMPTY RPC INFO" {
		return zGetbalance, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zGetbalanceJSON)

	var result APIResult

	json.Unmarshal([]byte(zGetbalanceJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zGetbalanceJSON), &zGetbalance)
		return zGetbalance, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zGetbalanceJSON), &zGetbalance)
	return zGetbalance, nil
}

type ZGetNewAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZGetNewAddress(tp string) (ZGetNewAddress, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_getnewaddress`,
		Params: `["` + tp + `"]`,
	}
	//fmt.Println(query)

	var zgetnewaddress ZGetNewAddress

	zgetnewaddressJSON := appName.APICall(query)
	if zgetnewaddressJSON == "EMPTY RPC INFO" {
		return zgetnewaddress, errors.New("EMPTY RPC INFO")
	}

	var result APIResult

	json.Unmarshal([]byte(zgetnewaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zgetnewaddressJSON), &zgetnewaddress)
		return zgetnewaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zgetnewaddressJSON), &zgetnewaddress)
	return zgetnewaddress, nil
}

type ZGetOperationResult struct {
	Result []struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		CreationTime int    `json:"creation_time"`
		Result       struct {
			Txid string `json:"txid"`
		} `json:"result"`
		ExecutionSecs float64 `json:"execution_secs"`
		Method        string  `json:"method"`
		Params        struct {
			Fromaddress string `json:"fromaddress"`
			Amounts     []struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
			} `json:"amounts"`
			Minconf int     `json:"minconf"`
			Fee     float64 `json:"fee"`
		} `json:"params"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZGetOperationResult(params APIParams) (ZGetOperationResult, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_getoperationresult`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zGetoperationresult ZGetOperationResult

	zGetoperationresultJSON := appName.APICall(query)
	if zGetoperationresultJSON == "EMPTY RPC INFO" {
		return zGetoperationresult, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zGetoperationresultJSON)

	var result APIResult

	json.Unmarshal([]byte(zGetoperationresultJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zGetoperationresultJSON), &zGetoperationresult)
		return zGetoperationresult, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zGetoperationresultJSON), &zGetoperationresult)
	return zGetoperationresult, nil
}

type ZGetOperationStatus struct {
	Result []struct {
		ID           string `json:"id"`
		Status       string `json:"status"`
		CreationTime int    `json:"creation_time"`
		Result       struct {
			Txid string `json:"txid"`
		} `json:"result"`
		ExecutionSecs float64 `json:"execution_secs"`
		Method        string  `json:"method"`
		Params        struct {
			Fromaddress string `json:"fromaddress"`
			Amounts     []struct {
				Address string  `json:"address"`
				Amount  float64 `json:"amount"`
			} `json:"amounts"`
			Minconf int     `json:"minconf"`
			Fee     float64 `json:"fee"`
		} `json:"params"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZGetOperationStatus(params APIParams) (ZGetOperationStatus, error) {
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_getoperationstatus`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zGetoperationstatus ZGetOperationStatus

	zGetoperationstatusJSON := appName.APICall(query)
	if zGetoperationstatusJSON == "EMPTY RPC INFO" {
		return zGetoperationstatus, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zGetoperationstatusJSON)

	var result APIResult

	json.Unmarshal([]byte(zGetoperationstatusJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zGetoperationstatusJSON), &zGetoperationstatus)
		return zGetoperationstatus, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zGetoperationstatusJSON), &zGetoperationstatus)
	return zGetoperationstatus, nil
}

type ZGetTotalBalance struct {
	Result struct {
		Transparent float64 `json:"transparent"`
		Interest    float64 `json:"interest"`
		Private     float64 `json:"private"`
		Total       float64 `json:"total"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZGetTotalBalance(params APIParams) (ZGetTotalBalance, error) {
	if params[0] == nil {
		params[0] = 1
	}
	if params[1] == nil {
		params[1] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_gettotalbalance`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zGettotalbalance ZGetTotalBalance

	zGettotalbalanceJSON := appName.APICall(query)
	if zGettotalbalanceJSON == "EMPTY RPC INFO" {
		return zGettotalbalance, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zGettotalbalanceJSON)

	var result APIResult

	json.Unmarshal([]byte(zGettotalbalanceJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zGettotalbalanceJSON), &zGettotalbalance)
		return zGettotalbalance, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zGettotalbalanceJSON), &zGettotalbalance)
	return zGettotalbalance, nil
}

type ZImportKey struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) ZImportKey(params APIParams) (ZImportKey, error) {
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = "whenkeyisnew"
		}
	}

	if len(params) == 3 {
		if params[2] == nil {
			params[2] = 0
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_importkey`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zImportkey ZImportKey

	zImportkeyJSON := appName.APICall(query)
	if zImportkeyJSON == "EMPTY RPC INFO" {
		return zImportkey, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zImportkeyJSON)

	var result APIResult

	json.Unmarshal([]byte(zImportkeyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zImportkeyJSON), &zImportkey)
		return zImportkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zImportkeyJSON), &zImportkey)
	return zImportkey, nil
}

type ZImportViewingKey struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) ZImportViewingKey(params APIParams) (ZImportViewingKey, error) {
	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = "whenkeyisnew"
		}
	}

	if len(params) == 3 {
		if params[2] == nil {
			params[2] = 0
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_importviewingkey`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zImportviewingkey ZImportViewingKey

	zImportviewingkeyJSON := appName.APICall(query)
	if zImportviewingkeyJSON == "EMPTY RPC INFO" {
		return zImportviewingkey, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zImportviewingkeyJSON)

	var result APIResult

	json.Unmarshal([]byte(zImportviewingkeyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zImportviewingkeyJSON), &zImportviewingkey)
		return zImportviewingkey, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zImportviewingkeyJSON), &zImportviewingkey)
	return zImportviewingkey, nil
}

type ZImportWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZImportWallet(wltfile string) (ZImportWallet, error) {
	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_importwallet`,
		Params: `["` + wltfile + `"]`,
	}
	//fmt.Println(query)

	var zImportwallet ZImportWallet

	zImportwalletJSON := appName.APICall(query)
	if zImportwalletJSON == "EMPTY RPC INFO" {
		return zImportwallet, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zImportwalletJSON)

	var result APIResult

	json.Unmarshal([]byte(zImportwalletJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zImportwalletJSON), &zImportwallet)
		return zImportwallet, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zImportwalletJSON), &zImportwallet)
	return zImportwallet, nil
}

type ZListAddresses struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) ZListAddresses(params APIParams) (ZListAddresses, error) {
	if params[0] == nil {
		params[0] = false
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_listaddresses`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zListaddresses ZListAddresses

	zListaddressesJSON := appName.APICall(query)
	if zListaddressesJSON == "EMPTY RPC INFO" {
		return zListaddresses, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zListaddressesJSON)

	var result APIResult

	json.Unmarshal([]byte(zListaddressesJSON), &result)

	if result.Result == nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zListaddressesJSON), &zListaddresses)
		return zListaddresses, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zListaddressesJSON), &zListaddresses)
	return zListaddresses, nil
}

type ZListOperationIDs struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) ZListOperationIDs(params APIParams) (ZListOperationIDs, error) {
	var paramsJSON string

	if params[0] == "" || params[0] == nil {
		paramsJSON = `[]`
		//fmt.Println(paramsJSON)
	} else {
		paramsBytes, _ := json.Marshal(params)
		paramsJSON = string(paramsBytes)
		//fmt.Println(paramsJSON)
	}

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_listoperationids`,
		Params: paramsJSON,
	}
	//fmt.Println(query)

	var zListoperationids ZListOperationIDs

	zListoperationidsJSON := appName.APICall(query)
	if zListoperationidsJSON == "EMPTY RPC INFO" {
		return zListoperationids, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zListoperationidsJSON)

	var result APIResult

	json.Unmarshal([]byte(zListoperationidsJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zListoperationidsJSON), &zListoperationids)
		return zListoperationids, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zListoperationidsJSON), &zListoperationids)
	return zListoperationids, nil
}

type ZListReceivedByAddress struct {
	Result []struct {
		Txid   string  `json:"txid"`
		Amount float64 `json:"amount"`
		Memo   string  `json:"memo"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZListReceivedByAddress(params APIParams) (ZListReceivedByAddress, error) {
	if len(params) == 2 {
		if params[1] == nil {
			params[1] = 1
		}
	}
	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_listreceivedbyaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zListreceivedbyaddress ZListReceivedByAddress

	zListreceivedbyaddressJSON := appName.APICall(query)
	if zListreceivedbyaddressJSON == "EMPTY RPC INFO" {
		return zListreceivedbyaddress, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zListreceivedbyaddressJSON)

	var result APIResult

	json.Unmarshal([]byte(zListreceivedbyaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zListreceivedbyaddressJSON), &zListreceivedbyaddress)
		return zListreceivedbyaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zListreceivedbyaddressJSON), &zListreceivedbyaddress)
	return zListreceivedbyaddress, nil
}

type ZListUnspent struct {
	Result []struct {
		Txid          string  `json:"txid"`
		Jsindex       int     `json:"jsindex"`
		Jsoutindex    int     `json:"jsoutindex"`
		Outindex      int     `json:"outindex"`
		Confirmations int     `json:"confirmations"`
		Spendable     bool    `json:"spendable"`
		Address       string  `json:"address"`
		Amount        float64 `json:"amount"`
		Memo          string  `json:"memo"`
		Change        bool    `json:"change"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZListUnspent(params APIParams) (ZListUnspent, error) {
	if len(params) >= 1 {
		if params[0] == nil {
			params[0] = 1
		}
	}

	if len(params) >= 2 {
		if params[1] == nil {
			params[1] = 9999999
		}
	}

	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = false
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_listunspent`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zListunspent ZListUnspent

	zListunspentJSON := appName.APICall(query)
	if zListunspentJSON == "EMPTY RPC INFO" {
		return zListunspent, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zListunspentJSON)

	var result APIResult

	json.Unmarshal([]byte(zListunspentJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zListunspentJSON), &zListunspent)
		return zListunspent, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zListunspentJSON), &zListunspent)
	return zListunspent, nil
}

type ZMergeToAddress struct {
	Result struct {
		RemainingUTXOs            int     `json:"remainingUTXOs"`
		RemainingTransparentValue float64 `json:"remainingTransparentValue"`
		RemainingNotes            int     `json:"remainingNotes"`
		RemainingShieldedValue    float64 `json:"remainingShieldedValue"`
		MergingUTXOs              int     `json:"mergingUTXOs"`
		MergingTransparentValue   float64 `json:"mergingTransparentValue"`
		MergingNotes              int     `json:"mergingNotes"`
		MergingShieldedValue      float64 `json:"mergingShieldedValue"`
		Opid                      string  `json:"opid"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZMergeToAddress(params APIParams) (ZMergeToAddress, error) {

	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 0.0001
		}
	}

	if len(params) >= 4 {
		if params[3] == nil {
			params[3] = 50
		}
	}

	if len(params) >= 5 {
		if params[4] == nil {
			params[4] = 10
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_mergetoaddress`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zMergetoaddress ZMergeToAddress

	zMergetoaddressJSON := appName.APICall(query)
	if zMergetoaddressJSON == "EMPTY RPC INFO" {
		return zMergetoaddress, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zMergetoaddressJSON)

	var result APIResult

	json.Unmarshal([]byte(zMergetoaddressJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zMergetoaddressJSON), &zMergetoaddress)
		return zMergetoaddress, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zMergetoaddressJSON), &zMergetoaddress)
	return zMergetoaddress, nil
}

type ZSendMany struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZSendMany(params APIParams) (ZSendMany, error) {
	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 1
		}
	}

	if len(params) >= 4 {
		if params[3] == nil {
			params[3] = 0.0001
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_sendmany`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zSendmany ZSendMany

	zSendmanyJSON := appName.APICall(query)
	if zSendmanyJSON == "EMPTY RPC INFO" {
		return zSendmany, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zSendmanyJSON)

	var result APIResult

	json.Unmarshal([]byte(zSendmanyJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zSendmanyJSON), &zSendmany)
		return zSendmany, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zSendmanyJSON), &zSendmany)
	return zSendmany, nil
}

type ZShieldCoinbase struct {
	Result struct {
		RemainingUTXOs int     `json:"remainingUTXOs"`
		RemainingValue float64 `json:"remainingValue"`
		ShieldingUTXOs int     `json:"shieldingUTXOs"`
		ShieldingValue float64 `json:"shieldingValue"`
		Opid           string  `json:"opid"`
	} `json:"result"`
	Error Error  `json:"error"`
	ID    string `json:"id"`
}

func (appName AppType) ZShieldCoinbase(params APIParams) (ZShieldCoinbase, error) {
	if len(params) >= 3 {
		if params[2] == nil {
			params[2] = 0.0001
		}
	}

	if len(params) >= 4 {
		if params[3] == nil {
			params[3] = 50
		}
	}

	paramsJSON, _ := json.Marshal(params)
	//fmt.Println(string(paramsJSON))

	query := new(APIQuery)
	query = &APIQuery{
		Method: `z_shieldcoinbase`,
		Params: string(paramsJSON),
	}
	//fmt.Println(query)

	var zShieldcoinbase ZShieldCoinbase

	zShieldcoinbaseJSON := appName.APICall(query)
	if zShieldcoinbaseJSON == "EMPTY RPC INFO" {
		return zShieldcoinbase, errors.New("EMPTY RPC INFO")
	}
	//fmt.Println(zShieldcoinbaseJSON)

	var result APIResult

	json.Unmarshal([]byte(zShieldcoinbaseJSON), &result)

	if result.Error != nil {
		answerError, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zShieldcoinbaseJSON), &zShieldcoinbase)
		return zShieldcoinbase, errors.New(string(answerError))
	}

	json.Unmarshal([]byte(zShieldcoinbaseJSON), &zShieldcoinbase)
	return zShieldcoinbase, nil
}
