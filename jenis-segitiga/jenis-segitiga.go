// Golang program to illustrate the usage of

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some variables
	var firstInput, secondInput, thirdInput int
	var triangle string

	fmt.Println("+===========================================+")
	fmt.Println("| NIM: 2111600827 ")
	fmt.Println("| Nama: Mohd Amiruddin Saddam ")
	fmt.Println("| Program untuk mencetak jenis segitiga dari tiga panjang sisi yang diinput ")
	fmt.Println("+===========================================+")
	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Printf("Masukkan Sisi Pertama : ")
	fmt.Scanf("%d", &firstInput)
	fmt.Printf("Masukkan Sisis Kedua : ")
	fmt.Scanf("%d", &secondInput)
	fmt.Printf("Masukkan Sisi Ketiga : ")
	fmt.Scanf("%d", &thirdInput)

	//check sisi
	if firstInput == secondInput && (secondInput == thirdInput && firstInput == thirdInput) {
		triangle = "Segitiga sama sisi"
	} else if firstInput == secondInput || secondInput == thirdInput || firstInput == thirdInput {
		triangle = "Segitiga sama kaki"
	} else {
		triangle = "Segitiga sembarang"
	}
	fmt.Println(fmt.Sprintf("Jenis Segitiga adalah %s ", triangle))

}
