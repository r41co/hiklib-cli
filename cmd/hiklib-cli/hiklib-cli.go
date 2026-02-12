package main

import "github.com/r41co/hiklib-cli/cmd/hiklib-cli/pkg/cmd"

func main() {
	cmd.Execute()
	// log.Printf("hiklib-cli [%s]\n", hiklib.HikVersion())
	// u, dev := hiklib.HikLogin("10.80.2.221", 8000, "admin", "251120oO")
	// // while True {
	// hiklib.HikCaptureImage(u, dev.ByStartChan, "./cam.jpeg")
	// }
}
