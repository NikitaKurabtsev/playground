package main

func getKeys[K comparable, V any](m map[K]V) []K {
	var keys []K

	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func main() {

}
