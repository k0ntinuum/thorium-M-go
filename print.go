package main
import "fmt"
func print_row(r [N]int) {
	for c := range r { fmt.Printf("%c ", alph[r[c]]) } 
	fmt.Println()
}

func print_vec(r []int) {
	for c := range r { fmt.Printf("%c", alph[r[c]]) } 
	fmt.Println()
}
func str(v []int) string {
	s := ""
	for i := range v { s += fmt.Sprintf("%c", alph[v[i]]) }
	return s
}

func print_key(k [N][N]int ) {
	for r := range k { print_row(k[r]) }
	fmt.Println()
}