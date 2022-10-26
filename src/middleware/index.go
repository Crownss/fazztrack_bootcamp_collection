package middleware

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/libs"
	"github.com/golang-jwt/jwt/v4"
)

var signatureKEY = []byte(os.Getenv("JWT_KEY"))

type Claims struct {
	Username string
	Is_Admin bool
	jwt.StandardClaims
}

func NewToken(username string, is_admin bool) *Claims {
	return &Claims{
		Username: username,
		Is_Admin: is_admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(signatureKEY)
}

func CheckToken(token string) (*Claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return signatureKEY, nil
	})
	if err != nil {
		return nil, err
	}
	claim := tokens.Claims.(*Claims)
	return claim, err
}

func CheckTokenString(token string) (string, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return signatureKEY, nil
	})
	if err != nil {
		return "", err
	}
	claim := tokens.Claims.(*Claims)
	return claim.Username, err
}

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		if !strings.Contains(headerToken, "Bearer") {
			libs.Response(w, "authorization invalid!", 500, true)
			return
		}
		tokens := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := CheckToken(tokens)
		if err != nil {
			libs.Response(w, err.Error(), 500, true)
			return
		}
		ctx := context.WithValue(r.Context(), "username", checkToken.Username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func CheckAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		if !strings.Contains(headerToken, "Bearer") {
			libs.Response(w, "authorization invalid!", 500, true)
			return
		}
		tokens := strings.Replace(headerToken, "Bearer ", "", -1)
		checkToken, err := CheckToken(tokens)
		if err != nil {
			libs.Response(w, err.Error(), 500, true)
			return
		}
		ctx := context.WithValue(r.Context(), "is_admin", checkToken.Is_Admin)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func UploadProfile(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err1 := r.ParseMultipartForm(10 << 20)
		if err1 != nil {
			libs.Response(w, err1.Error(), 500, true)
			return
		}
		file, _, err := r.FormFile("profile")
		if err != nil {
			libs.Response(w, err.Error(), 500, true)
			return
		}
		defer file.Close()
		tempFile, err := ioutil.TempFile("images", "upload-*.png")
		if err != nil {
			libs.Response(w, err.Error(), 500, true)
			return
		}
		defer tempFile.Close()
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			libs.Response(w, err.Error(), 500, true)
			return
		}
		tempFile.Write(fileBytes)
		ctx := context.WithValue(r.Context(), "filename", tempFile.Name())
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

type Middleware func(http.HandlerFunc) http.HandlerFunc

func MultipleMiddleware(hf http.HandlerFunc, middle ...Middleware) http.HandlerFunc {

	for _, v := range middle {
		hf = v(hf)
	}
	return hf
}

func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		(w).Header().Set("Access-Control-Allow-Origin", "*")
		(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}
