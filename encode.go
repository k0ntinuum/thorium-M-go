package main


func (t *Thorium) encode(p []int) []int {
	//t.soft_key = t.hard_key
	c := make([]int, len(p))
	for i := range p {
		c[i] = (trace(t.soft_key) + p[i] ) % 2 
		if p[i] == 0 { 
			t.spin_row(i%N, t.soft_key[i%N][i%N] + 1)
		} else {
			t.spin_col(i%N, t.soft_key[i%N][i%N] + 1)
		}
	}
	return c
}
func (t *Thorium) decode(c []int) []int {
	//t.soft_key = t.hard_key
	p := make([]int, len(c))
	for i := range p {
		p[i] = (trace(t.soft_key) + 2 - c[i] ) % 2 
		if p[i] == 0 { 
			t.spin_row(i%N, t.soft_key[i%N][i%N] + 1)
		} else {
			t.spin_col(i%N, t.soft_key[i%N][i%N] + 1)
		}
	}
	return p
}
