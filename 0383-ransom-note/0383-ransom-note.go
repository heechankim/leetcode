func canConstruct(ransomNote string, magazine string) bool {
  if len(ransomNote) > len(magazine) {
    return false
  }
  
  m := make(map[string]int)
  for _, v := range magazine {
    if m[string(v)] == 0 {
      m[string(v)] = 1
    } else {
      m[string(v)] += 1
    }    
  }
  
  for _, v := range ransomNote {
    if m[string(v)] == 0 {
      return false
    } else {
      m[string(v)] -= 1
    }
  }
  
  return true
}