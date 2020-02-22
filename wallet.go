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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `listunspent`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var listunspent ListUnspent

	listunspentJson := appName.APICall(query)
	if listunspentJson == "EMPTY RPC INFO!" {
		return listunspent, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listunspentJson)

	var result APIResult

	json.Unmarshal([]byte(listunspentJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listunspentJson), &listunspent)
		return listunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listunspentJson), &listunspent)
	return listunspent, nil
}

type AddMultiSigAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) AddMultiSigAddress(params APIParams) (AddMultiSigAddress, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `addmultisigaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var addmultisigaddress AddMultiSigAddress

	addmultisigaddressJson := appName.APICall(query)
	if addmultisigaddressJson == "EMPTY RPC INFO!" {
		return addmultisigaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(addmultisigaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(addmultisigaddressJson), &addmultisigaddress)
		return addmultisigaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(addmultisigaddressJson), &addmultisigaddress)
	return addmultisigaddress, nil
}

type BackupWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) BackupWallet(bkpwlt string) (BackupWallet, error) {
	query := APIQuery{
		Method: `backupwallet`,
		Params: `["` + bkpwlt + `"]`,
	}
	//fmt.Println(query)

	var backupwallet BackupWallet

	backupwalletJson := appName.APICall(query)
	if backupwalletJson == "EMPTY RPC INFO!" {
		return backupwallet, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(backupwalletJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(backupwalletJson), &backupwallet)
		return backupwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(backupwalletJson), &backupwallet)
	return backupwallet, nil
}

type DumpPrivKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) DumpPrivKey(taddr string) (DumpPrivKey, error) {
	query := APIQuery{
		Method: `dumpprivkey`,
		Params: `["` + taddr + `"]`,
	}
	//fmt.Println(query)

	var dumpprivkey DumpPrivKey

	dumpprivkeyJson := appName.APICall(query)
	if dumpprivkeyJson == "EMPTY RPC INFO!" {
		return dumpprivkey, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dumpprivkeyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dumpprivkeyJson), &dumpprivkey)
		return dumpprivkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dumpprivkeyJson), &dumpprivkey)
	return dumpprivkey, nil
}

type DumpWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) DumpWallet(dmpwlt string) (DumpWallet, error) {
	query := APIQuery{
		Method: `dumpwallet`,
		Params: `["` + dmpwlt + `"]`,
	}
	//fmt.Println(query)

	var dumpwallet DumpWallet

	dumpwalletJson := appName.APICall(query)
	if dumpwalletJson == "EMPTY RPC INFO!" {
		return dumpwallet, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(dumpwalletJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(dumpwalletJson), &dumpwallet)
		return dumpwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(dumpwalletJson), &dumpwallet)
	return dumpwallet, nil
}

type EncryptWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) EncryptWallet(phrase string) (EncryptWallet, error) {
	query := APIQuery{
		Method: `encryptwallet`,
		Params: `["` + phrase + `"]`,
	}
	//fmt.Println(query)

	var encryptwallet EncryptWallet

	encryptwalletJson := appName.APICall(query)
	if encryptwalletJson == "EMPTY RPC INFO!" {
		return encryptwallet, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(encryptwalletJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(encryptwalletJson), &encryptwallet)
		return encryptwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(encryptwalletJson), &encryptwallet)
	return encryptwallet, nil
}

type GetAccount struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetAccount(taddr string) (GetAccount, error) {
	query := APIQuery{
		Method: `getaccount`,
		Params: `["` + taddr + `"]`,
	}
	//fmt.Println(query)

	var getaccount GetAccount

	getaccountJson := appName.APICall(query)
	if getaccountJson == "EMPTY RPC INFO!" {
		return getaccount, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getaccountJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaccountJson), &getaccount)
		return getaccount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getaccountJson), &getaccount)
	return getaccount, nil
}

type GetAccountAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetAccountAddress(actname string) (GetAccountAddress, error) {
	query := APIQuery{
		Method: `getaccountaddress`,
		Params: `["` + actname + `"]`,
	}
	//fmt.Println(query)

	var getaccountaddress GetAccountAddress

	getaccountaddressJson := appName.APICall(query)
	if getaccountaddressJson == "EMPTY RPC INFO!" {
		return getaccountaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getaccountaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaccountaddressJson), &getaccountaddress)
		return getaccountaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getaccountaddressJson), &getaccountaddress)
	return getaccountaddress, nil
}

type GetAddressesByAccount struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) GetAddressesByAccount(actname string) (GetAddressesByAccount, error) {
	query := APIQuery{
		Method: `getaddressesbyaccount`,
		Params: `["` + actname + `"]`,
	}
	//fmt.Println(query)

	var getaddressesbyaccount GetAddressesByAccount

	getaddressesbyaccountJson := appName.APICall(query)
	if getaddressesbyaccountJson == "EMPTY RPC INFO!" {
		return getaddressesbyaccount, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getaddressesbyaccountJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getaddressesbyaccountJson), &getaddressesbyaccount)
		return getaddressesbyaccount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getaddressesbyaccountJson), &getaddressesbyaccount)
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getbalance`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getbalance GetBalance

	getbalanceJson := appName.APICall(query)
	if getbalanceJson == "EMPTY RPC INFO!" {
		return getbalance, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getbalanceJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getbalanceJson), &getbalance)
		return getbalance, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getbalanceJson), &getbalance)
	return getbalance, nil
}

type GetNewAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetNewAddress() (GetNewAddress, error) {
	query := APIQuery{
		Method: `getnewaddress`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getnewaddress GetNewAddress

	getnewaddressJson := appName.APICall(query)
	if getnewaddressJson == "EMPTY RPC INFO!" {
		return getnewaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getnewaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getnewaddressJson), &getnewaddress)
		return getnewaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getnewaddressJson), &getnewaddress)
	return getnewaddress, nil
}

type GetRawChangeAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) GetRawChangeAddress() (GetRawChangeAddress, error) {
	query := APIQuery{
		Method: `getrawchangeaddress`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getrawchangeaddress GetRawChangeAddress

	getrawchangeaddressJson := appName.APICall(query)
	if getrawchangeaddressJson == "EMPTY RPC INFO!" {
		return getrawchangeaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getrawchangeaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getrawchangeaddressJson), &getrawchangeaddress)
		return getrawchangeaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getrawchangeaddressJson), &getrawchangeaddress)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getreceivedbyaccount`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getreceivedbyaccount GetReceivedByAccount

	getreceivedbyaccountJson := appName.APICall(query)
	if getreceivedbyaccountJson == "EMPTY RPC INFO!" {
		return getreceivedbyaccount, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getreceivedbyaccountJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getreceivedbyaccountJson), &getreceivedbyaccount)
		return getreceivedbyaccount, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getreceivedbyaccountJson), &getreceivedbyaccount)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `getreceivedbyaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var getreceivedbyaddress GetReceivedByAddress

	getreceivedbyaddressJson := appName.APICall(query)
	if getreceivedbyaddressJson == "EMPTY RPC INFO!" {
		return getreceivedbyaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getreceivedbyaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getreceivedbyaddressJson), &getreceivedbyaddress)
		return getreceivedbyaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getreceivedbyaddressJson), &getreceivedbyaddress)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `gettransaction`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var gettransaction GetTransaction

	gettransactionJson := appName.APICall(query)
	if gettransactionJson == "EMPTY RPC INFO!" {
		return gettransaction, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(gettransactionJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(gettransactionJson), &gettransaction)
		return gettransaction, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(gettransactionJson), &gettransaction)
	return gettransaction, nil
}

type GetUnconfirmedBalance struct {
	Result float64 `json:"result"`
	Error  Error   `json:"error"`
	ID     string  `json:"id"`
}

