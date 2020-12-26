package pkg

func privateFunc() string {
	return ", its okay!"
}

func Greeting(name string) (string, string) {
	say := "Hello " + name
	return say, privateFunc()
}
