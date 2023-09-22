#include "callback.h"

int32_t callCallback(callbackFunc fn, int32_t value) {
    return fn(value);
}
