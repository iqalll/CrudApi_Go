package entity

type Mahasiswa struct {
	Id     int    `json: "id"`
	Nama   string `json: "nama"`
	Alamat string `json: "alamat"`
}
