package controllers

import "github.com/trulyworthless/chatter-blocks/pkg/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/identities", middlewares.SetMiddlewareJSON(s.CreateIdentity)).Methods("POST")
	s.Router.HandleFunc("/identities", middlewares.SetMiddlewareJSON(s.GetIdentities)).Methods("GET")
	s.Router.HandleFunc("/identities/{uuid}", middlewares.SetMiddlewareJSON(s.GetIdentity)).Methods("GET")
	s.Router.HandleFunc("/identities/{uuid}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateIdentity))).Methods("PUT")
	s.Router.HandleFunc("/identities/{uuid}", middlewares.SetMiddlewareAuthentication(s.DeletIdentity)).Methods("DELETE")

	//Posts routes
	s.Router.HandleFunc("/contacts", middlewares.SetMiddlewareJSON(s.CreateContact)).Methods("POST")
	s.Router.HandleFunc("/contacts", middlewares.SetMiddlewareJSON(s.GetContacts)).Methods("GET")
	s.Router.HandleFunc("/contacts/{uuid}", middlewares.SetMiddlewareJSON(s.GetContact)).Methods("GET")
	s.Router.HandleFunc("/contacts/{uuid}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateContact))).Methods("PUT")
	s.Router.HandleFunc("/contacts/{uuid}", middlewares.SetMiddlewareAuthentication(s.DeleteContact)).Methods("DELETE")
}
