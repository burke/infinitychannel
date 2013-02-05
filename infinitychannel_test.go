package infinitychannel

import "testing"

func TestInfinitychannel(t *testing.T) {
  recv, send := New()
    for i := 0; i < 100; i++ {
      send <- "omg"
    }
  
  for i := 0; i < 100; i++ {
    <- recv
    println("omg")
  }
}
