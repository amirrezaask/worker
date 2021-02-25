package worker

type Queue interface {
	Next() (Job, error)
	Add(j Job) error
	Len() int
}

type channelQueue struct {
	c chan Job
}

func NewChannelQueue(buffer int) Queue {
	return &channelQueue{
		c: make(chan Job, buffer),
	}
}

func (c *channelQueue) Next() (Job, error) {
	return <-c.c, nil
}

func (c *channelQueue) Add(j Job) error {
	c.c <- j
	return nil
}

func (c *channelQueue) Len() int {
	return len(c.c)
}
