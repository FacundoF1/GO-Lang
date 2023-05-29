package examples

func Add(num1, num2 int) int {
	return num1 + num2
}

type Dependency interface {
	Write(p []byte) (n int, err error)
}

func DependencyUser(writer Dependency) {
	writer.Write([]byte("a string"))
}
