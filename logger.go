package worker

type Logger interface {
	Println(s string)
	Printf(s string, args ...interface{})
}
