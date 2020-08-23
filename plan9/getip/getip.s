

TEXT ·getip(SB),$0x8-0
	MOVQ ip-8(FP) ,AX
	MOVQ AX,ret+0(FP)
	RET

