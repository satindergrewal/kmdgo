/******************************************************************************
 * Copyright © 2018-2019 Satinderjit Singh.                                   *
 *                                                                            *
 * See the AUTHORS, DEVELOPER-AGREEMENT and LICENSE files at                  *
 * the top-level directory of this distribution for the individual copyright  *
 * holder information and the developer policies on copyright and licensing.  *
 *                                                                            *
 * Unless otherwise agreed in a custom licensing agreement, no part of the    *
 * kmdgo software, including this file may be copied, modified, propagated.   *
 * or distributed except according to the terms contained in the LICENSE file *
 *                                                                            *
 * Removal or modification of this copyright notice is prohibited.            *
 *                                                                            *
 ******************************************************************************/
package main

import (
	"fmt"
	"log"
	"github.com/satindergrewal/kmdgo"
)

func main() {
	var appName kmdgo.AppType
	appName = `komodo`

	var pf kmdgo.VerifyTxOutProof

	txhex := `040000006d06946714aaaa0cf1e7cc3b5bdc530ac6132c71d850b9f3dd1148ae57fac00761102373fbb2f110c308fd14919f32090432fc7639410741527def45a2bc674d511fe835159b3826b0cd67bb87712ad3b490d81be3a09f01658e0d054a1f815391eb3f5c5e44021d02003c23093753d48fddf9cecd0895836fbf8d1cbd7950a14ae7be8baa750000fd400502ea9bce4b98771ff66a70a7cd1de42d783e5c7f622b6d158b51dd60ed45add60a98cafbf9e3b5d908171b757ea0fc7203bbc90456063172af05c59cf9f1ca24ece7e9ff5e07c555a4fbb852e8dd2b0686789ed70649bea69e8549d2515cd248555c9c6f36cc7a87fc4ba00cc6fdf499afcdfed508594bcd617a4cd9c26d0787e1b3bba308d3fec6a3c747a978da1f593957ce07d46f9e7d530b3346dc95006e4696db26bc5db542043881e02c617bab7b4e61ba555abf3926043ccd4543cdcd59ef53636f33e674960568464ac081fe3a764bdf02616fd6864f3cfeaa69b7f8781f86769e7e73591514734c1c7f730e67b69f24f0ec4f51e73da0d507a6c090d35c029317eac5adc478305da568bf9bc14c118eec12a5f2c39640486c73c6f4a65d8895ade31d2b44d51411b7095ca534245c4e89c5c15ffe3fd327d191bb9b5ce1b7d2d0138b1e2407d221c2b80aa3060b2ba0319cc32b3a1ef1fdbce311bd97e2943b460ecc9140bba2ed6b55c605c1f34f6e2a2e9c7e204a1e65b326c59d0379f540d49e38c7c3a185327e6f0d47cf1d72f55cc69b01f29822b9e26f6751863e4835075cf997ea1b5ea2ef9cd0dc3549f17d9c104e76941aa36f1b6ae4a1874b8643cfc45fe09a4f2c9dec720f87d496e24b0ecaa76380fa2afff4812643517de92272877915108e2ccc0ae4add54d59ce22fb933adf07154a7bb886e6d0854e41850c7aa8c94b751a4971086a436ccb9b99418ea6b6b867d134cf41e19e70e116dc2bcd91233693d5f7c19e37d95035ace170570d1d195d4ee88b6d57b42293304da22da2245a1909c50c91ecb1e3daec90fca7c0d9fa4dacb04710e3548e6b8855f7e59af0dfb8f3b82474cd432775d4be558227feb98fc8e08123f683a5ae5cdece6dc743d815c642d766de9f20a6cb15e1477a8465983a0754979e2503d042a7cd33f033e1953269d7a97065d3db5d3d911bcc0d2bd32319f78ee0041c52f615adabb49b0b102379063778dc8a5ba8bc56072d65e84d843f9bdc423d5d31ec2aa9639fc656166d1dba52833e433a5e811195661c95509f7b6c0302264ef389b51b84779b0672e48687d16f7e8993be97f556eb3312f0093a4bb41f94c7071ea299e3bc5e7239377ae8dca699d3f42a2c6c22252b5afe5d785558660b5bec4e660a16816a053f78c52b8310dd021313c2c7531f0d7be4f07b5d15a7315c95487939c79062b08f6a6d99ccc3bd04e61c2e610e0cd13200e79a85110abc8c3b9891bfbbeb2d9223eac0ab2fe7d056d31eddcf4626374cb400e7127720df45582bdb9a32d30932cbbd5a28177604bb215eb331bd8d9274ebd162deacedce652d010b41e81dcceb2c7e533a68e0add3a2ef7a2ad1531e3acacd2a62065b44597c30e8e014f2fe5564fb30613f4fbf05ccee14a583ae1d32a61119b72b3ec399ac7daa6d1e9fadd3e50e9a0bb7d294a61efb34d6f31097cce42642b4cd236915677cda446b26ed71054e987ee77934da7f6b6ca811e5559274fb8f5fd5c211d0941eb711ead398b60a08148edc235ada9eff8dd2ccb146e32d6eac81969b64e38fcaa55dbb56f54b429b7cea221755b5721d465afa86c465e77670c48585191d4fa767bb73b9312bb409af44294169e935527a155e9b4b08caabf19c8e0177329995640bbb1d81771e5062ba1ec1c35ce1a43ecfef1bf31eeb486c6dab49cf91852b3244498516cf90bf19b73b36c778b216299f9d043482bed8cbe4f26bcebaf3f78ad64f1b7ebddcf98316e21793488c5dd7fc9ff19f53fbaf4dbe7df9c93b220da9f3b1132e66abea72b2dfa1e80f11371997fe1b5f6a5255255fa9f9e9c1b6a12dc98dc821ce87f320ed7308ab8dcccba855b2af0da552e7418bfad6080d0000000568ae13c79c237b6a2da82e611cc06b8ce3761eb560e0068ce3307e82dd642e9df2d86fef834bb4770a23409271a322bb3246f2db6f9dd25a71716b3979e46006bda69b10ffddacc40832f386a0b98396c261b1b93002918216616b2fc320151349f0542189630e86e0ae99cbf9b6966d3aac2dd2bf75bd06c8b0aa9312fadb028cb37125de454bef96259821971e6dc89017adc9f40daa806336a6b50d9f09f3023f00`

	pf, err := appName.VerifyTxOutProof(txhex)
	if err != nil {
		fmt.Printf("Code: %v\n", pf.Error.Code)
		fmt.Printf("Message: %v\n\n", pf.Error.Message)
		log.Fatalln("Err happened", err)
	}

	fmt.Println("pf value", pf)
	fmt.Println("-------")
	fmt.Println(pf.Result)
	fmt.Println("-------")

	for i, v := range pf.Result {
		fmt.Println(i)
		fmt.Println(v)
	}
}