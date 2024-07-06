#include <stdio.h>

// A simple function that simulates encoding
void encode(const char* input, char* output) {
    while (*input) {
        *output = *input + 1; // Dummy encoding logic
        input++;
        output++;
    }
    *output = '\0';
}

// A simple function that simulates decoding
void decode(const char* input, char* output) {
    while (*input) {
        *output = *input - 1; // Dummy decoding logic
        input++;
        output++;
    }
    *output = '\0';
}
