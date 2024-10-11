package main

type config struct {
	nextLocationAreaURL     *string
	previousLocationAreaURL *string
}

func main() {

	var cfg config

	startRepl(&cfg)

}
