package models

import (
	"errors"
	"sync"
)

type JobRequest struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURLs []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}

type Job struct {
	ID      string
	Status  string
	Errors  []string
	Results []ImageResult
}

type ImageResult struct {
	ImageURL string
	Perimeter float64
}

type JobStore struct {
	mu   sync.Mutex
	jobs map[string]*Job
}

func NewJobStore() *JobStore {
	return &JobStore{
		jobs: make(map[string]*Job),
	}
}

func (s *JobStore) CreateJob(request JobRequest) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	jobID := generateID()
	s.jobs[jobID] = &Job{ID: jobID, Status: "ongoing"}
	return jobID
}

func (s *JobStore) GetJob(jobID string) (*Job, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	job, exists := s.jobs[jobID]
	if !exists {
		return nil, errors.New("job not found")
	}
	return job, nil
}

func (s *JobStore) UpdateJobStatus(jobID, status string, errors []string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	job, exists := s.jobs[jobID]
	if exists {
		job.Status = status
		job.Errors = errors
	}
}

func generateID() string {
	return "123"
}
