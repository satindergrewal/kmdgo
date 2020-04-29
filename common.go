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
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	// "fmt"
	"os"

	"github.com/satindergrewal/kmdgo/kmdutil"
)

const KMDGO_VERSION = `0.1.0`

type AppType string

type APIParams []interface{}

type APIResult struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type APIQuery struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewAppType(app AppType) *AppType {
	return &app
}

// APICall executes the blockchain RPC Calls and returns the data in JSON output.
// By default this API call tries to get the RPC information from the .conf file from the set blockchain's default directory on the system based on OS which is being executed on.
// If the local system doesn't has the blockchain running, you can also set this RPC settings in environment variables. The environment variable has to use the the combination of the appName + _RPCURL.
// Shell example:
// export komodo_RPCURL="http://192.168.1.1:"
// export komodo_RPCUSER="username"
// export komodo_RPCPASS="password"
// export komodo_RPCPORT="7771"
//
// These environment variables can also be set within Go code. Example:
// os.Setenv("komodo" + "_RPCURL", `http://127.0.0.1:`)
// os.Setenv("komodo" + "_RPCUSER", `username`)
// os.Setenv("komodo" + "_RPCPASS", `password`)
// os.Setenv("komodo" + "_RPCPORT", `7771`)
//
// As per this example, if for example using different SmartChain like "DEX", and the appName is set to example "DEX", just replace it word `komodo` with `DEX`.
func (appName AppType) APICall(q *APIQuery) string {
	var rpcurl, rpcuser, rpcpass, rpcport string

	key := string(appName) + "_RPCURL"
	// fmt.Println(key)
	_, ok := os.LookupEnv(key)
	if !ok {
		// fmt.Printf("%s not set\n", key)
		// Try to get RPC info from local default directory for set blockchain
		rpcuser, rpcpass, rpcport = kmdutil.AppRPCInfo(string(appName))
		rpcurl = `http://127.0.0.1:`
	} else {
		// fmt.Printf("%s=%s\n", key, val)
		rpcurl = os.Getenv(string(appName) + "_RPCURL")
		rpcuser = os.Getenv(string(appName) + "_RPCUSER")
		rpcpass = os.Getenv(string(appName) + "_RPCPASS")
		rpcport = os.Getenv(string(appName) + "_RPCPORT")
	}

	// fmt.Printf(" %s\n %s\n %s\n %s\n", rpcurl, rpcuser, rpcpass, rpcport)
	// fmt.Printf(" %T\n %T\n %T\n %T\n", rpcurl, rpcuser, rpcpass, rpcport)

	if rpcuser == "" && rpcpass == "" && rpcport == "" {
		// fmt.Println("EMPTY RPC INFO!")
		return "EMPTY RPC INFO!"
	}

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	url := rpcurl + rpcport
	// fmt.Println(url)

	var query_str string
	query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "` + q.Method + `", "params": ` + q.Params + ` }`
	//fmt.Println(query_str)

	query_byte := []byte(query_str)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
	req.Header.Set("Content-Type", "application/json")

	//req, err := http.NewRequest("POST", , nil)
	req.SetBasicAuth(rpcuser, rpcpass)
	resp, err := client.Do(req)

	if err != nil {
		// Check if the error is "connection refused", means if the daemon is running or not running or inaccessible due to any X reason.
		matched, _ := regexp.Match(`connection refused`, []byte(err.Error()))
		// fmt.Println(matched, err)
		if matched == true {
			s := string(`{"result": {}, "error": {"code":0,"message":"connection refused"}, "id":"kmdgo"}`)
			// fmt.Println(s)
			return s
		}
	}
	bodyText, err := ioutil.ReadAll(resp.Body)

	// fmt.Println(string(bodyText))

	var query_result map[string]interface{}
	if err := json.Unmarshal(bodyText, &query_result); err != nil {
		panic(err)
	}

	s := string(bodyText)
	return s
}
