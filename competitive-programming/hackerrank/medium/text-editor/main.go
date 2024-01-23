package main

type abc string

func (a abc) append(s string) abc {
	return a + abc(s)
}

func (a abc) delete(n int) abc {
	return a[:len(a)-n]
}

func history(operations []string, operation string) []string {
	operations = append(operations, operation)
	return operations
}

func (a abc) print(k int) string {
	return string(a[k-1])
}

func main() {

}
