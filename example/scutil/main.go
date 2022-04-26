package main

import (
	"fmt"
	"github.com/NuovoVita/scutil"
)

func main() {
	dig := 12344484.66
	str := "12,344,484.66"

	//12,344,484.66
	fmt.Println(scutil.Grading(dig, ',', 3))

	//12344484.66
	fmt.Printf("%.2f\n", scutil.CancelGrading(str, ','))
}
