package main
const N = 16
const alph = "0|"

func main() {
	t := new_thorium_system("O|")
	print_key(t.hard_key)
	for i := 0 ; i < 8 ; i++ {
		p := random_text(16)
		c := t.encrypt(p)
		print("f( ",str(p)," ) = ",str(c), "\n")
		d := t.decrypt(c)
		for i := range p { if p[i] != d[i] { print("ERROR ") } }
	}
}