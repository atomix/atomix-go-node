

package counter

import (
	counter "github.com/atomix/api/go/atomix/primitive/counter"
)

type Service interface {
    // Set sets the counter value
    Set(*counter.SetRequest) (*counter.SetResponse, error)
    // Get gets the current counter value
    Get(*counter.GetRequest) (*counter.GetResponse, error)
    // Increment increments the counter value
    Increment(*counter.IncrementRequest) (*counter.IncrementResponse, error)
    // Decrement decrements the counter value
    Decrement(*counter.DecrementRequest) (*counter.DecrementResponse, error)
}