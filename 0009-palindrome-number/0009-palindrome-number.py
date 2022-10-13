class Solution:
    def isPalindrome(self, x: int) -> bool:
        s = str(x)
        print(s, len(s))
        for i in range(len(s)):
            
            if (i == (len(s) - 1) - i ):
                break
            
            if (s[i] != s[(len(s) - 1) - i]):
                return False
        
        return True
            