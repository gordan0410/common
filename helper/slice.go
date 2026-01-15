package helper

// IntersectSlices returns a slice containing the common elements found in both input slices a and b.
func IntersectSlices[T comparable](a, b []T) []T {
	m := make(map[T]struct{}, len(a))
	for _, item := range a {
		m[item] = struct{}{}
	}

	result := make([]T, 0)
	for _, item := range b {
		if _, ok := m[item]; ok {
			result = append(result, item)
			delete(m, item)
		}
	}

	return result
}

func SplitToBatches[T any](items []T, batchSize int) [][]T {
	if batchSize <= 0 || len(items) == 0 {
		return nil
	}

	var batches [][]T
	for i := 0; i < len(items); i += batchSize {
		end := min(i+batchSize, len(items))

		batches = append(batches, items[i:end])
	}

	return batches
}

func SliceGroupBy[T any, K comparable](items []T, keyFunc func(T) K) map[K][]T {
	grouped := make(map[K][]T)
	for _, item := range items {
		key := keyFunc(item)
		grouped[key] = append(grouped[key], item)
	}
	return grouped
}

func SliceClone[T any](src []T) []T {
	dst := make([]T, len(src))
	copy(dst, src)
	return dst
}

// SliceContains 目前僅測試基礎型別
func SliceDistinct[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(input))

	for _, val := range input {
		if _, exists := seen[val]; !exists {
			seen[val] = struct{}{}
			result = append(result, val)
		}
	}

	return result
}

// InsertItemToSlice inserts an item at a specified index in a slice.
func InsertItemToSlice[T any](slice []T, item T, index int) []T {
	if index < 0 || index > len(slice) {
		return slice // out of range
	}

	slice = append(slice, item)          // extend
	copy(slice[index+1:], slice[index:]) // make space
	slice[index] = item                  // insert

	return slice
}

func FilterUnwantedValueForSlice[T comparable](slice []T, unwantedValue T) []T {
	idx := 0
	for _, v := range slice {
		if v != unwantedValue {
			slice[idx] = v
			idx++
		}
	}
	return slice[:idx]
}

func SliceMap[T any, R any](input []T, mapper func(T) R) []R {
	output := make([]R, len(input))
	for i, v := range input {
		output[i] = mapper(v)
	}
	return output
}

func SliceMapWithFilter[T any, R any](input []T, mapper func(T) (R, bool)) []R {
	var output []R
	for _, v := range input {
		if mapped, ok := mapper(v); ok {
			output = append(output, mapped)
		}
	}
	return output
}
