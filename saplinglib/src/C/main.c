#include <stdio.h>
#include <stdbool.h>

#include "saplinglib.h"

int main() {
	// rust_generate_wallet function takes four parameters
	// 1) nohd:				set it to false, if you don't want HD wallet
	// 2) zcount:			the number of sapling addresses you want to generate
	// 3) seed:				the user specified passphrase, which gives the same address everytime if given the same passphrase
	// 4) is_iguana_seed:	set this to true if you want the output to always give a deterministic address based on user specified seed phrase
	bool nohd = false;
	int zcount = 1;
	char *seed = "user specified seed phrase";
	bool is_iguana_seed = true;

	char * from_rust = rust_generate_wallet(nohd, zcount, seed, is_iguana_seed);
	char *stri = from_rust;
	printf("%s", stri);
	rust_free_string(from_rust);

	return 0;
}