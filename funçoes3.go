func conc_slices(x string) string {
	var concatenaçao string
	lista := strings.Split(x, " ")
	for _, palavra := range lista {
		concatenaçao += palavra
	}
	return concatenaçao
}
