package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	osType, _ := unix.Sysctl("kern.ostype")
	osRelease, _ := unix.Sysctl("kern.osrelease")
	cpuVendor, _ := unix.Sysctl("machdep.cpu.vendor")
	cpuBrand, _ := unix.Sysctl("machdep.cpu.brand_string")

	fmt.Println(osType)
	fmt.Println(osRelease)
	fmt.Println(cpuVendor)
	fmt.Println(cpuBrand)
}
