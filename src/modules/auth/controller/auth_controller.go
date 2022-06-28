package controller

import "github.com/gin-gonic/gin"

type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GoogleOauth(ctx *gin.Context)
}

//https://accounts.google.com/o/oauth2/auth/oauthchooseaccount?
//access_type=offline&
//client_id=339378826083-17qb64dfntrdv4p5kuvu1lnp3t73vn2v.apps.googleusercontent.com&
//redirect_uri=https%3A%2F%2Fapi.edmodo.com%2Fauth%2Fgoogle_oauth2%2Fcallback&
//response_type=code&
//scope=email%20profile%20https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fdrive&
//state=71dc3accb7ebcab8a322d85d88b0a6ae123a4ed6a788ea6d&
//flowName=GeneralOAuthFlow
