package main // karena didalam folder root atau sesuai nama didalam direktori atau sesuai func atau juga sesaui dengan kebutuhan aja

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
	// BanyakGenre()
	horror.KategoriHorror()

}
