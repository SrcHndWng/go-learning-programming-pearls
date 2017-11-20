package main

type searchRange struct {
	start int
	end   int
}

var ranges []searchRange

func binarySearch(target int, nums []int) (int, int) {
	return search(target, 0, len(nums)-1, nums)
}

func search(target int, start int, end int, nums []int) (int, int) {
	if searched(start, end) {
		return -1, -1
	}
	ranges = append(ranges, searchRange{start, end})
	i := start + binary(start, end)
	if i > len(nums)-1 {
		return -1, -1
	}
	n := nums[i]
	if n < target {
		return search(target, i, end, nums)
	} else if n > target {
		return search(target, start, i, nums)
	}
	return i, nums[i]
}

func searched(start int, end int) bool {
	r := searchRange{start, end}
	if len(ranges) == 0 {
		return false
	}
	last := ranges[len(ranges)-1]
	return r.start == last.start && r.end == last.end
}

func binary(start int, end int) int {
	result := (end - start) / 2
	if result == 0 && start > 0 {
		result = 1
	}
	return result
}
