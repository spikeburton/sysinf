package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func main() {
	ver, _ := unix.Sysctl("kern.version")
	cpu, _ := unix.Sysctl("machdep.cpu.brand_string")

	fmt.Printf("System Information\n---\n\n")
	fmt.Printf("Kernel: \t%s\n", ver)
	fmt.Printf("CPU: \t\t%s\n", cpu)
}
