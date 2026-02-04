package main

import (
	"context"
	"fmt"
	"json"
	"net/http"
)


func addHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Main page!!")

	ctx := r.Context()

	rows, err := db.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Send reponse back to request. (Crate a JSON)
	// w.WriteHeader(http.StatusOK)
	// responseBody := []byte(`{"message": "Hi, you are doing well"}`)
	// _, err := w.Write(responseBody)
	// if err != nil {
	// 	fmt.Println("Error: %v", err)
	// }


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rows)
}

func main(){
	// Context setup.

	// This already use go routine
	http.HandleFunc("/", addHandler)
	http.ListenAndServe(":8080", nil)
}
