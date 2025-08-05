class Solution(object):
    def lengthOfLastWord(self, s):
        words = s.strip().split()
        if words:
            print(len(words[-1]))
        else:
            print(0) 

sys = Solution()
sys.lengthOfLastWord("Hello World")