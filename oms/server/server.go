package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"oms/oms/core"
)

type Server struct {
	Core core.Core
}

// func (s *Server) HandleOrder() func(rw http.ResponseWriter, rq *http.Request) {
// 	return func(rw http.ResponseWriter, rq *http.Request) {
// 		switch rq.Method {
// 		case http.MethodGet:
// 			s.GetOrder(rq, rw)
// 		case http.MethodPost:
// 			s.CreateOrder(rq, rw)
// 		case http.MethodPut:
// 			s.UpdateOrder(rq, rw)
// 		}

// 	}
// }

// func (s *Server) CreateOrder(rq *http.Request, rw http.ResponseWriter) {
// 	defer rq.Body.Close()
// 	var payload payloadEntity.CreateOrderRequest
// 	err := json.NewDecoder(rq.Body).Decode(&payload)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	orderDetails, err := s.Core.CreateOrder(payload)
// 	if err != nil {
// 		fmt.Println("error during create order - ", err)
// 		rw.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	rw.WriteHeader(http.StatusOK)
// 	json.NewEncoder(rw).Encode(orderDetails)
// }

// func (s *Server) GetOrder(rq *http.Request, rw http.ResponseWriter) {
// 	defer rq.Body.Close()
// 	var payload payloadEntity.GetOrderRequest
// 	err := json.NewDecoder(rq.Body).Decode(&payload)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	orderDetails, err := s.Core.GetOrder(payload)
// 	if err != nil {
// 		fmt.Println("error during get order - ", err)
// 		rw.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	rw.WriteHeader(http.StatusOK)
// 	json.NewEncoder(rw).Encode(orderDetails)
// }

// func (s *Server) UpdateOrder(rq *http.Request, rw http.ResponseWriter) {
// 	defer rq.Body.Close()
// 	var payload payloadEntity.UpdateOrderRequest
// 	err := json.NewDecoder(rq.Body).Decode(&payload)
// 	if err != nil {
// 		rw.WriteHeader(http.StatusBadRequest)
// 		return
// 	}
// 	err = s.Core.UpdateOrder(payload)
// 	if err != nil {
// 		fmt.Println("error during update order - ", err)
// 		rw.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	rw.WriteHeader(http.StatusOK)
// 	rw.Write([]byte("Order updated successfully"))
// }

func (s *Server) HandleOrderMetrics() func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		switch rq.Method {
		case http.MethodGet:
			s.GetOrderMetrics(rq, rw)
		}
	}
}

func (s *Server) GetOrderMetrics(rq *http.Request, rw http.ResponseWriter) {
	orderMetrics, err := s.Core.GetOrderMetrics()
	if err != nil {
		fmt.Println("error during get order metrics - ", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(orderMetrics)
}
