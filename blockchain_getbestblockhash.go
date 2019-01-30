 //******************************************************************************
 //  Copyright © 2018-2019 Satinderjit Singh.                                   *
 //                                                                             *
 //  See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 //  the top-level directory of this distribution for the individual copyright  *
 //  holder information and the developer policies on copyright and licensing.  *
 //                                                                             *
 //  Unless otherwise agreed in a custom licensing agreement, no part of the    *
 //  kmdgo software, including this file may be copied, modified, propagated.   *
 //  or distributed except according to the terms contained in the LICENSE file *
 //                                                                             *
 //  Removal or modification of this copyright notice is prohibited.            *
 //                                                                             *
 //******************************************************************************/
package kmdgo

import (
    "errors"
    "encoding/json"
)

type GetBestBlockhash struct {
    Result interface{} `json:"result"`
    Error  Error       `json:"error"`
    ID     string      `json:"id"`
}

func (appName AppType) GetBestBlockhash() (GetBestBlockhash, error) {
    query := APIQuery {
        Method:     `getbestblockhash`,
        Params:   `[]`,
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