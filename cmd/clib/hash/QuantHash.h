/**
 * @company          MatrixTime[All Right Reserved]
 * @brief            基础运算公共接口定义
 * @author           gyang
 * @email            gyang@jmail.com
 * @version          1.0.0
 * @date:            2022-01-10 17:44:09
 *
 * Revison History:
 * Author      Date(YYYY-MM-DD)    Version     Description of changes
 * gyang        2022-01-10          1.0.0       创建文件
 * ------  ----------------  -------  ----------------------
 *
 */

#ifndef __COMMON_BASE_H__
#define __COMMON_BASE_H__

#include <immintrin.h>
#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

#define ALIGNED(a) __attribute__((__aligned__(a)))
#define MAX_U64 0XFFFFFFFFFFFFFFFF


// 128Bit 64位以内的左移
#define SHL128LT64(v, n) \
    ({ \
     __m128i v1, v2; \
     v1 = _mm_slli_epi64(v, n); \
     v2 = _mm_slli_si128(v, 8); \
     v2 = _mm_srli_epi64(v2, 64 - (n)); \
     v1 = _mm_or_si128(v1, v2); \
     v1; \
     })

// 128bit 左移n位
#define SHL128(v, n) \
    ({ \
     __m128i v1, v2; \
     \
     if ((n) >= 64) \
     { \
     v1 = _mm_slli_si128(v, 8); \
     v1 = _mm_slli_epi64(v1, (n) - 64); \
     } \
     else \
     { \
     v1 = _mm_slli_epi64(v, n); \
     v2 = _mm_slli_si128(v, 8); \
     v2 = _mm_srli_epi64(v2, 64 - (n)); \
     v1 = _mm_or_si128(v1, v2); \
     } \
     v1; \
     })

// 128Bit 64位以内的右移
#define SHR128LT64(v, n) \
    ({ \
     __m128i v1, v2; \
     v1 = _mm_srli_epi64(v, n); \
     v2 = _mm_srli_si128(v, 8); \
     v2 = _mm_slli_epi64(v2, 64 - (n)); \
     v1 = _mm_or_si128(v1, v2); \
     v1; \
     })

// 128bit 右移n位
#define SHR128(v, n) \
    ({ \
     __m128i v1, v2; \
     \
     if ((n) >= 64) \
     { \
     v1 = _mm_srli_si128(v, 8); \
     v1 = _mm_srli_epi64(v1, (n) - 64); \
     } \
     else \
     { \
     v1 = _mm_srli_epi64(v, n); \
     v2 = _mm_srli_si128(v, 8); \
     v2 = _mm_slli_epi64(v2, 64 - (n)); \
     v1 = _mm_or_si128(v1, v2); \
     } \
     v1; \
     })

/** 
 * @brief           定义128bit联合体
 *                      
 */
typedef union {
        int8_t              m128i_i8[16];
        int16_t             m128i_i16[8];
        int32_t             m128i_i32[4];    
        int64_t             m128i_i64[2];
        uint8_t             m128i_u8[16];
        uint16_t            m128i_u16[8];
        uint32_t            m128i_u32[4];
        uint64_t            m128i_u64[2];
        __m128i             m128i;
        //} ALIGNED xmm_t;
} xmm_t;

/** 
 * @brief           定义256bit联合体
 *                      
 */
typedef union
{
    int8_t              i8[32];
    int16_t             i16[16];
    int32_t             i32[8];
    int64_t             i64[4];
    uint8_t             u8[32];
    uint16_t            u16[16];
    uint32_t            u32[8];
    uint64_t            u64[4];
    xmm_t               m128i[2]; 
    __m256i             u256;
} ymm_t;

typedef struct _BIT8_
{                                                                                                                                                                          
    uint8_t b0:1;                                                                                                                                                          
    uint8_t b1:1;                                                                                                                                                          
    uint8_t b2:1;                                                                                                                                                          
    uint8_t b3:1;                                                                                                                                                          
    uint8_t b4:1;                                                                                                                                                          
    uint8_t b5:1;                                                                                                                                                          
    uint8_t b6:1;                                                                                                                                                          
    uint8_t b7:1;                                                                                                                                                          
}Bit8;

// 定义avx指令，主要解决的是低版本gcc没实现这几个接口
inline __m256i _mm256_loadu2_m128ic (__m128i const *__PH, __m128i const *__PL)
{
    return _mm256_insertf128_si256 (_mm256_castsi128_si256 (_mm_loadu_si128 (__PL)),
            _mm_loadu_si128 (__PH), 1);
}

inline __m256i _mm256_set_m128ic (__m128i __H, __m128i __L)
{
    return _mm256_insertf128_si256 (_mm256_castsi128_si256 (__L), __H, 1);
}


/* @brief   初始化
 *          使用前务必调用初始化
 */
void Init();

///* @brief   计算不可约多项式
// *
// * @param
// *      p   输入的随机数
// *      Gf  计算得到的不可约多项式
// *
// * @return
// *      -1 失败
// *      0  成功
// */
//int32_t GetGF(ymm_t &P, ymm_t &Gf);

/* @brief   计算不可约多项式
 *
 * @param
 *      pR256   输入的随机数
 *      pGf256  计算得到的不可约多项式
 *
 * @return
 *      -1 失败
 *      0  成功
 */
int32_t GetGF(char *pR256, char *pGf);

/**
 * @brief 量子Hash函数,C/C++版
 *
 * @param pMsg      待Hash的明文
 * @param uiLen     待Hash的明文长度
 * @param pRandom   量子随机数，16字节
 * @param p         不可约多项式, 16字节
 * @param pHash     Hash结果, 由外部分配内存(可以是栈也可以是堆变量), 16字节
 *
 * @return 
 *          -1      执行Hash失败，请检查参数是否合法
 *          0       成功，从Hash出参中那到最终的Hash值
 */
int32_t QuantumHash4C(char *pMsg, uint64_t uiLen, xmm_t *pRandom, xmm_t p, xmm_t *pHash);

/**
 * @brief 量子Hash函数,通用版
 *
 * @param pMsg      待Hash的明文
 * @param uiLen     待Hash的明文长度
 * @param pRandom   量子随机数，16字节
 * @param p         不可约多项式, 16字节
 * @param pHash     Hash结果, 由外部分配内存(可以是栈也可以是堆变量), 16字节
 *
 * @return 
 *          -1      执行Hash失败，请检查参数是否合法
 *          0       成功，从Hash出参中那到最终的Hash值
 */
int32_t QuantumHashCommon(char *pMsg, uint64_t uMsgLen, char *pRandom,
        char *p, char *pHash);

#ifdef __cplusplus
}
#endif

#endif
