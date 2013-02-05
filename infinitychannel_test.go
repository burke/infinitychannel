package infinitychannel

import "testing"

func TestInfinitychannel(t *testing.T) {
  recv, send := New()
  go func() {
    for i := 0; i < 100; i++ {
      send <- "omg"
    }
  }()
  
  for i := 0; i < 100; i++ {
    <- recv
    println("omg")
  }
}
