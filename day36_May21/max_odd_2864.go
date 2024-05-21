package main

func maximumOddBinaryNumber(s string) string {
	ones, zeros := make([]byte, 0), make([]byte, 0)
	for _, char := range s {
		if char == '1' {
			ones = append(ones, '1')
		} else {
			zeros = append(zeros, '0')
		}
	}
	if len(ones) == 1 {
		return string(zeros) + "1"
	}
	ones = ones[:len(ones)-1]
	return string(ones) + string(zeros) + "1"
}

func main() {

}
