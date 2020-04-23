import re
import matplotlib as plt

PATH = "C:/Users/Woo Kei Cheung/Desktop/CSHW/cs4700/Assign24_4700/Assignment24"

def getWords(filename):
    arr = []
    file = open(PATH + filename, "r")
    for f in file:
        f.replace('\n', '').replace(' ', '')
        n = len(f) - 1
        new = f[1:n]
        arr = re.split(', ', new)
    return arr

def graph(py, java, go):
    start = 1000
    timeArr = []
    while start <= 100000:
        timeArr.append(start)
        start += 2000
    plt.yscale('linear')
    plt.xscale('linear')
    plt.plot(timeArr, py, color='blue', label = "Python")
    plt.plot(timeArr, java, color='red', label = "Java")
    plt.plot(timeArr, go, color='green', label = "GO")
    plt.xlabel('Size of Array')
    plt.ylabel('Run Time (seconds)')
    plt.title("Run Time of Merge Sort Compared")
    plt.legend()
    plt.show()

def run():
    py = getWords("/mergeSortPy/pythonTimes.txt")
    java = getWords()
    go = getWords()
    graph(py, java, go)