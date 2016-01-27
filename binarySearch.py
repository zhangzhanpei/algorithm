#二分查找
#在一个有序数组中查找某个元素

def binarySearch(nums,target) :
	left = 0
	right = len(nums) - 1

	while (left <= right) :
		mid = int((left + right)/2)
		if (nums[mid] > target) :
			right = mid - 1
		elif (nums[mid] < target) :
			left = mid + 1
		else:
			return mid 
			#运行到这里才判断是否相等
			#因为一般情况下都是不等的多，应该尽快进入下一个减半的范围
	return -1

nums=[1,2,6,8,12,15,23]
print(binarySearch(nums,15))