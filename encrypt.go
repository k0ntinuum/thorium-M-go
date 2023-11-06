package main
import "slices"


func (T *Thorium) encrypt(p []int) []int {
	c := p
	for i := 0; i < N; i++ {
		//T.soft_key = T.hard_key
		T.spin(i)
		c = T.encode(c)
		slices.Reverse(c)
	}
	return c
}

func (T *Thorium) decrypt(c []int) []int {
	p := c
	for i := 0; i < N; i++ {
		//T.soft_key = T.hard_key
		T.spin(N - 1 - i)
		slices.Reverse(p)
		p = T.decode(p)
		
	}
	return p
}

func (T *Thorium) spin(c int) {
	T.soft_key = T.hard_key
	for i := 0 ; i < N ; i++ {
		if T.hard_key[i][c] == 0 { 
			T.spin_row(i, T.hard_key[i][i] + 1 ) 
		} else {  
			T.spin_col(i, T.hard_key[i][i] + 1 )  
		}
	}
}
