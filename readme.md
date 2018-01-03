## guitocons :
### attach the GUI stdout & stderr outputs to the calling process console (if there is one)



**install with :** `go get \github.com\ffred\guitocons`

On Windows if you build your Go project with '-ldflags="-H windowsgui' parameters, you got a "GUI" program, with no console.
even if you launch your program from a console, you won't have any output on it from 'fmt' or 'log' functions.
I tried a few GUI samples on Go witch was outputting values on the console and wasn't showing any of it for me on Windows, so I search a way to change that.

this small package attach the parent process console to your GUI program, so you can still output logs during development or whatever need you could have with your final program.

_I'm rather new to Go programming, so there's probably better or different ways to do that, but for now its working... :-)_

there's not much informations with Go language about that GUI/console "problem" online. I found what I needed [here.](https://stackoverflow.com/questions/23743217/printing-output-to-a-command-window-when-golang-application-is-compiled-with-ld/23744350) 


a simple way to test it is by building this small code with '-ldflags="-H windowsgui' parameters and launch it from a console, or not :

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ffred/guitocons"
)

func main() {
	err1 := guitocons.Guitocons()
	if err1 == nil {
		log.Println("err console OK..")
		fmt.Println("std console OK..")
		os.Stderr.WriteString("Msg to STDERR\n")

		var s string
		fmt.Scanln(&s)
		os.Exit(0)
	}
	log.Println(err1)
}
```
