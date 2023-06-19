package hash

/*
#cgo CFLAGS: -I../lib  -mavx2
#cgo LDFLAGS: -L../lib  -lrt -lm -ldl -l quant_hash
#include "hash.h"
*/
import "C"
import (
	"context"
	"encoding/hex"
	"fmt"
	"time"
	"unsafe"
)

var isSupportHashCard bool

const ContextsMaxLen = 1024

func reverse(s []byte) []byte {
	var ss = make([]byte, len(s), len(s))
	copy(ss, s)
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return ss
}

func GetP(ctx context.Context) []byte {
	var (
	// r16 []byte
	// err error
	)
	// if r16, err = GetRng(ctx, 16); err != nil {
	// 	fmt.Println("获取随机数r16失败. err=%v", err)
	// 	return nil
	// }

	temp := [16]byte{0x95, 0x9d, 0xf1, 0x41, 0x16, 0x22, 0x4c, 0xfa, 0x75, 0xeb, 0x51, 0xac, 0x86, 0x47, 0xa7, 0x60} //0x959df14116224cfa75eb51ac8647a761
	// temp := [16]byte{0xFF, 0x4B, 0xBE, 0xB2, 0x53, 0x3E, 0xF4, 0xAE, 0xF5, 0x4B, 0xFA, 0x26, 0x54, 0xFF, 0x3C, 0x0A} //0xFF4BBEB2533EF4AEF54BFA2654FF3C0B
	// temp := [16]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01, 0x86, 0x96}                                        //0x018697
	r16 := reverse(temp[:])

	zero := [16]byte{}
	r32 := append(r16[:], zero[:]...)
	fmt.Printf("计算不可约的输入随机数=", hex.EncodeToString(reverse(r32[:])))
	t1 := time.Now()
	var p [32]byte
	C.GetP((*C.char)(unsafe.Pointer(&r32[0])), (*C.char)(unsafe.Pointer(&p[0])))
	t2 := time.Since(t1)
	p16 := reverse(p[:16])
	fmt.Printf("得到不可约多项式=", hex.EncodeToString(p16), "time=", t2)
	return p16
}

// hash的长度为16字节
func GetHash(ctx context.Context, p, s, buf []byte, len uint64) []byte {
	return GetHashSw(ctx, p, s, buf, len)
}
func GetHashSw(ctx context.Context, p, s, buf []byte, len uint64) []byte {
	p16 := reverse(p[:16])
	var hash [16]byte
	C.GetHash((*C.char)(unsafe.Pointer(&s[0])),
		(*C.char)(unsafe.Pointer(&p16[0])),
		(*C.char)(unsafe.Pointer(&buf[0])),
		(C.ulong)(len),
		(*C.char)(unsafe.Pointer(&hash[0])))
	h16 := reverse(hash[:16])
	// g.Log().Debugf(ctx, "h16=", hex.EncodeToString(h16[:]))
	return h16
}
