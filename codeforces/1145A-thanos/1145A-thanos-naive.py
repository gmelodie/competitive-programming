

def isOrdered(arr):
    for i in range(len(arr) - 1):
        if arr[i] > arr[i+1]:
            return False
    return True


def solve(arr):
    if isOrdered(arr):
        return len(arr)

    a = solve(arr[:len(arr)//2])
    b = solve(arr[len(arr)//2:])

    if a > b:
        return a

    return b


if __name__ == '__main__':

    l = int(input())

    arr = [int(a) for a in input().split()]

    print(solve(arr))
