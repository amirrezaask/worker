# Worker

Worker is a simple job queue library for Golang with fluent API and simplicity in mind.

## Terminology

### Queue

Each queue contains jobs that are being add from the user of the library and worker use this queue to get new jobs and do them.

### Job

Job is the unit of work that needs to be done which is simply a function.

### Worker

Worker does the actual work from pushing new jobs to queue, getting new jobs for each worker and finally doing the jobs.
