#不使用临时变量，交换两个整数
#1 加减法[注意int类型的最大值]
a = 10
b = 20
a = a + b
b = a - b
a = a - b
print("a = %d, b = %d" %(a, b))

#2 异或
a = 10
b = 20
a ^= b
b ^= a
a ^= b
print("a = %d, b = %d" %(a, b))

#3 呵呵
a = 10
b = 20
a, b = b, a
print("a = %d, b = %d" %(a, b))