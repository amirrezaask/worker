package worker

type Queue interface {
	Next() (string, error)
	Add(jobID string) error
	Len() int
}

type channelQueue struct {
	c chan string
}

func NewChannelQueue(buffer int) Queue {
	return &channelQueue{
		c: make(chan string, buffer),
	}
}

func (c *channelQueue) Next() (string, error) {
	return <-c.c, nil
}

func (c *channelQueue) Add(jobID string) error {
	c.c <- jobID
	return nil
}

func (c *channelQueue) Len() int {
	return len(c.c)
}
