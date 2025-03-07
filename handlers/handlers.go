package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("secret_key")

var users = map[string]string{
	"ajay":     "ajaypwd",
	"raghu":    "raghupwd",
	"susan":    "susanpwd",
	"mike":     "mikepwd",
	"john":     "johnpwd",
	"emily":    "emilypwd",
	"robert":   "robertpwd",
	"lisa":     "lisapwd",
	"peter":    "peterpwd",
	"anna":     "annapwd",
	"tom":      "tompwd",
	"julia":    "juliapwd",
	"kevin":    "kevinpwd",
	"sophie":   "sophiepwd",
	"david":    "davidpwd",
	"rachel":   "rachelpwd",
	"charles":  "charlespwd",
	"mary":     "marypwd",
	"paul":     "paulpwd",
	"steve":    "stevepwd",
	"jack":     "jackpwd",
	"james":    "jamespwd",
	"karen":    "karenpwd",
	"sandra":   "sandrapwd",
	"ben":      "benpwd",
	"olivia":   "oliviapwd",
	"matt":     "mattpwd",
	"laura":    "laurapwd",
	"will":     "willpwd",
	"claire":   "clairepwd",
	"alex":     "alexpwd",
	"grace":    "gracepwd",
	"ian":      "ianpwd",
	"zoe":      "zoepwd",
	"daniel":   "danielpwd",
	"lucy":     "lucypwd",
	"george":   "georgepwd",
	"beth":     "bethpwd",
	"tim":      "timpwd",
	"isla":     "islapwd",
	"jim":      "jimpwd",
	"sarah":    "sarahpwd",
	"kyle":     "kylepwd",
	"vicky":    "vickypwd",
	"maggie":   "maggiepwd",
	"leah":     "leahpwd",
	"max":      "maxpwd",
	"joseph":   "josephpwd",
	"luke":     "lukepwd",
	"holly":    "hollypwd",
	"carl":     "carlpwd",
	"donna":    "donnapwd",
	"jenny":    "jennypwd",
	"greg":     "gregpwd",
	"amy":      "amypwd",
	"fiona":    "fionapwd",
	"brian":    "brianpwd",
	"toni":     "tonipwd",
	"lucas":    "lucaspwd",
	"olga":     "olgapwd",
	"morgan":   "morganpwd",
	"mia":      "miapwd",
	"nathan":   "nathanpwd",
	"ella":     "ellapwd",
	"sam":      "sampwd",
	"geoff":    "geoffpwd",
	"tina":     "tinapwd",
	"victor":   "victorpwd",
	"leon":     "leonpwd",
	"alison":   "alisonpwd",
	"hannah":   "hannahpwd",
	"roxy":     "roxypwd",
	"chris":    "chrispwd",
	"deborah":  "deborahpwd",
	"mark":     "markpwd",
	"nina":     "ninapwd",
	"hugo":     "hugopwd",
	"tess":     "tesspwd",
	"daisy":    "daisypwd",
	"aiden":    "aidenpwd",
	"rose":     "rosepwd",
	"simon":    "simonpwd",
	"helen":    "helenpwd",
	"natalie":  "nataliepwd",
	"chloe":    "chloepwd",
	"jackson":  "jacksonpwd",
	"quinn":    "quinnpwd",
	"penny":    "pennypwd",
	"ellen":    "ellenpwd",
	"sean":     "seanpwd",
	"toby":     "tobypwd",
	"jason":    "jasonpwd",
	"felix":    "felixpwd",
	"kim":      "kimpwd",
	"maria":    "mariapwd",
	"jasmine":  "jasminepwd",
	"josh":     "joshpwd",
	"jennifer": "jenniferpwd",
	"ron":      "ronpwd",
	"alina":    "alinapwd",
	"edward":   "edwardpwd",
	"jacob":    "jacobpwd",
	"valerie":  "valeriepwd",
	"hank":     "hankpwd",
	"nancy":    "nancypwd",
}

type Product struct {
	Name string
}
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Resp struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Login(response http.ResponseWriter, request *http.Request) {
	//fmt.Println("informatio reached here", request)
	var credentials Credentials
	err := json.NewDecoder(request.Body).Decode(&credentials)
	if err != nil {
		fmt.Println("error occured at ", err.Error())
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	expectedPassword, ok := users[credentials.Username]
	//fmt.Println("user expected password: ", credentials.Username, expectedPassword)

	if !ok || expectedPassword != credentials.Password {
		response.WriteHeader(http.StatusUnauthorized)
		//fmt.Println("expected password: ", ok, expectedPassword)
		return
	}

	expirationTime := time.Now().Add(2 * time.Minute)
	//fmt.Println("expiration time: ", expirationTime)

	claims := &Claims{
		Username: credentials.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	//fmt.Println("expirationTime and claims ", expirationTime, claims)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	fmt.Println("token", token)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Println("tokenString", tokenString)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(response, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	resp := Resp{
		Token: tokenString,
	}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(resp)
	if err != nil {
		// If there's an error encoding, return an internal server error
		fmt.Println("inside error encoding", err.Error())
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Println("returned at the end included cookie as well")
}
func Home(response http.ResponseWriter, request *http.Request) {
	// cookie, err := request.Cookie("token")

	// if err != nil {
	// 	if err == http.ErrNoCookie {
	// 		response.WriteHeader(http.StatusUnauthorized)
	// 		return
	// 	}
	// 	response.WriteHeader((http.StatusBadRequest))
	// 	return
	// }

	header := request.Header.Get("Authorization")
	tokenString := header
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			response.WriteHeader(http.StatusUnauthorized)
			return
		}
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		response.WriteHeader(http.StatusUnauthorized)
		return
	}

	response.Write([]byte(fmt.Sprintf("Hello , %s", claims.Username)))
	//response.WriteHeader(http.StatusInternalServerError)
	//return
}
func Refresh(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(http.StatusInternalServerError)
	return
}
