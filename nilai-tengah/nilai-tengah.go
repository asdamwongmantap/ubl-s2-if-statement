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
	var firstInput, secondInput, thirdInput, median int

	fmt.Println("+===========================================+")
	fmt.Println("| NIM: 2111600827 ")
	fmt.Println("| Nama: Mohd Amiruddin Saddam ")
	fmt.Println("| Program untuk mencetak nilai tengah dari tiga buah nilai yang diinput ")
	fmt.Println("+===========================================+")
	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Printf("Masukkan Bilangan Pertama : ")
	fmt.Scanf("%d", &firstInput)
	fmt.Printf("Masukkan Bilangan Kedua : ")
	fmt.Scanf("%d", &secondInput)
	fmt.Printf("Masukkan Bilangan Ketiga : ")
	fmt.Scanf("%d", &thirdInput)

	//check bigger number
	if firstInput == secondInput || secondInput == thirdInput || firstInput == thirdInput || (firstInput == secondInput && secondInput == thirdInput && firstInput == thirdInput) {
		fmt.Println("Angka tidak boleh sama")
	}

	if (firstInput < secondInput && secondInput < thirdInput) || (firstInput > secondInput && secondInput > thirdInput) {
		median = secondInput
	} else if (firstInput < thirdInput && thirdInput < secondInput) || (firstInput > thirdInput && thirdInput > secondInput) {
		median = thirdInput
	} else {
		median = firstInput
	}

	fmt.Println(fmt.Sprintf("Nilai Tengahnya adalah %d ", median))

}
