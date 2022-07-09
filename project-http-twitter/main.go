package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Tweet struct {
	Message  string `json:"message"`
	Location string `json:"location"`
}

func (s Server) addTweet(w http.ResponseWriter, r *http.Request) {
	tweet := Tweet{}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Printf("Could not read request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.Unmarshal(body, &tweet); err != nil {
		fmt.Printf("Failed to parse tweet")
	}

	id, addErr := s.repository.AddTweet(tweet)
	if addErr != nil {
		log.Println("Failed to add user:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		tr := TweetResponse{Id: id}
		serialized, _ := json.Marshal(tr)

		fmt.Fprint(w, string(serialized))
	}
}

func (s Server) listAllTweets(w http.ResponseWriter, r *http.Request) {
	tweets, err := s.repository.GetTweets()

	if err != nil {
		fmt.Fprint(w, "Could not get tweets")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tList := tweetsList{Tweets: tweets}
	tweetsJson, serializeError := json.Marshal(tList)

	if serializeError != nil {
		fmt.Fprint(w, "Error while serializing tweets")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(tweetsJson)
}
func (s Server) tweets(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	defer func() {
		duration := time.Since(start)
		fmt.Printf("%s %s %s\n", r.Method, r.URL, duration)
	}()

	if r.Method == http.MethodPost {
		s.addTweet(w, r)
	} else if r.Method == http.MethodGet {
		s.listAllTweets(w, r)
	}
}

func main() {
	s := Server{repository: &TweetMemoryRepository{}}

	http.HandleFunc("/tweets", s.tweets)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type TweetResponse struct {
	Id int `json:"ID"`
}

type TweetRepository interface {
	AddTweet(t Tweet) (int, error)
	GetTweets() ([]Tweet, error)
}

type TweetMemoryRepository struct {
	Tweets []Tweet
}

func (tweetRepo *TweetMemoryRepository) AddTweet(tweet Tweet) (int, error) {
	tweetRepo.Tweets = append(tweetRepo.Tweets, tweet)
	return len(tweetRepo.Tweets), nil
}

func (tweetRepo *TweetMemoryRepository) GetTweets() ([]Tweet, error) {
	return tweetRepo.Tweets, nil
}

type Server struct {
	repository TweetRepository
}

type tweetsList struct {
	Tweets []Tweet `json:"tweets"`
}
