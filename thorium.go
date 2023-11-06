package main
import "math/rand"

type Thorium struct  {
	hard_key [N][N]int
	soft_key [N][N]int
	tmp [N]int
}

func new_thorium_system() Thorium {
	var t Thorium;
	t.randomize_hard_key()
	t.soft_key = t.hard_key
	return t
}

func (T *Thorium) randomize_hard_key() {
	for r := range T.hard_key { for c := range T.hard_key[r] { T.hard_key[r][c] = rand.Intn(2)} }
}

func (T *Thorium) spin_row(r int, a int) {
	for c := range T.soft_key[r] { T.tmp[(c+a)%N] = T.soft_key[r][c] }
	for c := range T.soft_key[r] { T.soft_key[r][c] = T.tmp[c] }
}
func (T *Thorium) spin_col(c int, a int) {
	for r := range T.soft_key[c] { T.tmp[(r+a)%N] = T.soft_key[r][c] }
	for r := range T.soft_key[c] { T.soft_key[r][c] = T.tmp[r] }
}

func trace(M [N][N]int) int {
	T := 0; 
	for i := 0 ; i < N ; i++ { T += M[i][i] }
	return T
}

func random_text(n int ) []int {
	t := make([]int,n)
	for i := range t {  t[i] = rand.Intn(2)} 
	return t
}