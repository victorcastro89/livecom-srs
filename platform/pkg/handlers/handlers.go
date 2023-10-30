package handlers

import "livecom/pkg/services"



type Handlers struct {
    Service *services.Service
}

func New(s *services.Service) Handlers {
    return Handlers{Service: s}
}
