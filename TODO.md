


#### API calls to add later

To use the following command(s) command need to start blockchain with the options such as:

```
komodod -timestampindex=1 -ac_name=TEST
komodod -addressindex=-1 -ac_name=TEST
komodod -spentindex=-1 -ac_name=TEST
komodod -txindex=-1 -ac_name=TEST
```

```

# Need -timestampindex=1 active in komodod
komodo-cli getblockhashes 1547424000 1547526260

# Need -spentindex=-1 active in komodod
komodo-cli getspentinfo '{"txid": "16298ea1de39a2818210fb54435c47859ae8746ff9df6c557ea88fc612750e4f", "index": 0}'
```

