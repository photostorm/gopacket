//+build armbe arm64be ppc64 mips mips64 mips64p32 ppc s390 s390x sparc sparc64

package pcapgo

func htons(val uint16) uint16 {
	return val
}