# class Solution(object):
#     def isPalindrome(self, s):
#        newlower=s.lower()
#        news=newlower.strip()
#        notin=[":",",",":","<",">","/","?",".","@","#","$","%","^","&","(",")","-","_","!","~"," ","{","}","[","]","|","'","\""]
#        arr=list(news)
#        secarr=list(news)
#        for i in range(len(arr)- 1, -1, -1):
#            if arr[i] in notin:
#                del arr[i]

#        for i in range(len(secarr)- 1, -1, -1):
#            if secarr[i] in notin :
#                del secarr[i]

#        secarr.reverse()
#        if str(secarr)==str(arr):
#           print("True")
#           return True
#        else:
#            print("False")
#            return False


# sol=Solution()
# sol.isPalindrome("Marge, let's \"[went].\" I await {news} telegram.")
#############################################################################
class Solution(object):
    def isPalindrome(self, s):
        cleaned = [ch.lower() for ch in s if ch.isalnum()]
        result = cleaned == cleaned[::-1]
        print(result)
        return result