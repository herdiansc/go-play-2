package tenisball

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// Response wraps response for tennis ball container model.
type Response struct {
	ListContainers  Containers `json:"listContainers"`
	FullContainerNo int        `json:"fullContainerNo"`
}

// Containers contains balls.
type Containers map[int]int

// Length returns total of container.
func (c Containers) Length() int {
	return len(c)
}

// Load inserts a ball randomly into each container until its full.
func (c Containers) Load() int {
	number := 0
	for {
		number = c.GetRandContainer(1, c.Length())
		fmt.Println("loading a ball to container no. " + strconv.Itoa(number))
		c[number]++
		if c.GetTotalBall(number) >= 3 {
			break
		}
	}
	return number
}

// GetTotalBall returns total ball in a container.
func (c Containers) GetTotalBall(number int) int {
	return c[number]
}

// GetRandContainer gets random container number.
func (c Containers) GetRandContainer(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+min) + min
}
