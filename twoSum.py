#1. Two Sum
#寻找和为定值的两个数
#在一个有序数组中找出两个元素之和等于给定值

def twoSum(nums, target) :
    i = 0
    j = len(nums) -1
    while (i < j) :
        currSum = nums[i] + nums[j]
        if (currSum == target) :
            print("two elements is %d and %d" %(nums[i], nums[j]))
            return
        elif (currSum < target) :
            i += 1
        else :
            j -= 1
    print('not found')

nums = [2, 7, 11, 15]
twoSum(nums, 18)
