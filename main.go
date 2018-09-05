package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha256"

	"github.com/gorilla/mux"
)

type SignRequest struct {
	Namespace   string `json:"namespace"`
	Message string `json:"message"`
	PreviousHash   []byte `json:"previousHash"`
}

type SignResponse struct {
	NewHash []byte `json:"newHash"`
	Signature []byte `json:"signature"`
}

func signIt(namespace string, message string, previousHash []byte) (b []byte, err error) {
	var K *rsa.PrivateKey

	K, err = rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}

	rng := rand.Reader
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rng, K, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Println(err)
	}

	combined := append(signature[:], previousHash[:]...)
	hashed = sha256.Sum256(combined)

	res := SignResponse{NewHash: hashed[:], Signature: signature}

	b, err = json.Marshal(res)
	if err != nil {
		fmt.Println("error:", err)
	}
	return b, err
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sr SignRequest

	j, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(j, &sr)
	if err != nil {
		fmt.Println(err)
	}

	b, _ := signIt(sr.Namespace, sr.Message, sr.PreviousHash)
	w.Write(b)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sign", handle).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8001", nil))
}
