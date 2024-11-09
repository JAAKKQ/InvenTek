package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	http.HandleFunc("/posts", postsHandler)
	http.HandleFunc("/posts/", postHandler)

	fmt.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type EquipInfo struct {
	ID           int    `json:"id"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Year         int    `json:"year"`
	SerialNum    string `json:"serialNum"`
}

var posts = make(map[int]EquipInfo)
var nextID int
var postsMu sync.Mutex

func postsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetPosts(w, r)
	case "POST":
		handlePostPosts(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/posts/"):])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		handleGetPost(w, r, id)
	case "DELETE":
		handleDeletePost(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetPosts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	postsMu.Lock()

	defer postsMu.Unlock()

	ps := make([]EquipInfo, 0, len(posts))
	for _, p := range posts {
		ps = append(ps, p)
	}

	db, err := sql.Open("mysql", "root:YILHtSyE2QcFyRn4w6KmHEpx@tcp(10.9.0.20:3306)/junction")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success get all!")

	rows, err := db.QueryContext(context.Background(), "SELECT * FROM equipment")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	equips := make([]EquipInfo, 0)

	for rows.Next() {
		var newEquip EquipInfo
		if err := rows.Scan(&newEquip.ID, &newEquip.Manufacturer, &newEquip.Model, &newEquip.Year, &newEquip.SerialNum); err != nil {
			log.Fatal(err)
		}
		equips = append(equips, newEquip)
	}
	// If the database is being written to ensure to check for Close
	// errors that may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report the last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(equips)
}

func handlePostPosts(w http.ResponseWriter, r *http.Request) {
	var p EquipInfo

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body 1", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(body, &p); err != nil {
		http.Error(w, "Error parsing request body 2", http.StatusBadRequest)
		return
	}

	postsMu.Lock()
	defer postsMu.Unlock()

	db, err := sql.Open("mysql", "root:YILHtSyE2QcFyRn4w6KmHEpx@tcp(10.9.0.20:3306)/junction")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success post!")

	var maxId int
	err = db.QueryRowContext(context.Background(), "select MAX(id) from equipment").Scan(
		&maxId)
	if err != nil {
		log.Fatal("unable to get max id", err)
	}

	fmt.Println("maxId: ", maxId)
	p.ID = maxId + 1
	posts[p.ID] = p

	query := "INSERT INTO equipment (`id`, `manufacturer`, `model`, `year`, `serialNum`) VALUES(?, ?, ?, ?, ?)"
	insert, err := db.ExecContext(context.Background(), query, p.ID, p.Manufacturer, p.Model, p.Year, p.SerialNum)
	if err != nil {
		log.Fatalf("impossible insert: %s", err)
	}
	id, err := insert.LastInsertId()
	if err != nil {
		log.Fatalf("impossible to retrieve last inserted id: %s", err)
	}
	log.Printf("inserted id: %d", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}

func handleGetPost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	db, err := sql.Open("mysql", "root:YILHtSyE2QcFyRn4w6KmHEpx@tcp(10.9.0.20:3306)/junction")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success get by ID!")

	var newEquip EquipInfo
	err = db.QueryRowContext(context.Background(), "select * from equipment where id=?", id).Scan(
		&newEquip.ID, &newEquip.Manufacturer, &newEquip.Model, &newEquip.Year, &newEquip.SerialNum)
	if err != nil {
		log.Fatal("unable to execute search query", err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEquip)
}

func handleDeletePost(w http.ResponseWriter, r *http.Request, id int) {
	postsMu.Lock()
	defer postsMu.Unlock()

	// _, ok := posts[id]
	// if !ok {
	// 	http.Error(w, "Post not found", http.StatusNotFound)
	// 	return
	// }

	db, err := sql.Open("mysql", "root:YILHtSyE2QcFyRn4w6KmHEpx@tcp(10.9.0.20:3306)/junction")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Success delete!")
	fmt.Println("id to delete: ", id)

	db.ExecContext(context.Background(), "delete from equipment where id=?", id)

	//delete(posts, id)
	w.WriteHeader(http.StatusOK)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
