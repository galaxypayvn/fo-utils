package utarray

/*
ArrayToMap is a function returns a map
with keys representing the unique elements of the slice, and values being empty structs.

Example:

	mySlice := []int{1, 2, 3, 3, 4, 5, 5}
	uniqueMap := ArrayToMap(mySlice)
	fmt.Println(uniqueMap) // Output: map[1:{} 2:{} 3:{} 4:{} 5:{}]
*/
func ArrayToMap[T comparable](arr []T) map[T]struct{} {
	result := make(map[T]struct{}, len(arr))
	for _, elem := range arr {
		result[elem] = struct{}{}
	}
	return result
}

/*
Unique is a function returns a new slice with duplicates removed.

Example:

myArray := []int{1, 3, 3, 3, 4}

result := Unique(myArray) // [1, 3, 4]
*/
func Unique[T comparable](arr []T) []T {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]T, 0, len(arr))
	seen := make(map[T]bool)

	for _, v := range arr {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

// Contains checks whether the value is in the slice or not.
// It returns true if the value is found in the slice, false otherwise.
//
// Example:
//
//	slice := []int{1, 2, 3, 4, 5}
//	fmt.Println(Contains(slice, 3)) // true
//	fmt.Println(Contains(slice, 6)) // false
func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
