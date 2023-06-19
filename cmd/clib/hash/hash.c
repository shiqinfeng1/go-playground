
#include <stdio.h>
#include <stdint.h>
#include "hash.h"

void GetP(char* pRand, char* outP)
{
	Init();
	pRand[0] |= 1;
	pRand[16] = 1;
    int32_t iRet = GetGF(pRand, outP);
    if (0 != iRet)
    {
        printf("get p fail. ret=%d:\n",iRet);
        return;
    }
	printf("get p ok.\n");
	return;
}

// 计算hash值
void GetHash(char *pRand, char *P, char *pBuf, uint64_t iLen, char *outHash)
{
	Init();
    QuantumHashCommon(pBuf, iLen, pRand, P, outHash);
	return;
}

