package models

import "time"

type JobCard struct {
	ID       string
	Date     time.Time
	FleetNum string
	WorkDone []WorkDone
}

type WorkDone struct {
	description string
}
