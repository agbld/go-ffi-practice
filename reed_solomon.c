#include <stdio.h>

int shift = 1;

void set_shift(int s) {
    shift = s;
}

// A simple function that simulates encoding
void encode(const char* input, char* output) {
    while (*input) {
        *output = *input + shift; // Dummy encoding logic
        input++;
        output++;
    }
    *output = '\0';
}

// A simple function that simulates decoding
void decode(const char* input, char* output) {
    while (*input) {
        *output = *input - shift; // Dummy decoding logic
        input++;
        output++;
    }
    *output = '\0';
}
