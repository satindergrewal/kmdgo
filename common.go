package kmdgo

import (
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

type AppType string

type APIResult struct {
    Result interface{}  `json:"result"`
    Error interface{}   `json:"error"`
    ID string           `json:"id"`
}

type APIQuery struct {
    Method string `json:"method"`
    Params string `json:"params"`
}

type Error struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

func (appName AppType) APICall(q APIQuery) string {
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(string(appName))

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport

    var query_str string
    query_str = `{"jsonrpc": "1.0", "id":"kmdgo", "method": "`+q.Method+`", "params": `+q.Params+` }`

    query_byte := []byte(query_str)

    req, err := http.NewRequest("POST", url, bytes.NewBuffer(query_byte))
    req.Header.Set("Content-Type", "application/json")

    //req, err := http.NewRequest("POST", , nil)
    req.SetBasicAuth(rpcuser, rpcpass)
    resp, err := client.Do(req)
    if err != nil{
        log.Fatal(err)
    }
    bodyText, err := ioutil.ReadAll(resp.Body)

    var query_result map[string]interface{}
    if err := json.Unmarshal(bodyText, &query_result); err != nil {
        panic(err)
    }

    s := string(bodyText)
    return s
}