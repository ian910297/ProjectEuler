package main

import (
    "testing"
)

func TestP1easy(t *testing.T) {
    pairs := [...][2]int{{9, 23}, {999, 233168}}
    for _, pair := range pairs {
        in, out := pair[0], pair[1]
        if x:=p1easy(in); x!=out {
            t.Errorf("p1easy(%d) = %d, should be %d", in, x, out)
        } else {
            t.Log("p1easy(%d) = %d, should be %d", in, x, out)
        }
    }
}

func BenchmarkP1easy(b *testing.B) {
    in := 1000
    b.StartTimer()
    for i:=0; i<b.N; i++ {
        p1easy(in)
    }
    b.StopTimer()
}
