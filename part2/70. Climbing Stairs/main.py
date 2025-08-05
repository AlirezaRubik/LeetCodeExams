# from math import comb
# class Solution(object):
#     def climbStairs(self, n):
#         count = 0
#         bn = n
#         while bn >= 0:
#             bn -= 1
#             if bn == 0:
#                 count += 1
#                 break
#         bn = n
#         while bn >= 0:
#             bn -= 2
#             if bn == 0:
#                 count += 1
#                 break

#         for two_steps in range(1, n // 2 + 1):
#             one_steps = n - 2 * two_steps
#             if one_steps >= 0:
#                 count += comb(one_steps + two_steps, two_steps)

#         return count

######################################################
class Solution:
    def climbStairs(self, n):
        a,b = 1,0
        for _ in range(n):
            a,b = a+b,a
        return a