package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", 401)
			return
		}

		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "unauthorized", 401)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")

		if len(tokenParts) != 3 {
			http.Error(w, "unauthorized", 401)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		signature := tokenParts[2]

		fmt.Println(tokenParts)
		fmt.Println(jwtHeader)
		fmt.Println(jwtPayload)
		fmt.Println(signature)

		message := jwtHeader + "." + jwtPayload

		byteArrMessage := []byte(message)

		byteArrSecret := []byte(m.cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)

		newSignature := base64UrlEncode(hash)

		if signature != newSignature {
			http.Error(w, "Unauthorized tui hacker", 401)
			return
		}
	})
}
func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