func (appName AppType) GetUnconfirmedBalance() (GetUnconfirmedBalance, error) {
	query := APIQuery{
		Method: `getunconfirmedbalance`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getunconfirmedbalance GetUnconfirmedBalance

	getunconfirmedbalanceJson := appName.APICall(query)
	if getunconfirmedbalanceJson == "EMPTY RPC INFO!" {
		return getunconfirmedbalance, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getunconfirmedbalanceJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getunconfirmedbalanceJson), &getunconfirmedbalance)
		return getunconfirmedbalance, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getunconfirmedbalanceJson), &getunconfirmedbalance)
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
	query := APIQuery{
		Method: `getwalletinfo`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var getwalletinfo GetWalletInfo

	getwalletinfoJson := appName.APICall(query)
	if getwalletinfoJson == "EMPTY RPC INFO!" {
		return getwalletinfo, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(getwalletinfoJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(getwalletinfoJson), &getwalletinfo)
		return getwalletinfo, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(getwalletinfoJson), &getwalletinfo)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `importaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var importaddress ImportAddress

	importaddressJson := appName.APICall(query)
	if importaddressJson == "EMPTY RPC INFO!" {
		return importaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(importaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importaddressJson), &importaddress)
		return importaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(importaddressJson), &importaddress)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `importprivkey`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var importprivkey ImportPrivKey

	importprivkeyJson := appName.APICall(query)
	if importprivkeyJson == "EMPTY RPC INFO!" {
		return importprivkey, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(importprivkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importprivkeyJson), &importprivkey)
		return importprivkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(importprivkeyJson), &importprivkey)
	return importprivkey, nil
}

type ImportWallet struct {
	Result interface{} `json:"result"`
	Error  Error       `json:"error"`
	ID     string      `json:"id"`
}

func (appName AppType) ImportWallet(wltpth string) (ImportWallet, error) {
	query := APIQuery{
		Method: `importwallet`,
		Params: `["` + wltpth + `"]`,
	}
	//fmt.Println(query)

	var importwallet ImportWallet

	importwalletJson := appName.APICall(query)
	if importwalletJson == "EMPTY RPC INFO!" {
		return importwallet, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(importwalletJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(importwalletJson), &importwallet)
		return importwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(importwalletJson), &importwallet)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `keypoolrefill`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var keypoolrefill KeyPoolRefill

	keypoolrefillJson := appName.APICall(query)
	if keypoolrefillJson == "EMPTY RPC INFO!" {
		return keypoolrefill, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(keypoolrefillJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(keypoolrefillJson), &keypoolrefill)
		return keypoolrefill, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(keypoolrefillJson), &keypoolrefill)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `listaccounts`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var listaccounts ListAccounts

	listaccountsJson := appName.APICall(query)
	if listaccountsJson == "EMPTY RPC INFO!" {
		return listaccounts, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listaccountsJson)

	var result APIResult

	json.Unmarshal([]byte(listaccountsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listaccountsJson), &listaccounts)
		return listaccounts, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listaccountsJson), &listaccounts)
	return listaccounts, nil
}

type ListAddressGroupings struct {
	Result [][][]interface{} `json:"result"`
	Error  Error             `json:"error"`
	ID     string            `json:"id"`
}

func (appName AppType) ListAddressGroupings() (ListAddressGroupings, error) {
	query := APIQuery{
		Method: `listaddressgroupings`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var listaddressgroupings ListAddressGroupings

	listaddressgroupingsJson := appName.APICall(query)
	if listaddressgroupingsJson == "EMPTY RPC INFO!" {
		return listaddressgroupings, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listaddressgroupingsJson)

	var result APIResult

	json.Unmarshal([]byte(listaddressgroupingsJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listaddressgroupingsJson), &listaddressgroupings)
		return listaddressgroupings, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listaddressgroupingsJson), &listaddressgroupings)
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
	query := APIQuery{
		Method: `listlockunspent`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var listlockunspent ListLockUnspent

	listlockunspentJson := appName.APICall(query)
	if listlockunspentJson == "EMPTY RPC INFO!" {
		return listlockunspent, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listlockunspentJson)

	var result APIResult

	json.Unmarshal([]byte(listlockunspentJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listlockunspentJson), &listlockunspent)
		return listlockunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listlockunspentJson), &listlockunspent)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `listreceivedbyaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var listreceivedbyaddress ListReceivedByAddress

	listreceivedbyaddressJson := appName.APICall(query)
	if listreceivedbyaddressJson == "EMPTY RPC INFO!" {
		return listreceivedbyaddress, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listreceivedbyaddressJson)

	var result APIResult

	json.Unmarshal([]byte(listreceivedbyaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listreceivedbyaddressJson), &listreceivedbyaddress)
		return listreceivedbyaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listreceivedbyaddressJson), &listreceivedbyaddress)
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `listsinceblock`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var listsinceblock ListSinceBlock

	listsinceblockJson := appName.APICall(query)
	if listsinceblockJson == "EMPTY RPC INFO!" {
		return listsinceblock, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(listsinceblockJson)

	var result APIResult

	json.Unmarshal([]byte(listsinceblockJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listsinceblockJson), &listsinceblock)
		return listsinceblock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listsinceblockJson), &listsinceblock)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `listtransactions`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var listtransactions ListTransactions

	listtransactionsJson := appName.APICall(query)
	if listtransactionsJson == "EMPTY RPC INFO!" {
		return listtransactions, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(listtransactionsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
		return listtransactions, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(listtransactionsJson), &listtransactions)
	return listtransactions, nil
}

