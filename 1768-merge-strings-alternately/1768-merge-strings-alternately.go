func mergeAlternately(word1 string, word2 string) string {
  
  if len(word1) == 0 {
    return word2
  }
  
  if len(word2) == 0 {
    return word1
  }
  
  return word1[:1] + word2[:1] + mergeAlternately(word1[1:], word2[1:])
}