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

// GetMissingElements returns a slice of elements from the first slice
// that are not present in the second slice.
//
// Example:
//
//	arr1 := []string{"a", "b", "c", "d"}
//	arr2 := []string{"b", "c", "e"}
//	missingElements := GetMissingElements(arr1, arr2)
//	println(missingElements) // Output: [a d]
func GetMissingElements[T comparable](arr1, arr2 []T) []T {
	// Create a map to store elements from arr2 for efficient lookup.
	arr2Map := make(map[T]struct{}, len(arr2))
	for _, elem := range arr2 {
		arr2Map[elem] = struct{}{}
	}

	// Iterate over arr1 and collect elements not present in arr2Map.
	var missingElements []T
	for _, elem := range arr1 {
		if _, exists := arr2Map[elem]; !exists {
			missingElements = append(missingElements, elem)
		}
	}

	return missingElements
}
