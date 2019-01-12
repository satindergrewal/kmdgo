package kmdgo

import (
	"errors"
	"encoding/json"
	"strconv"
)

type GetBlock struct {
	Result struct {
		Hash				string	`json:"hash"`
		Confirmations		int		`json:"confirmations"`
		Rawconfirmations 	int		`json:"rawconfirmations"`
		Size 				int		`json:"size"`
		Height 				int		`json:"height"`
		Version 			int		`json:"version"`
		Merkleroot 			string	`json:"merkleroot"`
		Segid 				int		`json:"segid"`
		Finalsaplingroot 	string	`json:"finalsaplingroot"`
		Tx 					[]struct {
			Txid			string	`json:"txid"`
			Overwintered	bool	`json:"overwintered"`
			Version			int	`json:"version"`
			Locktime		int	`json:"locktime"`
			Vin		[]struct {
				Coinbase	string `json:"coinbase"`
				Sequence	int64  `json:"sequence"`
			} `json:"vin"`
			Vout	[]struct {
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
		}	`json:"result"`
	Error	Error	`json:"error"`
	ID		string	`json:"id"`
}


func (appName AppType) GetBlock(h string, v int) (GetBlock, error) {
	query := APIQuery {
		Method:	`getblock`,
		Params:	`["`+h+`",`+strconv.Itoa(v)+`]`,
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