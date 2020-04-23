import numpy as np
import time

def getRandArray(n):
    return np.random.randint(1000, size=n) 

def mergeSort(a):
    if len(a) > 1:
        size = len(a)//2
        left = mergeSort(a[:size])
        right = mergeSort(a[size:])

        arr = []

        while len(left) > 0 and len(right) > 0:
            if left[0] <= right[0]:
                arr.append(left[0])
                left = left[1:]
            else:
                arr.append(right[0])
                right = right[1:]

        if len(left) > 0:
            for i in left:
                arr.append(i)
        if len(right) > 0:
            for i in right:
                arr.append(i)
        return arr
    return a

def runTest():
    size = 1000
    timeArray = []
    while size <= 100000:
        arr = getRandArray(size)
        start = time.time()
        mergeSort(arr)
        end = time.time() - start
        timeArray.append(end)
        size = size + 2000
    print(timeArray)


runTest()