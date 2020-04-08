func parse(string string, varListP []string, varValsP []string) string {
	for i := 0; i < len(varListP); i++ {
		if string == varListP[i] {
			return varValsP[i]
		}
	}

	return string
}
