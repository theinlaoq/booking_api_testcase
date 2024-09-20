package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/theinlaoq/booking-api-testcase/db"
	"github.com/theinlaoq/booking-api-testcase/models"
	"gorm.io/gorm"
)

// PostBookingHandler godoc
// @Summary Создание бронирования
// @Description Создает новое бронирование
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param booking body models.Booking true "Данные бронирования"
// @Success 200 {object} models.Booking
// @Failure 400 {string} string "Bad Request"
// @Router /bookings [post]
func PostBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	json.NewDecoder(r.Body).Decode(&booking)
	validateErr := validate.Struct(booking)
	if validateErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(validateErr.Error())
		return
	}
	createdBooking := db.DB.Create(&booking)
	err := createdBooking.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("POST"))
	json.NewEncoder(w).Encode(booking)
}

// GetBookingsHandler godoc
// @Summary Получение всех бронирований
// @Description Возвращает список всех бронирований
// @Tags bookings
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Booking
// @Failure 400 {string} string "Bad Request"
// @Router /bookings [get]
func GetBookingsHandler(w http.ResponseWriter, r *http.Request) {
	var bookings []models.Booking
	db.DB.Find(&bookings)
	json.NewEncoder(w).Encode(&bookings)
}

// GetBookingsHandler godoc
// @Summary Получение бронирования по id
// @Description Возвращает бронирование
// @Tags bookings
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Booking
// @Failure 400 {string} string "Bad Request"
// @Router /bookings/{id} [get]
func GetBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	existence := db.DB.First(&booking, id)

	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Booking not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}
	json.NewEncoder(w).Encode(&booking)
}

// UpdateBookingHandler godoc
// @Summary Изменение бронирования
// @Description Изменяет существующие бронирование
// @Tags bookings
// @Accept  json
// @Produce  json
// @Param booking body models.Booking true "Данные бронирования"
// @Success 200 {object} models.Booking
// @Failure 400 {string} string "Bad Request"
// @Router /bookings/{id} [put]
func UpdateBookingHandler(w http.ResponseWriter, r *http.Request) {
	var updatedBooking models.Booking
	var booking models.Booking
	params := mux.Vars(r)
	id, IdError := strconv.Atoi(params["id"])
	if IdError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(IdError.Error()))
		return
	}
	dataError := json.NewDecoder(r.Body).Decode(&updatedBooking)
	if dataError != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(dataError.Error()))
		return
	}

	existence := db.DB.First(&booking, id)
	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Booking not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}

	booking.StartTime = updatedBooking.StartTime
	booking.EndTime = updatedBooking.EndTime

	update := db.DB.Save(&booking)
	updateError := update.Error
	if updateError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(updateError.Error()))
		return
	}

	json.NewEncoder(w).Encode(&booking)
}

// GetBookingsHandler godoc
// @Summary Удаление бронирования по id
// @Description Удаляет бронирование
// @Tags bookings
// @Accept  json
// @Produce  json
// @Success 200 {string} string "OK"
// @Failure 400 {string} string "Bad Request"
// @Router /bookings/{id} [delete]
func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	var booking models.Booking
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	existence := db.DB.First(&booking, id)

	if existence.Error != nil {
		if existence.Error == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Booking not found"))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(existence.Error.Error()))
		}
		return
	}
	db.DB.Unscoped().Delete(&booking)
	w.WriteHeader(http.StatusOK)
}
