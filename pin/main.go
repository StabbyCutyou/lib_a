package main

import "github.com/StabbyCutyou/pinner"

func main() {
	pinner.Register("github.com/StabbyCutyou/lib_c", "~> 1.0")
	pinner.Register("github.com/StabbyCutyou/lib_d", "= 1.2")

	pinner.RunMain()
}
