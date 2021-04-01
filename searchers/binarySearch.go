package searchers

// 二分查找算法之递归算法
func binarySearchRecursion(arr []int, target int, lowIndex int, highIndex int) int {
	// 先判断数组是否不为空 和索引大小值判断
	if len(arr) == 0 || highIndex < lowIndex {
		return -1
	}

	mid := (lowIndex + highIndex) / 2
	if arr[mid] > target {
		return binarySearchRecursion(arr, target, lowIndex, mid - 1)
	} else if arr[mid] < target {
		return binarySearchRecursion(arr, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// 二分查找之迭代算法
func binarySearchIteration(arr []int, target int, lowIndex int, highIndex int) int {
	// 先判断数组是否不为空 和索引大小值判断
	if len(arr) == 0  || highIndex > len(arr) || lowIndex < 0{
		return -1
	}

	for lowIndex < highIndex {
		midIndex := (lowIndex + highIndex) / 2
		if arr[midIndex] < target {
			highIndex = midIndex - 1
		} else if arr[midIndex] > target {
			lowIndex = midIndex + 1
		} else {
			return midIndex
		}
	}

	return -1
}