type LockUnspent struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) LockUnspent(params APIParams) (LockUnspent, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `lockunspent`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var lockunspent LockUnspent

	lockunspentJson := appName.APICall(query)
	if lockunspentJson == "EMPTY RPC INFO!" {
		return lockunspent, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(lockunspentJson)

	var result APIResult

	json.Unmarshal([]byte(lockunspentJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(lockunspentJson), &lockunspent)
		return lockunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(lockunspentJson), &lockunspent)
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `sendfrom`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var sendfrom SendFrom

	sendfromJson := appName.APICall(query)
	if sendfromJson == "EMPTY RPC INFO!" {
		return sendfrom, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(sendfromJson)

	var result APIResult

	json.Unmarshal([]byte(sendfromJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendfromJson), &sendfrom)
		return sendfrom, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendfromJson), &sendfrom)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `sendmany`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var sendmany SendMany

	sendmanyJson := appName.APICall(query)
	if sendmanyJson == "EMPTY RPC INFO!" {
		return sendmany, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(sendmanyJson)

	var result APIResult

	json.Unmarshal([]byte(sendmanyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendmanyJson), &sendmany)
		return sendmany, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendmanyJson), &sendmany)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `sendtoaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var sendtoaddress SendToAddress

	sendtoaddressJson := appName.APICall(query)
	if sendtoaddressJson == "EMPTY RPC INFO!" {
		return sendtoaddress, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(sendtoaddressJson)

	var result APIResult

	json.Unmarshal([]byte(sendtoaddressJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(sendtoaddressJson), &sendtoaddress)
		return sendtoaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(sendtoaddressJson), &sendtoaddress)
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
	query := APIQuery{
		Method: `setpubkey`,
		Params: `["` + pubkey + `"]`,
	}
	//fmt.Println(query)

	var setpubkey SetPubKey

	setpubkeyJson := appName.APICall(query)
	if setpubkeyJson == "EMPTY RPC INFO!" {
		return setpubkey, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(setpubkeyJson)

	var result APIResult

	json.Unmarshal([]byte(setpubkeyJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(setpubkeyJson), &setpubkey)
		return setpubkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(setpubkeyJson), &setpubkey)
	return setpubkey, nil
}

type SetTxFee struct {
	Result bool   `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SetTxFee(amount float64) (SetTxFee, error) {
	query := APIQuery{
		Method: `settxfee`,
		Params: `[` + strconv.FormatFloat(amount, 'f', 8, 64) + `]`,
	}
	//fmt.Println(query)

	var settxfee SetTxFee

	settxfeeJson := appName.APICall(query)
	if settxfeeJson == "EMPTY RPC INFO!" {
		return settxfee, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(settxfeeJson)

	var result APIResult

	json.Unmarshal([]byte(settxfeeJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(settxfeeJson), &settxfee)
		return settxfee, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(settxfeeJson), &settxfee)
	return settxfee, nil
}

type SignMessage struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) SignMessage(params APIParams) (SignMessage, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `signmessage`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var signmessage SignMessage

	signmessageJson := appName.APICall(query)
	if signmessageJson == "EMPTY RPC INFO!" {
		return signmessage, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(signmessageJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(signmessageJson), &signmessage)
		return signmessage, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(signmessageJson), &signmessage)
	return signmessage, nil
}

