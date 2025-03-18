package controllers

import (
	"database/sql"
	"encoding/json"
	"example.com/m/v2/API/src/database"
	"example.com/m/v2/API/src/models"
	"example.com/m/v2/API/src/repository"
	"example.com/m/v2/API/src/security"
	"example.com/m/v2/API/src/util/httpresponse"
	"io"
	"log"
	"net/http"
)

// Login authenticate user
func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	var user models.User
	if err := json.Unmarshal(requestBody, &user); err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	db, err := database.Connect()
	userRepository := repository.NewUserRepository(db)
	foundUser, err := userRepository.FindByEmail(user.Email)
	if err != nil {
		httpresponse.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	} else if foundUser.ID == 0 {
		httpresponse.Error(w, http.StatusNotFound, "User not found")
		return
	}

	if !security.CheckPasswordHash(foundUser.Password, user.Password) {
		httpresponse.Error(w, http.StatusUnauthorized, "Incorrect password")
		return
	}

	token, err := security.CreateToken(foundUser.ID)
	if err != nil {
		httpresponse.Error(w, http.StatusInternalServerError, err.Error())
	}
	httpresponse.JSON(w, http.StatusOK, models.Token{
		Token: token,
	})

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database connection")
		}
	}(db)
}

/*
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = seguranca.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		respostas.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := autenticacao.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	respostas.JSON(w, http.StatusOK, modelos.DadosAutenticacao{ID: usuarioID, Token: token})
*/
