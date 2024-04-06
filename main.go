package main // karena didalam folder root atau sesuai nama didalam direktori atau sesuai func atau juga sesaui dengan kebutuhan aja
// pada sebuah project cuma ada satul file yang memakai package main

import (
	"fmt"

	"github.com/ardileon/go-structure-project/horror"
	"github.com/ardileon/go-structure-project/horror/dark"
	"github.com/ardileon/go-structure-project/tembak"
)

func main() {
	fmt.Println("Structure Project di dalam golang!")
	tembak.UcapFps()
	dark.SayHorrorDark()
	horror.SayHelloHorror()
	horror.KategoriHorror()
	// fmt.Println(BanyakGenre())
}
