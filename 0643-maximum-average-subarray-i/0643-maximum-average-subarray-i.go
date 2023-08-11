func findMaxAverage(nums []int, k int) float64 {

  sum := 0

  for i := 0; i <= len(nums) - k; i++ {

    temp := 0

    for j := i; j < i + k; j++ {

      temp += nums[j]

    }

    if temp > sum || sum == 0 {

      sum = temp

    }

  }

  return float64(sum) / float64(k)

}