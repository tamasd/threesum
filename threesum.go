package threesum

func naive(input []int) bool {
	for i := range input {
		for j := range input {
			for k := range input {
				if i != j && i != k && j != k && input[i]+input[j]+input[k] == 0 {
					return true
				}
			}
		}
	}

	return false
}

func hash(input []int) bool {
	n := len(input)
	set := make(map[int]struct{}, n)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x := -(input[i] + input[j])
			if _, ok := set[x]; ok {
				return true
			}
			set[input[j]] = struct{}{}
		}

		// This will be recognized by the compiler and replaced with a
		// mapclear() call.
		for k := range set {
			delete(set, k)
		}
	}

	return false
}

func twopointer(input []int) bool {
	n := len(input)

	for i := 0; i < n-1; i++ {
		left := i + 1
		right := n - 1
		x := input[i]

		for left < right {
			sum := x + input[left] + input[right]
			if sum == 0 {
				return true
			}

			if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return false
}

func binarysearch(input []int) bool {
	n := len(input)

	for i := 0; i < n-2; i++ {
		left := bs(input, i+1, n-2, -(input[i] + input[n-1]))
		right := bs(input, left+1, n-1, -(input[i] + input[left]))
		for left < right && right < n {
			sum := input[i] + input[left] + input[right]
			if sum == 0 {
				return true
			}

			if sum < 0 {
				left++
			} else {
				right--
			}
		}
		if input[i] == 0 {
			break
		}
	}

	return false
}

func bs(input []int, min, max, target int) int {
	for {
		if max < min {
			return min
		}

		mid := (min + max) / 2
		if input[mid] < target {
			min = mid + 1
		} else if input[mid] > target {
			max = mid - 1
		} else {
			return mid
		}
	}
}
