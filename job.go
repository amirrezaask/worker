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
	run    func() error
	status JobStatus
}

func NewSimpleJob(f func() error) Job {
	return &simpleJob{
		run: f,
	}
}

func (j *simpleJob) ID() string {
	return ""
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
