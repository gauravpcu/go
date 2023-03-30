package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

/*
Write three APIs with GET/POST to minic the following
restore the backup
get a single record
get a single record by manipulating a string and a number
get a 100 records
get a 100 records by manipulating a string and a number
get 1000 records
post a single record to insert
post a single record to update
post 10 record inserts
post 10 record updates
post 100 record inserts
post 100 record updates
*/

var db *sql.DB

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

func main() {
	configDb()
	router := gin.Default()
	router.GET("/albums/:count/:manipulate", getAlbums)
	router.POST("/albums", postAlbums)
	router.POST("/albums/update", updateAlbum)
	router.Run("localhost:8080")

}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbums(c *gin.Context) {

	count, nil := strconv.ParseInt(c.Param("count"), 10, 32)
	fmt.Print(count)
	manipulate, nil := strconv.ParseBool(c.Param("manipulate"))
	fmt.Print(manipulate)

	albums, err := albumsByArtistCount(count, manipulate)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, albums)
	fmt.Printf("Albums found: %v\n", albums)
	return

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {

	var albums []album

	body, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	err = json.Unmarshal(body, &albums)

	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	for _, newAlbum := range albums {
		insert, err := db.Query("INSERT INTO album(title, artist, price) VALUES(?,?,?)", newAlbum.Title, newAlbum.Artist, newAlbum.Price)
		if err != nil {
			panic(err.Error())
		}
		defer insert.Close()

		fmt.Println("Successfully inserted into albums table")

	}

	//albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, albums)
}

//curl http://localhost:8080/albums/update     --include     --header "Content-Type: application/json"     --request "POST"     --data '{"id": "1","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

func updateAlbum(c *gin.Context) {

	var newAlbum album
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	update, err := db.Query("UPDATE album SET title=?, artist=?, price=? WHERE id=?", newAlbum.Title, newAlbum.Artist, newAlbum.Price, newAlbum.ID)
	if err != nil {
		panic(err.Error())
	}
	defer update.Close()

	fmt.Println("Successfully updated albums table")
	//albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func configDb() {
	// Capture connection properties.
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}
	// Get a database handle.
	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]album, error) {
	// An albums slice to hold data from returned rows.
	var albums []album

	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

func albumsByArtistCount(count int64, manipulate bool) ([]album, error) {
	// An albums slice to hold data from returned rows.
	var albums []album

	rows, err := db.Query("select * from album LIMIT ?", count)

	if err != nil {
		return nil, fmt.Errorf("albumsByArtistCount %q: %v", count, err)
	}

	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtistCount %q: %v", count, err)
		}
		if manipulate {
			alb = manipulateAlbum(alb)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtistCount %q: %v", count, err)
	}
	return albums, nil
}

func manipulateAlbum(alb album) album {
	alb.Title += " - added 10% markup (" + strconv.FormatFloat(alb.Price, 'f', 6, 64) + ")"
	alb.Artist += " -Gaurav "
	alb.Price += alb.Price * 0.1

	return alb
}
