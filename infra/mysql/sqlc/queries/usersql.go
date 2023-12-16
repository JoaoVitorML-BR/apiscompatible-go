package queries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"login/infra/tools/database"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type userType struct{
	ID uint32 `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

// find All users on database
func FindUsers(w http.ResponseWriter, r *http.Request){
	db, err := database.DatabaseAPI();

	if err != nil{
		w.Write([]byte("Erro ao conectar banco de dados"))
		return
	}

	defer db.Close()

	usersData, err := db.Query("SELECT id, name, password, created_at, updated_at FROM users")

	if err != nil{
		w.Write([]byte("Erro ao pegar todos os usuários da tabela user"))
		return
	}

	defer usersData.Close();

	var users []userType

	for usersData.Next(){
		var user userType

		if err := usersData.Scan(&user.ID, &user.Name, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println("Erro ao escanear os dados do banco de dados:", err)
			w.Write([]byte("Erro ao pegar usuario na struct"))
			return
		}

		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil{
		w.Write([]byte("Erro ao passar usuarios para json"))
		return
	}

}

// find a user especific
func FindUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r) // get params by router /user/{id}

	// converting string to int
	ID, err := strconv.ParseUint(params["id"], 10, 32) // 10 = base number, 32 = bits length

	if err != nil{
		w.Write([]byte("Erro ao converter id param para inteiro"))
		return
	}

	db, err := database.DatabaseAPI()

	if err != nil {
		w.Write([]byte("Erro ao concetar banco de dados para encontrar usuario por ID"))
	}

	userData, err := db.Query("SELECT id, name, password FROM users WHERE id = ?", ID)
	if err != nil{
		w.Write([]byte("Erro ao procurar usúario"))
		return
	}

	var user userType
	if userData.Next(){
		if err := userData.Scan(&user.ID, &user.Name, &user.Password); err != nil {
			fmt.Println("Erro ao escanear os dados do banco de dados:", err)
			w.Write([]byte("Erro ao  pegar usuario por id na struct"))
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil{
		w.Write([]byte("Erro ao passar usuario para json"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)

	if err != nil{
		w.Write([]byte("Erro ao converter id param para inteiro"))
		return
	}

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil{
		w.Write([]byte("Erro ao ler corpo da requisição"))
		return
	}

	// passing json to struct

	var user userType
	if err := json.Unmarshal(bodyReq, &user); err != nil{
		w.Write([]byte("Erro ao converter usuário de json para struct"))
		return
	}

	db, err := database.DatabaseAPI()
	if err != nil {
		w.Write([]byte("Erro ao concetar banco de dados para encontrar usuario por ID"))
	}

	defer db.Close()

	statement, err := db.Prepare("UPDATE users SET name = ?, password = ? WHERE id = ?")
	if err != nil{
		w.Write([]byte("Erro ao criar o statement"))
	}

	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Password, ID); err != nil{
		w.Write([]byte("Erro ao atualizar usúario"))
		return
	}else{
		w.WriteHeader(http.StatusAccepted)
	}

	w.WriteHeader(http.StatusNoContent)
}