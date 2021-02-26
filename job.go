package worker

type Job interface {
	ID() string
	Run() error
	Status() JobStatus
}
type JobStatus int

const (
	JobNotStarted = iota
	JobStarted
	JobFinishedNoErrs
	JobFinishedWithErrs
)

type simpleJob struct {
	id     string
	run    func() error
	status JobStatus
}

func NewSimpleJob(id string, f func() error) Job {
	return &simpleJob{
		id:  id,
		run: f,
	}
}

func (j *simpleJob) ID() string {
	return j.id
}

func (j *simpleJob) Run() error {
	j.status = JobStarted
	err := j.run()
	if err != nil {
		j.status = JobFinishedWithErrs
	} else {
		j.status = JobFinishedNoErrs
	}
	return err
}

func (j *simpleJob) Status() JobStatus {
	return j.status
}
