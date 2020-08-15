#include "textflag.h"


GLOBL ·T(SB),NOPTR,$8
DATA ·T+0(SB)/8,$"hello wd"

TEXT ·Hello(SB),$0
	MOVQ $8,s+8(FP)
	LEAQ ·T(SB),AX
	MOVQ AX,s+0(FP)
	RET
