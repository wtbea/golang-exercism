package lasagna

func PreparationTime(l []string, t int) int {
	if t == 0 {
		t = 2
	}
	return len(l) * t
}

func Quantities(l []string) (int, float64) {
	n := 0
	s := 0.0
	for _, v := range l {
		if v == "noodles" {
			n += 50
		} else if v == "sauce" {
			s += 0.2
		}
	}
	return n, s
}

func AddSecretIngredient(i1 []string, i2 []string) {
	i2[len(i2)-1] = i1[len(i1)-1]
}

func ScaleRecipe(a []float64, p int) []float64 {
	r := make([]float64, len(a))
	for i, v := range a {
		r[i] = v * float64(p) / 2
	}
	return r
}
