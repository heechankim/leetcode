func stringMerge(s1, s2 string) string {

  var res string

  for i := 0; i < len(s1); i++ {

    res += string(s1[i]) + string(s2[i])

  }

  return res

}

func mergeAlternately(word1 string, word2 string) string {

  diff := len(word1) - len(word2)

  switch {

    case diff == 0:

      return stringMerge(word1, word2)

    case diff > 0:

      return stringMerge(word1[:len(word2)], word2) + word1[len(word2):]

    case diff < 0:

      return stringMerge(word1, word2[:len(word1)]) + word2[len(word1):]

  }

  return ""

}