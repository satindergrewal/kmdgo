# kmdgo

I created this Go package as part of my learning process of Go language. This as of this stage is working to for the basic things required, like just using the "methods" with "parameters" and it will give you the expected result in Go's native data set, structs, which you can use further to code your application. Would really appreciate if any other experienced Go language developers may push any improvements of fixes required for this package.

As of this stage, this package is compatible with **Komodo and all of it's Assetchains**.

# Documentation and Examples
 - Godoc link: https://godoc.org/github.com/satindergrewal/kmdgo/
 - Examples: Each API call has it's own example file created under [examples](/examples) directory.


### Install dependencies and setup environemnt
This package depends on saplinglib, so you have to get that package and set environment variables.

For Linux setup these environment variables:
```shell
export CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/"
export CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/linux -lsaplinglib -lpthread -ldl -lm"
```

For MacOS setup these environment variables:
```shell
export CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/"
export CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/darwin -lsaplinglib -framework Security"
```

For MingW cross-platform windows setup these environment variables:
```shell
export CGO_CFLAGS="-I$HOME/go/src/github.com/satindergrewal/saplinglib/src/"
export CGO_LDFLAGS="-L$HOME/go/src/github.com/satindergrewal/saplinglib/dist/win64 -lsaplinglib -lws2_32 -luserenv"
export CC="x86_64-w64-mingw32-gcc"
```

Install dependency:
```shell
go get -u github.com/satindergrewal/saplinglib
```

### Quick Example
To use this package do

```shell
go get -u github.com/satindergrewal/kmdgo
```


Example code:

```go
//main.go
package main

import (
	"fmt"
	"github.com/satindergrewal/kmdgo/kmdutil"
)

func main() {
	appName := "VRSC"
	dir := kmdutil.AppDataDir(appName, false)
	fmt.Println(dir)

	rpcuser, rpcpass, rpcport := kmdutil.AppRPCInfo(appName)
	fmt.Printf("RPC User: %s\nRPC Password: %s\nRPC Port: %s\n", rpcuser, rpcpass, rpcport)
}
```

Output:

```shell
$ go run main.go
/home/satinder/.komodo/VRSC
RPC User: user1773837506
RPC Password: passbae43ecd576o8ariotkgjhasdfiyosidulrkhdf9390bf03b68
RPC Port: 27486
```


## Supported API Calls

#### KMDUtil Specific Methods

- [x]	AppDataDir - Gives the path of cryptocurrency's Data directory based on it's symbol. Supports Only Komodo and Assetchains so far.
- [x]	AppRPCInfo - Provides RPC username, password and it's port by reading it's config file from data directory.
- [x]	GetTAddress - GetTAddress generates a public address using a seed phrase
- [x]	GetZAddress - GetZAddress generates a shielded sapling address using a seed phrase
- [x]	GetIguanaWallet - GetIguanaWallet returns a set of public and shielded addresses generated based on a mnemonic seed phrase
- [x]	BytesToString - BytesToString converts and returns bytes to string

##### NOTE: [saplinglib](saplinglib) is required for kmdutil methods `GetZAddress` and `GetIguanaWallet`


#### CClib
- [ ]	cclib method [evalcode] [JSON params]
- [x]	cclibaddress [evalcode] [pubkey]
- [x]	cclibinfo

#### DEX
- [x]	DEX_anonsend message priority destpub33
- [x]	DEX_broadcast hex [priority [tagA [tagB [pubkey33 [volA [volB]]]]]]
- [x]	DEX_cancel id [pubkey33 [tagA tagB]]
- [x]	DEX_get id
- [x]	DEX_list stopat minpriority tagA tagB pubkey33 [minA maxA minB maxB [stophash]]
- [ ]	DEX_notarize coin height
- [x]	DEX_orderbook maxentries minpriority tagA tagB pubkey33 [minA maxA minB maxB]
- [x]	DEX_publish filename priority sliceid
- [x]	DEX_setpubkey pubkey33
- [x]	DEX_stats
- [ ]	DEX_stream filename priority
- [ ]	DEX_streamsub filename priority pubkey
- [ ]	DEX_subscribe filename priority id [publisher33]

