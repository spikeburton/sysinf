package main

import (
	"fmt"

	sys "golang.org/x/sys/unix"
)

func main() {
	ver, _ := sys.Sysctl("kern.version")
	cpu, _ := sys.Sysctl("machdep.cpu.brand_string")
	cores, _ := sys.SysctlUint32("machdep.cpu.core_count")

	fmt.Printf("System Information\n---\n\n")
	fmt.Printf("Kernel: \t%s\n", ver)
	fmt.Printf("CPU: \t\t%s\n", cpu)
	fmt.Printf("Cores: \t\t%v\n", cores)
}
