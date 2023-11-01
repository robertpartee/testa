package testa

func GetTestAName(nCode int) (NameOutput string) {

	switch nCode {
	case 1:
		NameOutput = "Name-a"
	case 2:
		NameOutput = "Name-b"
	case 3:
		NameOutput = "Name-c"
	default:
		NameOutput = "Name-zeta"
	}
	return NameOutput
}
