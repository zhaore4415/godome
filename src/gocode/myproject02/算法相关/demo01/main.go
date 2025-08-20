package main

// 给你一个整数数组 nums 和一个整数 k，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

// 示例 1:
// 输入：nums = [1,1,1,2,2,3], k = 2
// 输出：[1,2]
// 示例 2:
// 输入：nums = [1], k = 1
// 输出：[1]
// 提示：
// - 1 <= nums.length <= 105
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
import (
	"fmt"
)

func topKFrequent(nums []int, k int) []int {
	// 步骤1: 统计频率
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}
	// fmt.Println(count) //如map[1:2 2:3 3:1]
	// 步骤2: 创建空桶，索引是频率
	n := len(nums)
	bucket := make([][]int, n+1) // bucket[freq] = 桶，切片的切片，列表的列表
	for i := range bucket {
		bucket[i] = []int{}
	}
	//fmt.Println(bucket)   //如[[] [] [] [] [] [] []]
	for num, freq := range count {
		bucket[freq] = append(bucket[freq], num)
	}
	// fmt.Println(bucket) //如[[] [3] [1] [2] [] [] []]
	// 步骤3: 从高频率到低频率收集结果
	var result []int
	for freq := n; freq >= 1; freq-- {
		if len(bucket[freq]) > 0 {
			result = append(result, bucket[freq]...)
			if len(result) >= k {
				break
			}
		}
	}

	// 返回前 k 个
	return result[:k] //[2 1]，也可以直接 return result
}

// 示例使用
func main() {
	fmt.Println(topKFrequent([]int{1, 1, 3, 2, 2, 2}, 2)) // 输出: [2 1]
	//fmt.Println(topKFrequent([]int{4, 4, 4, 5, 5, 6}, 1)) // 输出: [4]
	//fmt.Println(topKFrequent([]int{1}, 1))                // 输出: [1]
}
