go
package handlers

import (
    "encoding/json"
    "net/http"
    "myproject/db"
    "myproject/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    // Store user in MySQL or MongoDB...
}