#### FSM
- [ ]	FSMaddress [pubkey]
- [ ]	FSMcreate name states
- [ ]	FSMinfo fundingtxid
- [ ]	FSMlist

#### Addressindex
- [ ]	getaddressbalance
- [ ]	getaddressdeltas
- [ ]	getaddressmempool
- [ ]	getaddresstxids
- [ ]	getaddressutxos
- [ ]	getsnapshot

#### Auction
- [ ]	auctionaddress [pubkey]

#### Blockchain
- [x]	coinsupply <height>
- [x]	getbestblockhash
- [x]	getblock "hash|height" ( verbosity )
- [x]	getblockchaininfo
- [x]	getblockcount
- [x]	getblockhash index
- [ ]	getblockhashes timestamp -- Todo
- [x]	getblockheader "hash" ( verbose )
- [x]	getchaintips
- [x]	getdifficulty
- [x]	getmempoolinfo
- [x]	getrawmempool ( verbose )
- [ ]	getspentinfo -- Todo
- [x]	gettxout "txid" n ( includemempool )
- [x]	gettxoutproof ["txid",...] ( blockhash )
- [x]	gettxoutsetinfo
- [ ]	kvsearch key -- Todo
- [ ]	kvupdate key "value" days passphrase -- Todo
- [x]	minerids needs height
- [x]	notaries height timestamp
- [x]	verifychain ( checklevel numblocks )
- [x]	verifytxoutproof "proof"

#### Channels
- [ ]	channelsaddress destpubkey
- [ ]	channelsclose opentxid
- [ ]	channelsinfo [opentxid]
- [ ]	channelsinfo
- [ ]	channelsopen destpubkey numpayments payment
- [ ]	channelspayment opentxid amount [secret]
- [ ]	channelsrefund opentxid closetxid

#### Control
- [x]	getinfo
- [ ]	help ( "command" )
- [x]	stop

#### Crosschain
- [ ]	MoMoMdata symbol kmdheight ccid
- [ ]	assetchainproof needs a txid
- [ ]	calc_MoM height MoMdepth
- [ ]	getNotarisationsForBlock blockHash
- [ ]	height_MoM height
- [ ]	migrate_completeimporttransaction importTx
- [ ]	migrate_converttoexport rawTx dest_symbol export_amount
- [ ]	migrate_createimporttransaction burnTx payouts
- [ ]	scanNotarisationsDB blockHeight symbol [blocksLimit=1440]

#### Dice
- [ ]	diceaddfunds name fundingtxid amount
- [ ]	diceaddress [pubkey]
- [ ]	dicebet name fundingtxid amount odds
- [ ]	dicefinish name fundingtxid bettxid
- [ ]	dicefund name funds minbet maxbet maxodds timeoutblocks
- [ ]	diceinfo fundingtxid
- [ ]	dicelist
- [ ]	dicestatus name fundingtxid bettxid

#### Disclosure
- [ ]	z_getpaymentdisclosure "txid" "js_index" "output_index" ("message") 
- [ ]	z_validatepaymentdisclosure "paymentdisclosure"

#### Faucet
- [x]	faucetaddress [pubkey]
- [x]	faucetfund amount
- [x]	faucetget
- [x]	faucetinfo


#### Rogue Game
- [x]	cclib newgame [evalcode] "[maxplayers, buyin]"
- [x]	cclib gameinfo [evalcode] "[gametxid]"
- [x]	cclib pending [evalcode]
- [x]	cclib register [evalcode] "[gametxid, [playertxid]]"
- [ ]	cclib keystrokes [evalcode] "[gametxid, keystrokes]" -- This is only executed by the ROGUE dApp. Not need to run manully
- [x]	cclib bailout [evalcode] "[gametxid]"
- [x]	cclib highlander [evalcode] "[gametxid]"
- [x]	cclib playerinfo [evalcode] "[playertxid]"
- [x]	cclib players [evalcode]
- [x]	cclib games [evalcode]
- [x]	cclib setname [evalcode] "[pname]"

