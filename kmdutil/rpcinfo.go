// Copyright (c) 2018 Satinder Grewal
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package kmdutil

import (
	"path/filepath"
	"runtime"
	"log"
	"strings"
	"regexp"
	"io/ioutil"
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
		log.Fatal( err )
	}

	var rpcu = regexp.MustCompile("(?m)^rpcuser=.+$")
	var rpcpass = regexp.MustCompile("(?m)^rpcpassword=.+$")
	var rpcport = regexp.MustCompile("(?m)^rpcport=.+$")

	bytestr := BytesToString(confdata)

	rpcuser_line := rpcu.FindString(bytestr)
	rpcpass_line := rpcpass.FindString(bytestr)
	rpcpport_line := rpcport.FindString(bytestr)

	//AppRPCuser := strings.TrimLeft(strings.TrimLeft(rpcuser_line,`rpcuser`),`=`)
	//AppRPCpass := strings.TrimLeft(strings.TrimLeft(rpcpass_line,`rpcpassword`),`=`)
	AppRPCport := strings.TrimLeft(strings.TrimLeft(rpcpport_line,`rpcport`),`=`)

	//AppRPCuser := rpcuser_line[8:]
	AppRPCuser := strings.TrimLeft(rpcuser_line,`rpcuser`)[1:]
	AppRPCpass := strings.TrimLeft(rpcpass_line,`rpcpassword`)[1:]
	//AppRPCport := strings.TrimLeft(rpcpport_line,`rpcport`)[1:]

	if AppRPCport == "" && appName == "komodo" {
		AppRPCport = "7771"
	}

	return AppRPCuser, AppRPCpass, AppRPCport
}