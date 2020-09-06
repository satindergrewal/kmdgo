#include <stdbool.h>

char * rust_generate_wallet(unsigned int count, const char* entropy);
void   rust_free_string(char* s);