package validation

import (
	"result_management_system/internal/domain"
	"strings"
)

const (
	MinNameLength = 2
	MaxNameLength = 50
	MinGrade      = 0
	MaxGrade      = 100
)

func ValidateStudent(id string, name string) error {
	if strings.TrimSpace(id) == "" {
		return domain.ErrInvalidID
	}

	nameLen := len(strings.TrimSpace(name))
	if nameLen < MinNameLength || nameLen > MaxNameLength {
		return domain.ErrInvalidName
	}

	return nil
}

func ValidateGrade(grade int) error {
	if grade < MinGrade || grade > MaxGrade {
		return domain.ErrInvalidGrade
	}

	return nil
}
