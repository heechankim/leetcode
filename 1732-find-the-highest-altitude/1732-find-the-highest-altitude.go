func largestAltitude(gain []int) int {
  
  max := 0
  temp := 0
  for _, v := range gain {
    
    temp += v
    
    if temp > max {
      max = temp
    }
  }
  
  return max
}