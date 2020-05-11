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

package kmdutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func BytesToString(data []byte) string {
	return string(data[:])
}

func AppRPCInfo(appName string) (string, string, string) {
	appDir := AppDataDir(appName, false)

	var appConf string

	if strings.ToLower(appName) == "komodo" || strings.ToLower(appName) == ".komodo" {
		if runtime.GOOS == "darwin" {
			appConf = filepath.Join(appDir, strings.Title(appName)+`.conf`)
		} else {
			appConf = filepath.Join(appDir, strings.ToLower(appName)+`.conf`)
		}
	} else {
		appConf = filepath.Join(appDir, strings.ToUpper(appName)+`.conf`)
	}

	confdata, err := ioutil.ReadFile(appConf)
	if err != nil {
		switch err {
		case os.ErrInvalid:
			//Do stuff
		case os.ErrPermission:
			//Do stuff
		case os.ErrNotExist:
			// fmt.Println("File Does not exists")
			// log.Println(err)
		default:
			// fmt.Println(err)
		}
		return "", "", ""
	}

	var rpcu = regexp.MustCompile("(?m)^rpcuser=.+$")
	var rpcpass = regexp.MustCompile("(?m)^rpcpassword=.+$")
	var rpcport = regexp.MustCompile("(?m)^rpcport=.+$")

	// fmt.Println("confdata: ", confdata)

	bytestr := BytesToString(confdata)

	rpcuser_line := rpcu.FindString(bytestr)
	rpcpass_line := rpcpass.FindString(bytestr)
	rpcpport_line := rpcport.FindString(bytestr)

	// fmt.Println(bytestr)

	//AppRPCuser := strings.TrimLeft(strings.TrimLeft(rpcuser_line,`rpcuser`),`=`)
	//AppRPCpass := strings.TrimLeft(strings.TrimLeft(rpcpass_line,`rpcpassword`),`=`)
	AppRPCport := strings.TrimLeft(strings.TrimLeft(rpcpport_line, `rpcport`), `=`)

	//AppRPCuser := rpcuser_line[8:]
	AppRPCuser := strings.TrimLeft(rpcuser_line, `rpcuser`)[1:]
	AppRPCpass := strings.TrimLeft(rpcpass_line, `rpcpassword`)[1:]
	//AppRPCport := strings.TrimLeft(rpcpport_line,`rpcport`)[1:]

	if AppRPCport == "" && appName == "komodo" {
		AppRPCport = "7771"
	}

	//TODO
	// Check if server=1 and update/change it
	// Check if rpcuser and rpcpassword is the known default value from QT wallet,
	// example "komodo" and "local321", and prompt/update it with random values.

	// Removing any \n or \r\n from the stripped string from .conf file
	AppRPCuser = strings.ReplaceAll(AppRPCuser, "\n", "")
	AppRPCuser = strings.ReplaceAll(AppRPCuser, "\r", "")
	AppRPCpass = strings.ReplaceAll(AppRPCpass, "\n", "")
	AppRPCpass = strings.ReplaceAll(AppRPCpass, "\r", "")
	AppRPCport = strings.ReplaceAll(AppRPCport, "\n", "")
	AppRPCport = strings.ReplaceAll(AppRPCport, "\r", "")
	return AppRPCuser, AppRPCpass, AppRPCport
}
