package handlers

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PudgeKim/go-holdem-authorization/application"
	"github.com/PudgeKim/go-holdem-authorization/domain/entity"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"math/rand"
	"net/http"
	"os"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	Endpoint:     google.Endpoint,
}

func (h *Handler) GoogleLoginHandler(c *gin.Context) {
	var state string

	cookie, err := c.Cookie("oauth")
	fmt.Println("cookie: ", cookie)
	if err != nil {
		state = getRandomState() // csrf 공격을 막기 위함
		c.SetCookie("oauth", state, 3600, "/", "localhost", false, false)
	} else {
		state = cookie
	}

	// 유저가 리다이렉트 해야할 주소를 가르쳐줌
	url := googleOauthConfig.AuthCodeURL(state)
	// 유저는 해당 주소로 리다이렉트한 후 구글에 로그인
	// 유저가 로그인을 성공했다면 구글에서 googleOauthConfig.RedirectURL에 해당하는 주소로 요청을 보냄
	// (그러므로 이 서버에서 해당 url에 대한 처리를 해줘야함)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func getRandomState() string {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	return state
}

func (h *Handler) GoogleAuthCallBack(c *gin.Context) {
	ctx := c.Request.Context()

	oauthCookie, err := c.Cookie("oauth")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no oauth cookie"})
		return
	}

	// 구글로부터 온 요청에 코드가 포함되어있음
	responseState := c.Query("state")
	fmt.Println("responseState: ", responseState)
	if responseState != oauthCookie {
		fmt.Println("invalid google oauth state cookie")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	// 포함된 코드를 통해 토큰을 받아낸 후
	// 해당 토큰으로 다시 구글에 요청하여 유저 데이터 얻어냄
	responseCode := c.Query("code")
	fmt.Println("responseCode: ", responseCode)
	userInfo, err := getGoogleUserInfo(responseCode)
	if err != nil {
		fmt.Println("getGoogleUserInfo error: ", err)
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	interactor := application.UserInteractor{Repository: h.userRepo}

	var user *entity.User

	user, err = interactor.GetUser(ctx, "123")
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			user = entity.NewUser(userInfo.Id, userInfo.GivenName, userInfo.Email)
			addUserErr := interactor.AddUser(ctx, user)
			if addUserErr != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"id":      user.Id,
		"name":    user.Name,
		"email":   user.Email,
		"balance": user.Balance,
	})
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

type GoogleUserReq struct {
	Id        string `json:"id"`
	GivenName string `json:"given_name"`
	Email     string `json:"email"`
}

func getGoogleUserInfo(code string) (*GoogleUserReq, error) {
	// 구글로부터 적절한 코드를 받고
	// 그 코드를 이용해 다시 구글에 요청을 해야 비로소 accessToken과 refreshToken을 받을 수 있음
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("new request error %s\n", err.Error())
	}

	var googleUserReq GoogleUserReq

	if err := json.NewDecoder(resp.Body).Decode(&googleUserReq); err != nil {
		return nil, err
	}

	return &googleUserReq, nil
}
