package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"math/rand"
	"encoding/json"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}


type Card struct {
	Id 			int
	PhraseFrom 	string
	PhraseTo	string
	LngFrom 	string
	LngTo 		string
}

type Cards []Card

// =============== test function ============================
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
// =============== test function ============================


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Write([]byte("¯\\_(ツ)_/¯"))
}

func GetListCards(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// todo: pagination
	rows, err := db.Query("SELECT * FROM card LIMIT 100")
	if err != nil {
		log.Panic(err)
		w.Write([]byte("500: Internal Server Error"))
	}

	defer rows.Close()

	cards := Cards{}
	for rows.Next() {
		card := Card{}
		rows.Scan(&card.Id, &card.PhraseTo, &card.PhraseFrom, &card.LngTo, &card.LngFrom)
		cards = append(cards, card)
	}

	if err != nil {
		log.Panic(err)
	}

	js, err := json.Marshal(cards)
	if err != nil {
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func GetCard(w http.ResponseWriter, r *http.Request, args httprouter.Params){
	cardId := args.ByName("id")

	card := Card{}

	err := db.QueryRow("SELECT * FROM card WHERE card.id = $1", cardId).Scan(
		&card.Id, &card.PhraseTo, &card.PhraseFrom, &card.LngTo, &card.LngFrom)
	if err != nil {
		log.Panic(err)
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	js, err := json.Marshal(card)
	if err != nil {
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func CreateCard(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	// todo: make a real data
	card := Card{
		Id:10,
		PhraseFrom:RandStringBytes(8),
		PhraseTo:RandStringBytes(8),
		LngFrom:RandStringBytes(4),
		LngTo:RandStringBytes(4),
	}

	var lastId int
	err := db.QueryRow(
		"INSERT INTO card (phrase_from, phrase_to, lng_from, lng_to) VALUES ($1, $2, $3, $4) RETURNING id;",
		 card.PhraseFrom, card.PhraseTo, card.LngFrom, card.LngTo).Scan(&lastId)
	if err != nil {
		log.Panic(err)
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	card.Id = int(lastId)

	js, err := json.Marshal(card)
	if err != nil {
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
//
//func UpdateCard(w http.ResponseWriter, r *http.Request, args httprouter.Params){
//	title := args.ByName("id")
//}

func DeleteCard(w http.ResponseWriter, r *http.Request, args httprouter.Params){
	var deletedCardId int
	cardId := args.ByName("id")
	err := db.QueryRow("DELETE FROM card WHERE card.id = $1 RETURNING id;", cardId).Scan(&deletedCardId)
	if err != nil {
		log.Panic(err)
		// fixme: does not work
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Write([]byte(string(deletedCardId)))
}


func main(){
	InitDB("postgres://dictcard:dictcard@db/dictcard?sslmode=disable")
	defer db.Close()

	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/card/", GetListCards)
	router.POST("/card/", CreateCard)
	router.GET("/card/:id/", GetCard)
	//router.PUT("/card/:id/", UpdateCard)
	router.DELETE("/card/:id/", DeleteCard)
	log.Fatal(http.ListenAndServe(":8080", router))
}