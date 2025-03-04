package services

type Service interface {
	GetData() string
	SearchData() string
	AddData(dta string) error
}
