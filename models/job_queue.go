package models

import (
	"sync"
)

type JobQueue struct {
	Jobs []Job
	Mu   sync.Mutex
}
