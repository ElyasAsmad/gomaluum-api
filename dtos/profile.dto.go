package dtos

type Profile struct {
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	MatricNo string `json:"matric_no"`
	Level    string `json:"level"`
	Kuliyyah string `json:"kuliyyah"`
}
