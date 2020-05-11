// kmdgo/kmdutils troubleshooting file
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/satindergrewal/kmdgo/kmdutil"
)

//RPCUsername, RPCPassword string = "user60de7828fd8985d3", "ce3f74430f82aa34b58aeba4b37a3373"

func BytesToString(data []byte) string {
	return string(data[:])
}

type GetInfo struct {
	Result struct {
		Version             int     `json:"version"`
		Protocolversion     int     `json:"protocolversion"`
		KMDversion          string  `json:"KMDversion"`
		Notarized           int     `json:"notarized"`
		PrevMoMheight       int     `json:"prevMoMheight"`
		Notarizedhash       string  `json:"notarizedhash"`
		Notarizedtxid       string  `json:"notarizedtxid"`
		NotarizedtxidHeight string  `json:"notarizedtxid_height"`
		KMDnotarizedHeight  int     `json:"KMDnotarized_height"`
		NotarizedConfirms   int     `json:"notarized_confirms"`
		Walletversion       int     `json:"walletversion"`
		Balance             float64 `json:"balance"`
		Blocks              int     `json:"blocks"`
		Longestchain        int     `json:"longestchain"`
		Timeoffset          int     `json:"timeoffset"`
		Tiptime             int     `json:"tiptime"`
		Connections         int     `json:"connections"`
		Proxy               string  `json:"proxy"`
		Difficulty          float64 `json:"difficulty"`
		Testnet             bool    `json:"testnet"`
		Keypoololdest       int     `json:"keypoololdest"`
		Keypoolsize         int     `json:"keypoolsize"`
		Paytxfee            float64 `json:"paytxfee"`
		Relayfee            float64 `json:"relayfee"`
		Errors              string  `json:"errors"`
		CCid                int     `json:"CCid"`
		Name                string  `json:"name"`
		P2Pport             int     `json:"p2pport"`
		Rpcport             int     `json:"rpcport"`
		Magic               int     `json:"magic"`
		Premine             int     `json:"premine"`
		Reward              int64   `json:"reward"`
		Halving             int     `json:"halving"`
		Commission          int     `json:"commission"`
	} `json:"result"`
	Error interface{} `json:"error"`
	ID    string      `json:"id"`
}

func basicAuth() string {
	appName := "komodo"

	//appDir := kmdutil.AppDataDir(appName, false)
	//fmt.Println(appDir)

	rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)

	client := &http.Client{}

	url := `http://127.0.0.1:` + rpcport
	fmt.Println("URL:>", url)

	query_byte := []byte(`{"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }`)
	fmt.Printf("Query: %s\n\n", query_byte)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
	req.Header.Set("Content-Type", "application/json")

	//    req, err := http.NewRequest("POST", , nil)
	req.SetBasicAuth(rpcuser, rpcpass)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%T\n\n", bodyText)

	var query_result map[string]interface{}

	if err := json.Unmarshal(bodyText, &query_result); err != nil {
		panic(err)
	}
	//fmt.Println(query_result)
	fmt.Printf("\n\n")

	//fmt.Println(query_result["result"].(map[string]interface{})["connections"])

	/*    parsed_result := query_result["result"]
	      fmt.Println("result: ", parsed_result)
	      fmt.Printf("\n\n")
	*/

	s := string(bodyText)
	return s
}

func main() {

	fmt.Println("requesting...\n")
	getinfoJson := basicAuth()
	fmt.Println(getinfoJson)

	fmt.Printf("%T\n", getinfoJson)

	var getinfo GetInfo
	json.Unmarshal([]byte(getinfoJson), &getinfo)
	fmt.Println(getinfo.Result.Version)
	//fmt.Printf("Version: %d", getinfo[0].Version)
}
