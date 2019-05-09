
target = 100
coins = [37, 19, 5, 2, 1]


def soln (used_coins):
    sum_used_coins = sum(used_coins)
    if (sum_used_coins == target) :
        print(used_coins)
        return True
    if (sum_used_coins > target) :
        return False
    

    for x in coins:
        used_coins.append(x)
        if (soln(used_coins)):
            return True
        used_coins.pop()
    return False


soln([])

    # // let sum_used_coins = sum(used_coins)
    # // if (sum_used_coins == target) {
    # //     print(used_coins)
    # //     return true
    # // }

    # // if (sum_used_coins > target) {
    # //     return false
    # // }

    # // let i = 0
    # // for (i < len(coins)) {
    # //     push(used_coins, coins[i])
    # //     if (soln(used_coins)) {
    # //         return true
    # //     }
    # //     pop(used_coins)
    # //     let i = i + 1
    # // }
    # // return false