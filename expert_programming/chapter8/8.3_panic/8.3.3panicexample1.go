package panic

func compile() {
	panic("aa")
}

/*
go tool compile -S ./8.3.3anicexample1.go
"".compile STEXT size=66 args=0x0 locals=0x18 funcid=0x0
        0x0000 00000 (./8.3.3panicexample1.go:3)        TEXT    "".compile(SB), ABIInternal, $24-0
        0x0000 00000 (./8.3.3panicexample1.go:3)        MOVQ    (TLS), CX
        0x0009 00009 (./8.3.3panicexample1.go:3)        CMPQ    SP, 16(CX)
        0x000d 00013 (./8.3.3panicexample1.go:3)        PCDATA  $0, $-2
        0x000d 00013 (./8.3.3panicexample1.go:3)        JLS     58
        0x000f 00015 (./8.3.3panicexample1.go:3)        PCDATA  $0, $-1
        0x000f 00015 (./8.3.3panicexample1.go:3)        SUBQ    $24, SP
        0x0013 00019 (./8.3.3panicexample1.go:3)        MOVQ    BP, 16(SP)
        0x0018 00024 (./8.3.3panicexample1.go:3)        LEAQ    16(SP), BP
        0x001d 00029 (./8.3.3panicexample1.go:3)        FUNCDATA        $0, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (./8.3.3panicexample1.go:3)        FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x001d 00029 (./8.3.3panicexample1.go:4)        LEAQ    type.string(SB), AX
        0x0024 00036 (./8.3.3panicexample1.go:4)        MOVQ    AX, (SP)
        0x0028 00040 (./8.3.3panicexample1.go:4)        LEAQ    ""..stmp_0(SB), AX
        0x002f 00047 (./8.3.3panicexample1.go:4)        MOVQ    AX, 8(SP)
        0x0034 00052 (./8.3.3panicexample1.go:4)        PCDATA  $1, $0
        0x0034 00052 (./8.3.3panicexample1.go:4)        CALL    runtime.gopanic(SB)
        0x0039 00057 (./8.3.3panicexample1.go:4)        XCHGL   AX, AX
        0x003a 00058 (./8.3.3panicexample1.go:4)        NOP
        0x003a 00058 (./8.3.3panicexample1.go:3)        PCDATA  $1, $-1
        0x003a 00058 (./8.3.3panicexample1.go:3)        PCDATA  $0, $-2
        0x003a 00058 (./8.3.3panicexample1.go:3)        CALL    runtime.morestack_noctxt(SB)
        0x003f 00063 (./8.3.3panicexample1.go:3)        PCDATA  $0, $-1
        0x003f 00063 (./8.3.3panicexample1.go:3)        NOP
        0x0040 00064 (./8.3.3panicexample1.go:3)        JMP     0
        0x0000 65 48 8b 0c 25 00 00 00 00 48 3b 61 10 76 2b 48  eH..%....H;a.v+H
        0x0010 83 ec 18 48 89 6c 24 10 48 8d 6c 24 10 48 8d 05  ...H.l$.H.l$.H..
        0x0020 00 00 00 00 48 89 04 24 48 8d 05 00 00 00 00 48  ....H..$H......H
        0x0030 89 44 24 08 e8 00 00 00 00 90 e8 00 00 00 00 90  .D$.............
        0x0040 eb be                                            ..
        rel 2+0 t=25 type.string+0
        rel 5+4 t=17 TLS+0
        rel 32+4 t=16 type.string+0
        rel 43+4 t=16 ""..stmp_0+0
        rel 53+4 t=8 runtime.gopanic+0
        rel 59+4 t=8 runtime.morestack_noctxt+0
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 70 61 6e 69 63                                   panic
go.string."aa" SRODATA dupok size=2
        0x0000 61 61                                            aa
""..stmp_0 SRODATA static size=16
        0x0000 00 00 00 00 00 00 00 00 02 00 00 00 00 00 00 00  ................
        rel 0+8 t=1 go.string."aa"+0
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........
*/
