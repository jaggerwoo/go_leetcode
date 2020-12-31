package main

func main() {
    println(isPalindrome(121))
}

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0){
        return false
    }

    res := 0
    for x > res {
        res = res * 10 + x % 10
        x = x / 10
    }

    return res == x || x == res / 10
}