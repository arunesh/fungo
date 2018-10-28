package world

import "testing"

func TestEven(t *testing.T) {
    if !Even(3) {
        t.Log("forced error")
        t.Fail()
    }
}
