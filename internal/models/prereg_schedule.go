package models

type PreregScheduleRequest struct {
	Kulliyyah string `json:"kulliyyah"`
	Semester  int64  `json:"semester"`
	Session   string `json:"session"`
}
