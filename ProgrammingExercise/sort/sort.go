package sort

// BubbleSort 1. 冒泡排序 (Bubble Sort)
// 时间复杂度：O(n²) 最好O(n) 空间复杂度：O(1)
// 通过相邻元素比较和交换，将最大元素"冒泡"到数组末尾
func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false // 优化：如果本轮无交换说明已有序
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

// SelectionSort 2. 选择排序 (Selection Sort)
// 时间复杂度：O(n²) 空间复杂度：O(1)
// 每次从未排序部分选择最小元素放到已排序末尾
func SelectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

// InsertionSort 3. 插入排序 (Insertion Sort)
// 时间复杂度：O(n²) 最好O(n) 空间复杂度：O(1)
// 将未排序元素逐个插入到已排序部分的合适位置
func InsertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// 将大于key的元素后移
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// QuickSort 4. 快速排序 (Quick Sort)
// 时间复杂度：平均O(n log n) 最差O(n²) 空间复杂度：O(log n)
// 分治策略：选择基准，将小元素放左边，大元素放右边，递归排序
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2] // 选择中间元素作为基准
	var left, right, equal []int

	for _, num := range arr {
		switch {
		case num < pivot:
			left = append(left, num)
		case num == pivot:
			equal = append(equal, num)
		default:
			right = append(right, num)
		}
	}

	return append(append(QuickSort(left), equal...), QuickSort(right)...)
}

// MergeSort 5. 归并排序 (Merge Sort)
// 时间复杂度：O(n log n) 空间复杂度：O(n)
// 分治策略：将数组分成两半分别排序，然后合并有序数组
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