type WalletLock struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletLock() (WalletLock, error) {
	query := APIQuery{
		Method: `walletlock`,
		Params: `[]`,
	}
	//fmt.Println(query)

	var walletlock WalletLock

	walletlockJson := appName.APICall(query)
	if walletlockJson == "EMPTY RPC INFO!" {
		return walletlock, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(walletlockJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletlockJson), &walletlock)
		return walletlock, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(walletlockJson), &walletlock)
	return walletlock, nil
}

type WalletPassPhrase struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletPassPhrase(params APIParams) (WalletPassPhrase, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `walletpassphrase`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var walletpassphrase WalletPassPhrase

	walletpassphraseJson := appName.APICall(query)
	if walletpassphraseJson == "EMPTY RPC INFO!" {
		return walletpassphrase, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(walletpassphraseJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletpassphraseJson), &walletpassphrase)
		return walletpassphrase, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(walletpassphraseJson), &walletpassphrase)
	return walletpassphrase, nil
}

type WalletPassPhrasechangeChange struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) WalletPassPhrasechangeChange(params APIParams) (WalletPassPhrasechangeChange, error) {
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `walletpassphrasechange`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var walletpassphrasechange WalletPassPhrasechangeChange

	walletpassphrasechangeJson := appName.APICall(query)
	if listtransactionsJson walletpassphrasechangeJson "EMPTY RPC INFO!" {
		return walletpassphrasechange, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(walletpassphrasechangeJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(walletpassphrasechangeJson), &walletpassphrasechange)
		return walletpassphrasechange, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(walletpassphrasechangeJson), &walletpassphrasechange)
	return walletpassphrasechange, nil
}

type ZExportKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportKey(zaddr string) (ZExportKey, error) {
	query := APIQuery{
		Method: `z_exportkey`,
		Params: `["` + zaddr + `"]`,
	}
	//fmt.Println(query)

	var zexportkey ZExportKey

	zexportkeyJson := appName.APICall(query)
	if zexportkeyJson == "EMPTY RPC INFO!" {
		return zexportkey, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(zexportkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zexportkeyJson), &zexportkey)
		return zexportkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(zexportkeyJson), &zexportkey)
	return zexportkey, nil
}

type ZExportViewingKey struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportViewingKey(zaddr string) (ZExportViewingKey, error) {
	query := APIQuery{
		Method: `z_exportviewingkey`,
		Params: `["` + zaddr + `"]`,
	}
	//fmt.Println(query)

	var zexportviewingkey ZExportViewingKey

	zexportviewingkeyJson := appName.APICall(query)
	if zexportviewingkeyJson == "EMPTY RPC INFO!" {
		return zexportviewingkey, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(zexportviewingkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zexportviewingkeyJson), &zexportviewingkey)
		return zexportviewingkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(zexportviewingkeyJson), &zexportviewingkey)
	return zexportviewingkey, nil
}

type ZExportWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZExportWallet(wltfile string) (ZExportWallet, error) {
	query := APIQuery{
		Method: `z_exportwallet`,
		Params: `["` + wltfile + `"]`,
	}
	//fmt.Println(query)

	var z_exportwallet ZExportWallet

	z_exportwalletJson := appName.APICall(query)
	if z_exportwalletJson == "EMPTY RPC INFO!" {
		return z_exportwallet, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_exportwalletJson)

	var result APIResult

	json.Unmarshal([]byte(z_exportwalletJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_exportwalletJson), &z_exportwallet)
		return z_exportwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_exportwalletJson), &z_exportwallet)
	return z_exportwallet, nil
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_getbalance`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_getbalance ZGetBalance

	z_getbalanceJson := appName.APICall(query)
	if z_getbalanceJson == "EMPTY RPC INFO!" {
		return z_getbalance, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_getbalanceJson)

	var result APIResult

	json.Unmarshal([]byte(z_getbalanceJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_getbalanceJson), &z_getbalance)
		return z_getbalance, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_getbalanceJson), &z_getbalance)
	return z_getbalance, nil
}

type ZGetNewAddress struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZGetNewAddress(tp string) (ZGetNewAddress, error) {
	query := APIQuery{
		Method: `z_getnewaddress`,
		Params: `["` + tp + `"]`,
	}
	//fmt.Println(query)

	var zgetnewaddress ZGetNewAddress

	zgetnewaddressJson := appName.APICall(query)
	if zgetnewaddressJson == "EMPTY RPC INFO!" {
		return zgetnewaddress, errors.New("EMPTY RPC INFO!")
	}

	var result APIResult

	json.Unmarshal([]byte(zgetnewaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(zgetnewaddressJson), &zgetnewaddress)
		return zgetnewaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(zgetnewaddressJson), &zgetnewaddress)
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_getoperationresult`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_getoperationresult ZGetOperationResult

	z_getoperationresultJson := appName.APICall(query)
	if z_getoperationresultJson == "EMPTY RPC INFO!" {
		return z_getoperationresult, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_getoperationresultJson)

	var result APIResult

	json.Unmarshal([]byte(z_getoperationresultJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_getoperationresultJson), &z_getoperationresult)
		return z_getoperationresult, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_getoperationresultJson), &z_getoperationresult)
	return z_getoperationresult, nil
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_getoperationstatus`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_getoperationstatus ZGetOperationStatus

	z_getoperationstatusJson := appName.APICall(query)
	if z_getoperationstatusJson == "EMPTY RPC INFO!" {
		return z_getoperationstatus, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_getoperationstatusJson)

	var result APIResult

	json.Unmarshal([]byte(z_getoperationstatusJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_getoperationstatusJson), &z_getoperationstatus)
		return z_getoperationstatus, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_getoperationstatusJson), &z_getoperationstatus)
	return z_getoperationstatus, nil
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_gettotalbalance`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_gettotalbalance ZGetTotalBalance

	z_gettotalbalanceJson := appName.APICall(query)
	if z_gettotalbalanceJson == "EMPTY RPC INFO!" {
		return z_gettotalbalance, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_gettotalbalanceJson)

	var result APIResult

	json.Unmarshal([]byte(z_gettotalbalanceJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_gettotalbalanceJson), &z_gettotalbalance)
		return z_gettotalbalance, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_gettotalbalanceJson), &z_gettotalbalance)
	return z_gettotalbalance, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_importkey`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_importkey ZImportKey

	z_importkeyJson := appName.APICall(query)
	if z_importkeyJson == "EMPTY RPC INFO!" {
		return z_importkey, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_importkeyJson)

	var result APIResult

	json.Unmarshal([]byte(z_importkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_importkeyJson), &z_importkey)
		return z_importkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_importkeyJson), &z_importkey)
	return z_importkey, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_importviewingkey`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_importviewingkey ZImportViewingKey

	z_importviewingkeyJson := appName.APICall(query)
	if z_importviewingkeyJson == "EMPTY RPC INFO!" {
		return z_importviewingkey, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_importviewingkeyJson)

	var result APIResult

	json.Unmarshal([]byte(z_importviewingkeyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_importviewingkeyJson), &z_importviewingkey)
		return z_importviewingkey, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_importviewingkeyJson), &z_importviewingkey)
	return z_importviewingkey, nil
}

type ZImportWallet struct {
	Result string `json:"result"`
	Error  Error  `json:"error"`
	ID     string `json:"id"`
}

