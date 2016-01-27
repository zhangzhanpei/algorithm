#冒泡排序

def bubbleSort(nums) :
    n = len(nums)
    k = 0
    for i in range(n) :
       for j in range(i+1, n) : #一趟过后会把最小的数交换到最左边
           k += 1
           if (nums[i] > nums[j]) :
           tmp = nums[i]
           nums[i] = nums[j]
           nums[j] = tmp
    print(nums)

nums = [2, 7, 12, 6, 0]
bubbleSort(nums)
