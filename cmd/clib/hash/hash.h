
#include "mthash.h"
#include "QuantHash.h"

void GetP(char* pRand, char* outP);
void GetHash(char *pRand, char *P, char *pBuf, uint64_t iLen, char *outHash);
int InitHW(char* sover, char* xdmaver);
void GetHashHw(char *pRand, char *P, char *pBuf, uint64_t iLen, char *outHash);
void DeInit();