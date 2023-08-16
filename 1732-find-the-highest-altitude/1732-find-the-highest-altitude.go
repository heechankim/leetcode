func largestAltitude(gain []int) int {
  
  height := []int{0}
  
  for i := 0; i < len(gain); i++ {
    height = append(height, height[i] + gain[i])
  }
  
  max := 0
  
  for _, v := range(height) {
    if v > max {
      max = v
    }
  }
  
  return max
}