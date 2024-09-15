package arrays

func Merge(nums1 []int, m int, nums2 []int, n int) {
    pointerM := m - 1
    pointerN := n - 1
    pointer  := m + n - 1

    for pointerM >= 0 && pointerN >= 0 {
        if nums1[pointerM] > nums2[pointerN] {
            nums1[pointer] = nums1[pointerM]
            pointerM--
        } else {
            nums1[pointer] = nums2[pointerN]
            pointerN--
        }
        pointer--
    }

    for pointerN >= 0 {
        nums1[pointer] = nums2[pointerN]
        pointerN--
        pointer--
    }
}