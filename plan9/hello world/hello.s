#include "textflag.h"


//GLOBL ·T(SB),NOPTR,$8
//DATA ·T+0(SB)/8,$"hello wd"

TEXT ·Hello(SB),$8-0
	MOVQ $8,s+8(FP)
	MOVQ $0x4457004f4c4c4548,AX
	MOVQ AX,tmp+0(SP)
	LEAQ tmp+0(SP),AX

	MOVQ AX,s+0(FP)
	RET
