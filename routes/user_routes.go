package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/theinlaoq/booking-api-testcase/db"
	"github.com/theinlaoq/booking-api-testcase/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var validate = validator.New()

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// PostUserHandler godoc
// @Summary Создание пользователя
// @Description Создает нового пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /users [post]
func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	validateErr := validate.Struct(user)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validateErr.Error())
		return
	}
	hashedPassword, hashErr := hashPassword(user.Password)
	if hashErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(hashErr.Error()))
		return
	}

	user.Password = hashedPassword
	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("POST"))
	json.NewEncoder(w).Encode(user)
}

// GetUsersHandler godoc
// @Summary Получение всех пользователей
// @Description Возвращает список всех пользователей
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /users [get]
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Preload("Booking").Find(&users)
	json.NewEncoder(w).Encode(&users)
}

// GetUserHandler godoc
// @Summary Получение пользователя по id
// @Description Возвращает пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /users/{id} [get]
func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	existence := db.DB.Preload("Booking").First(&user, id)
	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// UpdateUserHandler godoc
// @Summary Обновление данных пользователя по полю id
// @Description Обновляет данные пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body models.User true "Username(optional), Password(optional)"
// @Success 200 {object} models.User
// @Failure 400 {string} string "Bad Request"
// @Router /users/{id} [put]
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var updatedUser models.User
	var user models.User
	params := mux.Vars(r)
	id, IdError := strconv.Atoi(params["id"])
	if IdError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(IdError.Error()))
		return
	}
	dataError := json.NewDecoder(r.Body).Decode(&updatedUser)
	if dataError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(dataError.Error()))
		return
	}
	existence := db.DB.Preload("Booking").First(&user, id)
	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}

	if user.Username != updatedUser.Username {
		user.Username = updatedUser.Username
	}

	if updatedUser.Password != "" {
		hashedPassword, hashErr := hashPassword(updatedUser.Password)
		if hashErr != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(hashErr.Error()))
			return
		}
		user.Password = hashedPassword
	}
	user.UpdatedAt = time.Now()
	update := db.DB.Save(&user)
	updateError := update.Error
	if updateError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(updateError.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}

// DeleteUserHandler godoc
// @Summary Удаление пользователя по полю id
// @Description Удаляет пользователя
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /users/{id} [delete]
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	existence := db.DB.Preload("Booking").First(&user, id)
	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("User not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusOK)
}
