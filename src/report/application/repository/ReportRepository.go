package repository

import (
	"apidos/src/report/domain"
	"apidos/src/report/domain/entities"
)

type ProcessReport struct {
	repo domain.IReport
}

func NewProcessReport(repo domain.IReport) *ProcessReport {
	return &ProcessReport{repo: repo}
}

func (pr *ProcessReport) Execute(id int, tittle, content, status string) int {
	report := entities.Report{
		ID:      id,
		Title:   tittle,
		Content: content,
		Status:  status,
	}

	err := pr.repo.Update(report.ID, report.Title, report.Content, report.Status)
	if err != nil {
		return -1
	}
	return 0
}
