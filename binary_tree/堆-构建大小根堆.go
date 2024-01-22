package binary_tree

import "fmt"

// minHead 构建小根堆
func minHead(root int, end int, arr []int) {
	for {
		child := 2*root + 1
		if child >= end {
			break
		}

		// 判断该root的右孩子是否存在，如果存在，则跟其左孩子比较，去最小的那一个
		if child+1 < end && arr[child] > arr[child+1] {
			child = child + 1
		}

		// 判断root节点和child节点的大小
		if arr[root] > arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}

func maxHead(root int, end int, arr []int) {
	for {
		child := 2*root + 1
		if child >= end {
			break
		}

		if child+1 < end && arr[child] < arr[child+1] {
			child = child + 1
		}

		// 判断root节点和child节点的大小
		if arr[root] < arr[child] {
			arr[root], arr[child] = arr[child], arr[root]
			root = child
		} else {
			break
		}
	}
}

// createHead 创建堆
func createMinHead(arr []int, size int) {
	fmt.Println("开始创建小根堆")
	// 建小根堆
	for root := size / 2; root >= 0; root-- {
		minHead(root, size, arr)
	}
	fmt.Println("小根堆建立完成", arr)
}

func createMaxHead(arr []int, size int) {
	fmt.Println("开始创建大根堆")
	// 建小根堆
	for root := size / 2; root >= 0; root-- {
		maxHead(root, size, arr)
	}
	fmt.Println("大根堆建立完成", arr)
}
