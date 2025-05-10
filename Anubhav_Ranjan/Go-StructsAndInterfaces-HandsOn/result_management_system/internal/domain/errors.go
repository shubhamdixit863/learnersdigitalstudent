package domain

import "errors"

var (
	ErrInvalidName        = errors.New("invalid name: must be between 2 and 50 characters")
	ErrInvalidID          = errors.New("invalid student ID: must not be empty")
	ErrInvalidGrade       = errors.New("invalid grade: must be between 0 and 100")
	ErrInvalidSubjects    = errors.New("invalid subjects: must have at least one subject")
	ErrSubjectNotFound    = errors.New("subject not found")
	ErrDuplicateSubject   = errors.New("duplicate subject found")
	ErrInvalidStudentType = errors.New("invalid student type")
)
