package main

func p1easy(n int) int {
    sum := 0
    for i:=0; i<=n; i++ {
        if i%3==0 || i%5==0 {
            sum += i
        }
    }
    
    return sum
}

