class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        backwardIndex = len(s) - 1
        for i in range(len(s)):
            
            if (i == backwardIndex - i ):
                break
            
            if (s[i] != s[backwardIndex - i]):
                return False
        
        return True
            