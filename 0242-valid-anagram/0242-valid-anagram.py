class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        hash_table = {}

        for i in range(len(s)):

            if s[i] not in hash_table:
                hash_table[s[i]] = 1
            else:
                hash_table[s[i]] += 1

        for i in range(len(t)):

            if t[i] not in hash_table:
                return False
            else:
                hash_table[t[i]] -= 1

        result = map(lambda x: x != 0, hash_table.values())
        if True in result:
            return False

        return True
        