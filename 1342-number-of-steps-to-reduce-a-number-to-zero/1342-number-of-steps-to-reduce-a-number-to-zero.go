func numberOfSteps(num int) int {
  
  var step int = 0
  
  if num == 0 {
    step = 0
  }
  
  for num != 0 {
    if num % 2 == 0 {
      num /= 2
    } else {
      num -= 1
    }
    
    step += 1
  }
  
  return step
}