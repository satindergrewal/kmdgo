#include <stdio.h>
#include <stdbool.h>

#include "saplinglib.h"

int main() {
	bool nohd = false;
	int zcount = 1;
	char *seed = "user specified seed phrase";
	bool iguana_seed = true;

	char * from_rust = rust_generate_wallet(nohd, zcount, seed, iguana_seed);
	char *stri = from_rust;
	printf("%s", stri);
	rust_free_string(from_rust);

	return 0;
}