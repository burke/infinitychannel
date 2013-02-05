package infinitychannel

import "testing"
import "fmt"

func TestInfinitychannel(t *testing.T) {
	send, recv := New()
	for i := 0; i < 100; i++ {
		send <- i
	}

	for i := 0; i < 100; i++ {
		x, _ := (<-recv).(int)
		fmt.Println("omg ", x)
	}
}
