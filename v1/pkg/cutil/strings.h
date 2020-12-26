#ifndef __GO_C_UTILS_STRING__
#define __GO_C_UTILS_STRING__

static inline char** createStringArray(int len) {
  return (char**) malloc(sizeof(char*) * len);
}

static inline void putString(char** arr, int ind, char* value) {
  arr[ind] = value;
}

#endif // __GO_C_UTILS_STRING__