#### Gateways
- [ ]	gatewaysaddress [pubkey]
- [ ]	gatewaysbind tokenid oracletxid coin tokensupply M N pubkey(s)
- [ ]	gatewaysclaim bindtxid coin deposittxid destpub amount
- [ ]	gatewayscompletesigning withdrawtxid coin hex
- [ ]	gatewaysdeposit bindtxid height coin cointxid claimvout deposithex proof destpub amount
- [ ]	gatewaysinfo bindtxid
- [ ]	gatewayslist
- [ ]	gatewaysmarkdone completesigningtx coin
- [ ]	gatewaysmultisig txidaddr
- [ ]	gatewayspartialsign txidaddr refcoin hex
- [ ]	gatewayspending bindtxid coin
- [ ]	gatewaysprocessed bindtxid coin
- [ ]	gatewayswithdraw bindtxid coin withdrawpub amount

#### Generating
- [ ]	generate numblocks
- [ ]	getgenerate
- [ ]	setgenerate generate ( genproclimit )

#### Heir
- [ ]	heiraddress func txid amount [destpubkey]

#### Lotto
- [ ]	lottoaddress [pubkey]

#### Mining
- [x]	getblocksubsidy height
- [x]	getblocktemplate ( "jsonrequestobject" )
- [x]	getlocalsolps
- [x]	getmininginfo
- [x]	getnetworkhashps ( blocks height )
- [x]	getnetworksolps ( blocks height )
- [x]	prioritisetransaction <txid> <priority delta> <fee delta>
- [x]	submitblock "hexdata" ( "jsonparametersobject" )

#### Network
- [x]	addnode "node" "add|remove|onetry"
- [x]	clearbanned
- [x]	disconnectnode "node" 
- [x]	getaddednodeinfo dns ( "node" )
- [x]	getconnectioncount
- [x]	getdeprecationinfo
- [x]	getnettotals
- [x]	getnetworkinfo
- [x]	getpeerinfo
- [x]	listbanned
- [x]	ping
- [x]	setban "ip(/netmask)" "add|remove" (bantime) (absolute)

#### Oracles
- [ ]	oraclesaddress [pubkey]
- [ ]	oraclescreate name description format
- [ ]	oraclesdata oracletxid hexstr
- [ ]	oraclesinfo oracletxid
- [ ]	oracleslist
- [ ]	oraclesregister oracletxid datafee
- [ ]	oraclessamples oracletxid batonutxo num
- [ ]	oraclessubscribe oracletxid publisher amount

#### Payments
- [ ]	paymentsaddress [pubkey]

#### Pegs
- [ ]	pegssaddress [pubkey]

#### Prices
- [ ]	pricesaddfunding fundingtxid bettoken amount
- [ ]	pricesaddress [pubkey]
- [ ]	pricesbet fundingtxid bettoken amount leverage
- [ ]	pricescreate bettoken oracletxid margin mode longtoken shorttoken maxleverage funding N [pubkeys]
- [ ]	pricesfinish fundingtxid bettoken bettxid
- [ ]	pricesinfo fundingtxid
- [ ]	priceslist
- [ ]	pricesstatus fundingtxid bettoken bettxid

#### Rawtransactions
- [x]	createrawtransaction [{"txid":"id","vout":n},...] {"address":amount,...} ( locktime ) ( expiryheight )
- [x]	decoderawtransaction "hexstring"
- [x]	decodescript "hex"
- [x]	fundrawtransaction "hexstring"
- [x]	getrawtransaction "txid" ( verbose )
- [x]	sendrawtransaction "hexstring" ( allowhighfees )
- [x]	signrawtransaction "hexstring" ( [{"txid":"id","vout":n,"scriptPubKey":"hex","redeemScript":"hex"},...] ["privatekey1",...] sighashtype )

#### Rewards
- [ ]	rewardsaddfunding name fundingtxid amount
- [ ]	rewardsaddress [pubkey]
- [ ]	rewardscreatefunding name amount APR mindays maxdays mindeposit
- [ ]	rewardsinfo fundingtxid
- [ ]	rewardslist
- [ ]	rewardslock name fundingtxid amount
- [ ]	rewardsunlock name fundingtxid [txid]

