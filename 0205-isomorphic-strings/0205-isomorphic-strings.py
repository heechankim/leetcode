class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        hash_table = {}
        
        for i in range(len(s)):
            
            if s[i] not in hash_table:
                if t[i] in hash_table.values():
                    return False
                
                hash_table.update({s[i]: t[i]})
            else:
                if hash_table[s[i]] != t[i]:
                    return False
        return True
        