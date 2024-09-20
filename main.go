package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"github.com/theinlaoq/booking-api-testcase/db"
	_ "github.com/theinlaoq/booking-api-testcase/docs"
	"github.com/theinlaoq/booking-api-testcase/routes"
)

// @title Booking API
// @version 1.0
// @description Это API для бронирования игровых мест в компьютерном клубе.
// @host localhost:3000
// @BasePath /

func main() {
	db.DBConnection()
	r := mux.NewRouter()

	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.UpdateUserHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	r.HandleFunc("/bookings", routes.PostBookingHandler).Methods("POST")
	r.HandleFunc("/bookings", routes.GetBookingsHandler).Methods("GET")
	r.HandleFunc("/bookings/{id}", routes.GetBookingHandler).Methods("GET")
	r.HandleFunc("/bookings/{id}", routes.UpdateBookingHandler).Methods("PUT")
	r.HandleFunc("/bookings/{id}", routes.DeleteBookingHandler).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("server is running")
	log.Fatal(http.ListenAndServe(":3000", r))

}