#### Tokens
- [ ]	tokenaddress [pubkey]
- [ ]	tokenask numtokens tokenid price
- [ ]	tokenbalance tokenid [pubkey]
- [ ]	tokenbid numtokens tokenid price
- [ ]	tokencancelask tokenid asktxid
- [ ]	tokencancelbid tokenid bidtxid
- [ ]	tokenconvert evalcode tokenid pubkey amount
- [ ]	tokencreate name supply description
- [ ]	tokenfillask tokenid asktxid fillunits
- [ ]	tokenfillbid tokenid bidtxid fillamount
- [ ]	tokeninfo tokenid
- [ ]	tokenlist
- [ ]	tokenorders [tokenid]
- [ ]	tokentransfer tokenid destpubkey amount

#### Triggers
- [ ]	triggersaddress [pubkey]

#### Util
- [ ]	createmultisig nrequired ["key",...]
- [ ]	estimatefee nblocks
- [ ]	estimatepriority nblocks
- [ ]	invalidateblock "hash"
- [ ]	jumblr_deposit "depositaddress"
- [ ]	jumblr_pause
- [ ]	jumblr_resume
- [ ]	jumblr_secret "secretaddress"
- [ ]	reconsiderblock "hash"
- [ ]	txnotarizedconfirmed txid
- [x]	validateaddress "komodoaddress"
- [x]	verifymessage "komodoaddress" "signature" "message"
- [x]	z_validateaddress "zaddr"

#### Wallet
- [x]	addmultisigaddress nrequired ["key",...] ( "account" )
- [x]	backupwallet "destination"
- [x]	dumpprivkey "t-addr"
- [x]	dumpwallet "filename"
- [x]	encryptwallet "passphrase"
- [x]	getaccount "KMD_address"
- [x]	getaccountaddress "account"
- [x]	getaddressesbyaccount "account"
- [x]	getbalance ( "account" minconf includeWatchonly )
- [x]	getnewaddress ( "account" )
- [x]	getrawchangeaddress
- [x]	getreceivedbyaccount "account" ( minconf )
- [x]	getreceivedbyaddress "KMD_address" ( minconf )
- [x]	gettransaction "txid" ( includeWatchonly )
- [x]	getunconfirmedbalance
- [x]	getwalletinfo
- [x]	importaddress "address" ( "label" rescan )
- [x]	importprivkey "komodoprivkey" ( "label" rescan )
- [x]	importwallet "filename"
- [x]	keypoolrefill ( newsize )
- [x]	listaccounts ( minconf includeWatchonly)
- [x]	listaddressgroupings
- [x]	listlockunspent
- [ ]	listreceivedbyaccount ( minconf includeempty includeWatchonly) -- DEPRECATED
- [x]	listreceivedbyaddress ( minconf includeempty includeWatchonly)
- [x]	listsinceblock ( "blockhash" target-confirmations includeWatchonly)
- [x]	listtransactions ( "account" count from includeWatchonly)
- [x]	listunspent ( minconf maxconf  ["address",...] )
- [x]	lockunspent unlock [{"txid":"txid","vout":n},...]
- [ ]	move "fromaccount" "toaccount" amount ( minconf "comment" ) -- DEPRECATED
- [ ]	resendwallettransactions -- "Intended only for testing;" --- May be Later --- Todo
- [x]	sendfrom "fromaccount" "toKMDaddress" amount ( minconf "comment" "comment-to" )
- [x]	sendmany "fromaccount" {"address":amount,...} ( minconf "comment" ["address",...] )
- [x]	sendtoaddress "KMD_address" amount ( "comment" "comment-to" subtractfeefromamount )
- [ ]	setaccount "KMD_address" "account" -- DEPRECATED
- [x]	setpubkey
- [x]	settxfee amount
- [x]	signmessage "t-addr" "message"
- [x]	walletlock
- [x]	walletpassphrase "passphrase" seconds
- [x]	walletpassphrasechange "oldpassphrase" "newpassphrase"
- [x]	z_exportkey "zaddr"
- [x]	z_exportviewingkey "zaddr"
- [x]	z_exportwallet "filename"
- [x]	z_getbalance "address" ( minconf )
- [x]	z_getnewaddress ( type )
- [x]	z_getoperationresult (["operationid", ... ]) 
- [x]	z_getoperationstatus (["operationid", ... ]) 
- [x]	z_gettotalbalance ( minconf includeWatchonly )
- [x]	z_importkey "zkey" ( rescan startHeight )
- [x]	z_importviewingkey "vkey" ( rescan startHeight )
- [x]	z_importwallet "filename"
- [x]	z_listaddresses ( includeWatchonly )
- [x]	z_listoperationids
- [x]	z_listreceivedbyaddress "address" ( minconf )
- [x]	z_listunspent ( minconf maxconf includeWatchonly ["zaddr",...] )
- [x]	z_mergetoaddress ["fromaddress", ... ] "toaddress" ( fee ) ( transparent_limit ) ( shielded_limit ) ( memo )
- [x]	z_sendmany "fromaddress" [{"address":... ,"amount":...},...] ( minconf ) ( fee )
- [x]	z_shieldcoinbase "fromaddress" "tozaddress" ( fee ) ( limit )
- [ ]	zcbenchmark benchmarktype samplecount -- Todo
- [ ]	zcrawjoinsplit rawtx inputs outputs vpub_old vpub_new  -- Todo
- [ ]	zcrawkeygen  -- Todo
- [ ]	zcrawreceive zcsecretkey encryptednote  -- Todo
- [ ]	zcsamplejoinsplit  -- Todo

