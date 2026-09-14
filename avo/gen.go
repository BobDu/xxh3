package main

import (
	"flag"

	"github.com/mmcloughlin/avo/build"
	"github.com/mmcloughlin/avo/ir"
	"github.com/mmcloughlin/avo/operand"
)

//go:generate go run . -avx -out ../accum_vector_avx_amd64.s -pkg xxh3
//go:generate go run . -avx512 -out ../accum_vector_avx512_amd64.s -pkg xxh3
//go:generate go run . -sse -out ../accum_vector_sse_amd64.s -pkg xxh3

var (
	avx    = flag.Bool("avx", false, "run avx generation")
	avx512 = flag.Bool("avx512", false, "run avx512 generation")
	sse    = flag.Bool("sse", false, "run sse generation")
)

func main() {
	flag.Parse()

	if *avx {
		AVX()
	} else if *sse {
		SSE()
	} else if *avx512 {
		AVX512()
	}
}

// align64 starts the current function on a cache line; Go aligns functions to
// 32 bytes and avo has no PCALIGN of its own.
func align64() {
	build.Instruction(&ir.Instruction{Opcode: "PCALIGN", Operands: []operand.Op{operand.Imm(64)}})
}
