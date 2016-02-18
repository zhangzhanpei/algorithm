#162. Find Peak Element
#给定一个数组，返回一个峰值的位置[峰值：该元素大于相邻的两个元素]

def findPeakElement(Arr):
    if len(Arr)<3:
        return 0
    for i in range(1,len(Arr)):
        if Arr[i-1]<Arr[i]>Arr[i+1]:
            return i

arr=[1,3,4,5,7,6]
print(findPeakElement(arr)) #return 4
