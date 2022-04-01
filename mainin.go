// package main

// import (
// 	"bytes"
// 	"database/sql"
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
// 	err = db.Ping()
// 	if err != nil {
// 		panic("Gagal Menghubungkan ke Database...")
// 	}
// 	defer db.Close()

// 	router := gin.Default()

// 	type Mahasiswa struct {
// 		Id     int    `json: "id"`
// 		Nama   string `json: "nama"`
// 		Alamat string `json: "alamat"`
// 	}

// 	// Menampilkan Detail Data Berdasarkan ID
// 	router.GET("/:id", func(c *gin.Context) {
// 		var (
// 			mahasiswa Mahasiswa
// 			result    gin.H
// 		)
// 		id := c.Param("id")
// 		row := db.QueryRow("select id, nama, alamat from mahasiswa where id = ?;", id)
// 		err = row.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Alamat)
// 		if err != nil {
// 			// If no results send null
// 			result = gin.H{
// 				"Hasile":  "Yahh Tidak ada data yang ditemukan",
// 				"jumlahe": 0,
// 			}
// 		} else {
// 			result = gin.H{
// 				"Hasile":  mahasiswa,
// 				"Jumlahe": 1,
// 			}
// 		}
// 		c.JSON(http.StatusOK, result)
// 	})

// 	// Menampilkan Semua Data
// 	router.GET("/", func(c *gin.Context) {
// 		var (
// 			mahasiswa  Mahasiswa
// 			mahasiswas []Mahasiswa
// 		)
// 		rows, err := db.Query("select id, nama, alamat from mahasiswa;")
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		for rows.Next() {
// 			err = rows.Scan(&mahasiswa.Id, &mahasiswa.Nama, &mahasiswa.Alamat)
// 			mahasiswas = append(mahasiswas, mahasiswa)
// 			if err != nil {
// 				fmt.Print(err.Error())
// 			}
// 		}
// 		defer rows.Close()
// 		c.JSON(http.StatusOK, gin.H{
// 			"Hasile":  mahasiswas,
// 			"Jumlahe": len(mahasiswas),
// 		})
// 	})

// 	// Menambahkan Data Program Studi
// 	router.POST("/", func(c *gin.Context) {
// 		var buffer bytes.Buffer
// 		id := c.PostForm("id")
// 		nama := c.PostForm("nama")
// 		alamat := c.PostForm("alamat")
// 		stmt, err := db.Prepare("insert into mahasiswa (id, nama, alamat) values(?,?,?);")
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		_, err = stmt.Exec(id, nama, alamat)

// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}

// 		// Fastest way to append strings
// 		buffer.WriteString(nama)
// 		buffer.WriteString(" ")
// 		buffer.WriteString(alamat)
// 		defer stmt.Close()
// 		datane := buffer.String()
// 		c.JSON(http.StatusOK, gin.H{
// 			"Pesane": fmt.Sprintf(" Yeyyy Berhasil menambahkan Mahasiswa %s ", datane),
// 		})
// 	})

// 	// PUT Merubah Data
// 	router.PUT("/", func(c *gin.Context) {
// 		var buffer bytes.Buffer
// 		id := c.PostForm("id")
// 		nama := c.PostForm("nama")
// 		alamat := c.PostForm("alamat")
// 		stmt, err := db.Prepare("update mahasiswa set nama= ?, alamat= ? where id= ?;")
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		_, err = stmt.Exec(nama, alamat, id)
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}

// 		// Fastest way to append strings
// 		buffer.WriteString(nama)
// 		buffer.WriteString(" ")
// 		buffer.WriteString(alamat)
// 		defer stmt.Close()
// 		datane := buffer.String()
// 		c.JSON(http.StatusOK, gin.H{
// 			"Pesane": fmt.Sprintf("Berhasil Merubah Menjadi %s", datane),
// 		})
// 	})

// 	// Delete resources
// 	router.DELETE("/", func(c *gin.Context) {
// 		id := c.PostForm("id")
// 		stmt, err := db.Prepare("delete from mahasiswa where id= ?;")
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		_, err = stmt.Exec(id)
// 		if err != nil {
// 			fmt.Print(err.Error())
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"Pesane": fmt.Sprintf("Berhasil Menghapus %s", id),
// 		})
// 	})
// 	router.Run(":8080")
// }
