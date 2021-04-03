package responseBase

type Data struct {
	Success bool
	Loading bool
	Data    interface{}
	Message string
	Status  int
}
