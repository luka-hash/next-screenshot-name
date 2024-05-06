// Copyright © 2024- Luka Ivanović
// This code is licenced under the terms of the MIT licence (see LICENCE for details).

package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

func main() {
	ss_home, ok := os.LookupEnv("XDG_SCREENSHOTS_HOME")
	if !ok {
		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalln(err)
		}
		ss_home = path.Join(home, "Pictures/")
	}
	current_time := time.Now().Format("2006-01-2T15:04:05.000Z07:00")
	fmt.Print(path.Join(ss_home, current_time+".png"))
}
