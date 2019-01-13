# kmdgo


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
- [ ]	getblockhashes timestamp
- [x]	getblockheader "hash" ( verbose )
- [x]	getchaintips
- [x]	getdifficulty
- [ ]	getmempoolinfo
- [ ]	getrawmempool ( verbose )
- [ ]	getspentinfo
- [ ]	gettxout "txid" n ( includemempool )
- [ ]	gettxoutproof ["txid",...] ( blockhash )
- [ ]	gettxoutsetinfo
- [ ]	kvsearch key
- [ ]	kvupdate key "value" days passphrase
- [ ]	minerids needs height
- [ ]	notaries height timestamp
- [ ]	verifychain ( checklevel numblocks )
- [ ]	verifytxoutproof "proof"

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
- [ ]	stop

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
- [ ]	faucetaddress [pubkey]
- [ ]	faucetfund amount
- [ ]	faucetget
- [ ]	faucetinfo

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
- [ ]	getblocksubsidy height
- [ ]	getblocktemplate ( "jsonrequestobject" )
- [ ]	getlocalsolps
- [ ]	getmininginfo
- [ ]	getnetworkhashps ( blocks height )
- [ ]	getnetworksolps ( blocks height )
- [ ]	prioritisetransaction <txid> <priority delta> <fee delta>
- [ ]	submitblock "hexdata" ( "jsonparametersobject" )

#### Network
- [ ]	addnode "node" "add|remove|onetry"
- [ ]	clearbanned
- [ ]	disconnectnode "node" 
- [ ]	getaddednodeinfo dns ( "node" )
- [ ]	getconnectioncount
- [ ]	getdeprecationinfo
- [ ]	getnettotals
- [ ]	getnetworkinfo
- [ ]	getpeerinfo
- [ ]	listbanned
- [ ]	ping
- [ ]	setban "ip(/netmask)" "add|remove" (bantime) (absolute)

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
- [ ]	createrawtransaction [{"txid":"id","vout":n},...] {"address":amount,...} ( locktime ) ( expiryheight )
- [ ]	decoderawtransaction "hexstring"
- [ ]	decodescript "hex"
- [ ]	fundrawtransaction "hexstring"
- [ ]	getrawtransaction "txid" ( verbose )
- [ ]	sendrawtransaction "hexstring" ( allowhighfees )
- [ ]	signrawtransaction "hexstring" ( [{"txid":"id","vout":n,"scriptPubKey":"hex","redeemScript":"hex"},...] ["privatekey1",...] sighashtype )

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
- [ ]	validateaddress "komodoaddress"
- [ ]	verifymessage "komodoaddress" "signature" "message"
- [ ]	z_validateaddress "zaddr"

#### Wallet
- [ ]	addmultisigaddress nrequired ["key",...] ( "account" )
- [ ]	backupwallet "destination"
- [ ]	dumpprivkey "t-addr"
- [ ]	dumpwallet "filename"
- [ ]	encryptwallet "passphrase"
- [ ]	getaccount "KMD_address"
- [ ]	getaccountaddress "account"
- [ ]	getaddressesbyaccount "account"
- [ ]	getbalance ( "account" minconf includeWatchonly )
- [ ]	getnewaddress ( "account" )
- [ ]	getrawchangeaddress
- [ ]	getreceivedbyaccount "account" ( minconf )
- [ ]	getreceivedbyaddress "KMD_address" ( minconf )
- [ ]	gettransaction "txid" ( includeWatchonly )
- [ ]	getunconfirmedbalance
- [ ]	getwalletinfo
- [ ]	importaddress "address" ( "label" rescan )
- [ ]	importprivkey "komodoprivkey" ( "label" rescan )
- [ ]	importwallet "filename"
- [ ]	keypoolrefill ( newsize )
- [ ]	listaccounts ( minconf includeWatchonly)
- [ ]	listaddressgroupings
- [ ]	listlockunspent
- [ ]	listreceivedbyaccount ( minconf includeempty includeWatchonly)
- [ ]	listreceivedbyaddress ( minconf includeempty includeWatchonly)
- [ ]	listsinceblock ( "blockhash" target-confirmations includeWatchonly)
- [ ]	listtransactions ( "account" count from includeWatchonly)
- [ ]	listunspent ( minconf maxconf  ["address",...] )
- [ ]	lockunspent unlock [{"txid":"txid","vout":n},...]
- [ ]	move "fromaccount" "toaccount" amount ( minconf "comment" )
- [ ]	resendwallettransactions
- [ ]	sendfrom "fromaccount" "toKMDaddress" amount ( minconf "comment" "comment-to" )
- [ ]	sendmany "fromaccount" {"address":amount,...} ( minconf "comment" ["address",...] )
- [ ]	sendtoaddress "KMD_address" amount ( "comment" "comment-to" subtractfeefromamount )
- [ ]	setaccount "KMD_address" "account"
- [ ]	setpubkey
- [ ]	settxfee amount
- [ ]	signmessage "t-addr" "message"
- [ ]	z_exportkey "zaddr"
- [ ]	z_exportviewingkey "zaddr"
- [ ]	z_exportwallet "filename"
- [ ]	z_getbalance "address" ( minconf )
- [ ]	z_getnewaddress ( type )
- [ ]	z_getoperationresult (["operationid", ... ]) 
- [ ]	z_getoperationstatus (["operationid", ... ]) 
- [ ]	z_gettotalbalance ( minconf includeWatchonly )
- [ ]	z_importkey "zkey" ( rescan startHeight )
- [ ]	z_importviewingkey "vkey" ( rescan startHeight )
- [ ]	z_importwallet "filename"
- [ ]	z_listaddresses ( includeWatchonly )
- [ ]	z_listoperationids
- [ ]	z_listreceivedbyaddress "address" ( minconf )
- [ ]	z_listunspent ( minconf maxconf includeWatchonly ["zaddr",...] )
- [ ]	z_mergetoaddress ["fromaddress", ... ] "toaddress" ( fee ) ( transparent_limit ) ( shielded_limit ) ( memo )
- [ ]	z_sendmany "fromaddress" [{"address":... ,"amount":...},...] ( minconf ) ( fee )
- [ ]	z_shieldcoinbase "fromaddress" "tozaddress" ( fee ) ( limit )
- [ ]	zcbenchmark benchmarktype samplecount
- [ ]	zcrawjoinsplit rawtx inputs outputs vpub_old vpub_new
- [ ]	zcrawkeygen
- [ ]	zcrawreceive zcsecretkey encryptednote
- [ ]	zcsamplejoinsplit
