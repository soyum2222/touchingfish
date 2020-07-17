#include "textflag.h"

GLOBL ·data(SB),NOPTR,$8
DATA ·data+0(SB)/8,$"ABCDEFG"


GLOBL ·S(SB),$16
DATA ·S+0(SB)/8,$·data(SB)
DATA ·S+8(SB)/8,$8