## Verus blockchain specific Methods

#### Verus Identity
- [x]	getidentity "name"
- [x]	listidentities (includecansign) (includewatchonly)
- [ ]	recoveridentity "jsonidentity" (returntx)
- [x]	registeridentity "jsonidregistration" feeoffer
- [x]	registernamecommitment "name" "controladdress" ("referralidentity")
- [ ]	revokeidentity "nameorID" (returntx)
- [x]	updateidentity "jsonidentity" (returntx)

#### Verus Multichain
- [ ]	addmergedblock "hexdata" ( "jsonparametersobject" )
- [ ]	definecurrency '{"name": "BAAS", ..., "nodes":[{"networkaddress":"identity"},..]}'
- [ ]	getcrossnotarization "systemid" '["notarizationtxid1", "notarizationtxid2", ...]' - TODO - Need example command/output
- [x]	getcurrency "chainname"
- [x]	getcurrencyconverters ["currency1","currency2",...]'
- [x]	getcurrencystate "n"
- [x]	getexports "chainname"
- [x]	getimports "chainname"
- [x]	getinitialcurrencystate "name"
- [x]	getlastimportin "fromname"
- [ ]	getlatestimportsout "name" "lastimporttransaction" "importtxtemplate"
- [ ]	getblocktemplate ( "jsonrequestobject" )
- [x]	getnotarizationdata "currencyid" accepted
- [ ]	getpendingtransfers "chainname"
- [ ]	getsaplingtree "n"
- [ ]	listcurrencies (includeexpired)
- [ ]	paynotarizationrewards "currencyid" "amount" "billingperiod"
- [ ]	refundfailedlaunch "currencyid"
- [ ]	reserveexchange '[{"toreserve": 1, "recipient": "RRehdmUV7oEAqoZnzEGBH34XysnWaBatct", "amount": 5.0}]'
- [ ]	sendcurrency "fromaddress" '[{"address":... ,"amount":...},...]' (returntx)
- [ ]	submitacceptednotarization "hextx"

#### Verus Wallet
- [ ]	convertpassphrase "walletpassphrase"
- [ ]	signfile "address or identity" "filepath/filename" "curentsig"
- [ ]	z_getmigrationstatus
- [ ]	z_setmigration enabled
- [ ]	z_viewtransaction "txid"