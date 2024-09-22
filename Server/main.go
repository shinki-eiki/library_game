package main

import (
	"patchouli_lib/controler"
)

func main() {
	// fmt.Println("Game Begin.")
	// h := wsConn.NewHub()
	// go wsConn.RunHub(h)
	controler.AllConnect()

	// com := &controler.CommandLine{}
	// for {
	// 	res, _ := com.AskIntegerRange("Select One:", 0, 5)
	// 	fmt.Println("receive ", res)
	// 	if res == 5 {
	// 		break
	// 	}
	// }

	// wsConn.Demo()
	// controler.MainLoop()
}
