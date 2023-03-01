package model

type Student struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Nama  string `json:"nama"`
	Umur  uint   `json:"umur"`
	Kelas string `json:"kelas"`
}

type StudentResponse struct {
	ID    string `json:"id" bson:"_id,omitempty"`
	Nama  string `json:"nama"`
	Umur  uint   `json:"umur"`
	Kelas string `json:"kelas"`
}
