package main

import (
  "github.com/satindergrewal/kmdgo/kmdutil/bip39"
  "fmt"
)

func main(){
  // Generate a mnemonic for memorization or user-friendly seeds
  entropy, _ := bip39.NewEntropy(256)
  mnemonic, _ := bip39.NewMnemonic(entropy)

  // Generate a Bip32 HD wallet for the mnemonic and a user supplied password
  //seed := bip39.NewSeed(mnemonic, "Secret Passphrase")

  //fmt.Println(entropy)
  fmt.Println(mnemonic)
  //fmt.Println(seed)
}