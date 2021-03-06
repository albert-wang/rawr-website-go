package routes

import (
	"log"
	"net/http"
)

func CheckAuth(r *http.Request, ctx Context) error {
	forwardedFor := r.Header["X-Forwarded-For"]

	//This is just the server's IP address.
	if len(forwardedFor) == 1 && forwardedFor[0] == ctx.Config.IPAddress {
		log.Print("Validated admin auth though ip address")
		return nil
	}

	auth, err := r.Cookie("auth")
	if err != nil {
		log.Print("Attempted to access a protected area, but failed to authenticate!")
		return MakeHttpError(err, http.StatusForbidden, r)
	}

	expected := ctx.Config.AdminSecret
	if auth.Value != expected {
		log.Print("Attempted to access a protected area, but failed to authenticate!")
		return MakeHttpError(nil, http.StatusForbidden, r)
	}

	return nil
}

func CheckAuthNonfatal(r *http.Request, ctx Context) error {
	forwardedFor := r.Header["X-Forwarded-For"]

	//This is just the server's IP address.
	if len(forwardedFor) == 1 && forwardedFor[0] == ctx.Config.IPAddress {
		return nil
	}

	auth, err := r.Cookie("auth")
	if err != nil {
		log.Print("No auth cookie?")
		return MakeHttpError(err, http.StatusForbidden, r)
	}

	expected := ctx.Config.AdminSecret
	if auth.Value != expected {
		log.Print("Secret mismatch")
		return MakeHttpError(nil, http.StatusForbidden, r)
	}

	return nil
}
