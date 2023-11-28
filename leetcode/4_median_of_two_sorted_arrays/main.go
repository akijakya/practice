package main

func main() {
	nums1 := []int{1, 2, 4}
	nums2 := []int{1, 3, 4}

	println(findMedianSortedArrays(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	lenCombined := int(len(nums1)) + int(len(nums2))
	repeat := (lenCombined / 2) + 1

	last := -10 ^ 6
	beforeLast := -10 ^ 6
	i1 := 0
	i2 := 0

	for i := 0; i < repeat; i++ {
		if i2 == len(nums2) && last <= nums1[i1] {
			beforeLast = last
			last = nums1[i1]
			i1++
			continue
		}

		if i1 == len(nums1) && last <= nums2[i2] {
			beforeLast = last
			last = nums2[i2]
			i2++
			continue
		}

		if last <= nums1[i1] && nums1[i1] <= nums2[i2] {
			beforeLast = last
			last = nums1[i1]
			i1++
			continue
		}

		if last <= nums2[i2] && nums2[i2] <= nums1[i1] {
			beforeLast = last
			last = nums2[i2]
			i2++
		}
	}

	if lenCombined%2 == 0 {
		return (float64(last) + float64(beforeLast)) / float64(2)
	}

	return float64(last)
}
