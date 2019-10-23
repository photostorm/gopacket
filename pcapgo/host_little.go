//+build 386 amd64 amd64p32 arm arm64 ppc64le mipsle mips64le mips64p32le

package pcapgo

func htons(val uint16) uint16 {
	return (val&0x00FF)<<8 | (val&0xFF00)>>8
}