
GLOBL ·S(SB),$24
DATA ·S+0(SB)/8,$·S+16(SB)
DATA ·S+8(SB)/8,$8
DATA ·S+16(SB)/8,$"HELLO WD"


TEXT ·main(SB), $16-0
    MOVQ ·S+0(SB),AX
    MOVQ AX,0(SP)
    MOVQ ·S+8(SB),BX
    MOVQ BX,8(SP)
    CALL ·runtime·printstring(SB)
    RET
