package collections

func Pop[K Key, V Value](_map map[K]V) map[K]V {
	popped := make(map[K]V)

	index := -1
	for key, value := range _map {
		index++
		if index != 0 {
			popped[key] = value
		}
	}

	return popped
}

func Peek[K Key, V Value](_map map[K]V) K {
	for key, _ := range _map {
		return key
	}

	// never reacher
	var empty K
	return empty
}

func HasKey[K Key, V Value](_map map[K]V, searchedKey K) bool {
	for key := range _map {
		if key == searchedKey {
			return true
		}
	}

	return false
}

func HasValue[K Key, V Value](_map map[K]V, searchedValue V) bool {
	for _, value := range _map {
		if value == searchedValue {
			return true
		}
	}

	return false
}

func Sum[K Key, V NumberValue](_map map[K]V) V {
	var sum V
	sum = 0

	for _, element := range _map {
		sum += element
	}

	return sum
}