func (appName AppType) ZImportWallet(wltfile string) (ZImportWallet, error) {
	query := APIQuery{
		Method: `z_importwallet`,
		Params: `["` + wltfile + `"]`,
	}
	//fmt.Println(query)

	var z_importwallet ZImportWallet

	z_importwalletJson := appName.APICall(query)
	if z_importwalletJson == "EMPTY RPC INFO!" {
		return z_importwallet, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_importwalletJson)

	var result APIResult

	json.Unmarshal([]byte(z_importwalletJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_importwalletJson), &z_importwallet)
		return z_importwallet, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_importwalletJson), &z_importwallet)
	return z_importwallet, nil
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_listaddresses`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_listaddresses ZListAddresses

	z_listaddressesJson := appName.APICall(query)
	if z_listaddressesJson == "EMPTY RPC INFO!" {
		return z_listaddresses, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_listaddressesJson)

	var result APIResult

	json.Unmarshal([]byte(z_listaddressesJson), &result)

	if result.Result == nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listaddressesJson), &z_listaddresses)
		return z_listaddresses, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listaddressesJson), &z_listaddresses)
	return z_listaddresses, nil
}

type ZListOperationIDs struct {
	Result []string `json:"result"`
	Error  Error    `json:"error"`
	ID     string   `json:"id"`
}

func (appName AppType) ZListOperationIDs(params APIParams) (ZListOperationIDs, error) {
	var params_json string

	if params[0] == "" || params[0] == nil {
		params_json = `[]`
		//fmt.Println(params_json)
	} else {
		params_bytes, _ := json.Marshal(params)
		params_json = string(params_bytes)
		//fmt.Println(params_json)
	}

	query := APIQuery{
		Method: `z_listoperationids`,
		Params: params_json,
	}
	//fmt.Println(query)

	var z_listoperationids ZListOperationIDs

	z_listoperationidsJson := appName.APICall(query)
	if z_listoperationidsJson == "EMPTY RPC INFO!" {
		return z_listoperationids, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_listoperationidsJson)

	var result APIResult

	json.Unmarshal([]byte(z_listoperationidsJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listoperationidsJson), &z_listoperationids)
		return z_listoperationids, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listoperationidsJson), &z_listoperationids)
	return z_listoperationids, nil
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
	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_listreceivedbyaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_listreceivedbyaddress ZListReceivedByAddress

	z_listreceivedbyaddressJson := appName.APICall(query)
	if listtransactionsJson z_listreceivedbyaddressJson "EMPTY RPC INFO!" {
		return z_listreceivedbyaddress, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_listreceivedbyaddressJson)

	var result APIResult

	json.Unmarshal([]byte(z_listreceivedbyaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listreceivedbyaddressJson), &z_listreceivedbyaddress)
		return z_listreceivedbyaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listreceivedbyaddressJson), &z_listreceivedbyaddress)
	return z_listreceivedbyaddress, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_listunspent`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_listunspent ZListUnspent

	z_listunspentJson := appName.APICall(query)
	if z_listunspentJson == "EMPTY RPC INFO!" {
		return z_listunspent, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_listunspentJson)

	var result APIResult

	json.Unmarshal([]byte(z_listunspentJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_listunspentJson), &z_listunspent)
		return z_listunspent, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_listunspentJson), &z_listunspent)
	return z_listunspent, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_mergetoaddress`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_mergetoaddress ZMergeToAddress

	z_mergetoaddressJson := appName.APICall(query)
	if z_mergetoaddressJson == "EMPTY RPC INFO!" {
		return z_mergetoaddress, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_mergetoaddressJson)

	var result APIResult

	json.Unmarshal([]byte(z_mergetoaddressJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_mergetoaddressJson), &z_mergetoaddress)
		return z_mergetoaddress, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_mergetoaddressJson), &z_mergetoaddress)
	return z_mergetoaddress, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_sendmany`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_sendmany ZSendMany

	z_sendmanyJson := appName.APICall(query)
	if z_sendmanyJson == "EMPTY RPC INFO!" {
		return z_sendmany, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_sendmanyJson)

	var result APIResult

	json.Unmarshal([]byte(z_sendmanyJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_sendmanyJson), &z_sendmany)
		return z_sendmany, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_sendmanyJson), &z_sendmany)
	return z_sendmany, nil
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

	params_json, _ := json.Marshal(params)
	//fmt.Println(string(params_json))

	query := APIQuery{
		Method: `z_shieldcoinbase`,
		Params: string(params_json),
	}
	//fmt.Println(query)

	var z_shieldcoinbase ZShieldCoinbase

	z_shieldcoinbaseJson := appName.APICall(query)
	if z_shieldcoinbaseJson == "EMPTY RPC INFO!" {
		return z_shieldcoinbase, errors.New("EMPTY RPC INFO!")
	}
	//fmt.Println(z_shieldcoinbaseJson)

	var result APIResult

	json.Unmarshal([]byte(z_shieldcoinbaseJson), &result)

	if result.Error != nil {
		answer_error, err := json.Marshal(result.Error)
		if err != nil {
		}
		json.Unmarshal([]byte(z_shieldcoinbaseJson), &z_shieldcoinbase)
		return z_shieldcoinbase, errors.New(string(answer_error))
	}

	json.Unmarshal([]byte(z_shieldcoinbaseJson), &z_shieldcoinbase)
	return z_shieldcoinbase, nil
}
