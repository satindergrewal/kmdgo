#include <stdio.h>

#include "saplinglib.h"

int main() {
  char * from_rust = rust_generate_wallet(1,"user-provided-entropy");
  char *stri = from_rust;
  printf("%s", stri);
  rust_free_string(from_rust);

  return 0;
}