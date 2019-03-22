package main

func main() {
	numterms, sum := add(1, 4, 5, 6, 8)
	println("Added: ", numterms, "to create a total of ", sum)

}

func add(terms ...int) (numterms int, sum int) {
	for _, v := range terms {
		sum += v
	}
	numterms = len(terms)
	return
}
