def binary_search(list, item):
    low = 0
    half = 0
    high = len(list) -1
    while low <= high:
        half = (low + high) // 2
        guess = list[half]
        if guess == item:
            return half
        if guess > item:
            high = half -1
        else:
            low = half + 1
    return None
my_list = [1,3,5,7,9]

print(binary_search(my_list, 3))
print(binary_search(my_list,-1))