TEXT ·WithASM(SB),$0-16
  MOVQ    n+0(FP), AX
  MOVQ    AX, CX          // y ← x
  SARQ    $63, CX         // y ← y ⟫ 63
  XORQ    CX, AX          // x ← x ⨁ y
  SUBQ    CX, AX          // x ← x - y
  MOVQ    AX, ret+8(FP)
  RET
