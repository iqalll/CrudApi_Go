package repository

import (
	"bytes"
	"fmt"
	"task/db"
	. "task/entity"

	"github.com/gin-gonic/gin"
)

var (
	mahasiswa  Mahasiswa
	mahasiswas []Mahasiswa
	result     gin.H
)

func GetId(id string) (interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	var result gin.H

	row := conn.QueryRow("select id, nama, alamat from mahasiswa where id = ?;", id)
	err = row.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Alamat)
	if err != nil {
		// If no results send null
		result = gin.H{
			"Output": "Data tidak ada",
		}
	} else {
		result = gin.H{
			"Output": mahasiswa,
		}
	}

	return result, nil

}
func GetAll() (interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	rows, err := conn.Query("select id, nama, alamat from mahasiswa;")
	if err != nil {
		fmt.Print(err.Error())
	}
	for rows.Next() {
		err = rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Alamat)
		mahasiswas = append(mahasiswas, mahasiswa)
		if err != nil {
			fmt.Print(err.Error())
		}
	}
	defer rows.Close()
	result = gin.H{
		"Data":        mahasiswas,
		"Jumlah Data": len(mahasiswas),
	}
	return result, nil

}
func AddData(id string, nama string, alamat string) (interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}
	var buffer bytes.Buffer

	stmt, err := conn.Prepare("insert into mahasiswa (id, nama, alamat) values(?,?,?);")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id, nama, alamat)

	if err != nil {
		fmt.Print(err.Error())
	}

	buffer.WriteString(nama)
	buffer.WriteString(" ")
	buffer.WriteString(alamat)
	defer stmt.Close()
	datane := buffer.String()
	result = gin.H{
		"Pesane": fmt.Sprintf(" Yeyyy Berhasil menambahkan Mahasiswa %s ", datane),
	}
	return result, nil
}
func UpdateData(id string, nama string, alamat string) (interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer

	stmt, err := conn.Prepare("update mahasiswa set nama= ?, alamat= ? where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(nama, alamat, id)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Fastest way to append strings
	buffer.WriteString(nama)
	buffer.WriteString(" ")
	buffer.WriteString(alamat)
	defer stmt.Close()
	data := buffer.String()
	result = gin.H{
		"Notif": fmt.Sprintf("Berhasil Merubah Menjadi %s", data),
	}
	return result, nil
}
func DeleteData(id string) (interface{}, error) {
	conn, err := db.Connect()
	if err != nil {
		return nil, err
	}

	stmt, err := conn.Prepare("delete from mahasiswa where id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err.Error())
	}
	result = gin.H{
		"Pesane": fmt.Sprintf("Berhasil Menghapus %s", id),
	}
	return result, nil
}
