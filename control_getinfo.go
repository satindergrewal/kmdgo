package kmdgo

import(
    "net/http"
    "io/ioutil"
    "log"
    "bytes"
    "encoding/json"
    "github.com/satindergrewal/kmdgo/kmdutil"
)

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

func GetinfoJsonValue(appName string) string {
    //appName := "komodo"
    rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)

    client := &http.Client{}
    url := `http://127.0.0.1:`+rpcport

    query_byte := []byte(`{"jsonrpc": "1.0", "id":"curltest", "method": "getinfo", "params": [] }`)

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

func (i GetInfo) DisplayGetinfo() GetInfo {
    return i
}

func ResultGetInfo(appName string) GetInfo {
    getinfoJson := GetinfoJsonValue(appName)
    var getinfo GetInfo
    json.Unmarshal([]byte(getinfoJson), &getinfo)
    return getinfo
}