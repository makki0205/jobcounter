package jobcounter

type JobCounter struct {
	counter int
	c       chan int
}

func NewJob() JobCounter {
	jc := JobCounter{
		counter: 0,
		c:       make(chan int),
	}
	go func(jc JobCounter) {
		for {
			jc.counter += <-jc.c
		}
	}(jc)
	return jc
}

func (j *JobCounter) Add(c uint) {
	j.c <- int(c)
}

func (j *JobCounter) Done() {
	j.c <- -1
}
func (j *JobCounter) IsEndJob() bool {
	return j.counter <= 0
}
