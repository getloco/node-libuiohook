#include <stdint.h>

typedef int32_t (*callbackFunc)(int32_t);

int32_t callCallback(callbackFunc fn, int32_t value);
