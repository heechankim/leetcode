func maximumWealth(accounts [][]int) int {
  var maxWealth int = 0
  for _, v := range accounts {
    
    var temp int = 0
    for _, m := range v {
      temp += m
    }
    
    if temp > maxWealth {
      maxWealth = temp
    }
  }
  
  return maxWealth
}