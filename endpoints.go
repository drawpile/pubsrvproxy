package main

import (
	"github.com/drawpile/pubsrvproxy/queries"
	"log"
	"net/http"
)

/**
 * Server info
 */
type ServerInfoResponse struct {
	ApiName     string `json:"api_name"`
	Version     string `json:"version"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Favicon     string `json:"favicon"`
	ReadOnly    bool   `json:"read_only"`
}

func (r ServerInfoResponse) WriteResponse(w http.ResponseWriter) {
	writeJsonResponse(w, r, http.StatusOK)
}

func ServerInfoEndpoint(ctx *apiContext) apiResponse {
	if len(ctx.path) == 0 {
		return ServerInfoResponse{
			ApiName:     "drawpile-session-list",
			Version:     "1.5",
			Name:        ctx.cfg.Name,
			Description: ctx.cfg.Description,
			Favicon:     ctx.cfg.Favicon,
			ReadOnly:    true,
		}
	} else {
		return notFoundResponse()
	}
}

/**
 * List sessions
 *
 * The returned response is compatible with the one that the listserver returns.
 */

type SessionListResponse struct {
	sessions []queries.SessionInfo
}

func (r SessionListResponse) WriteResponse(w http.ResponseWriter) {
	writeJsonResponse(w, r.sessions, http.StatusOK)
}

func SessionListEndpoint(ctx *apiContext) apiResponse {
	if len(ctx.path) == 0 {
		list, err := queries.QuerySessionInfo(ctx.GetQueryOpts())
		if err != nil {
			log.Println("Session listing error:", err)
			return internalServerError()
		}

		for i, _ := range list {
			list[i].Host = ctx.cfg.ServerHost
			list[i].Port = ctx.cfg.ServerPort
		}

		return SessionListResponse{list}

	} else {
		return notFoundResponse()
	}
}

/**
 * List users
 */
type UserListResponse struct {
	users []queries.UserInfo
}

func (r UserListResponse) WriteResponse(w http.ResponseWriter) {
	writeJsonResponse(w, r.users, http.StatusOK)
}

func UserListEndpoint(ctx *apiContext) apiResponse {
	if len(ctx.path) == 0 {
		list, err := queries.QueryUserList(ctx.GetQueryOpts())
		if err != nil {
			log.Println("User listing error:", err)
			return internalServerError()
		}

		if ctx.cfg.ShowUserIps == false {
			// Redact IP addresses
			for i, _ := range list {
				list[i].Ip = ""
			}
		}

		return UserListResponse{list}

	} else {
		return notFoundResponse()
	}
}
