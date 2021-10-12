package app

import (
	"encoding/json"
	"net/http"

	"github.com/Lol-MBTI/data"
	"github.com/Lol-MBTI/model"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type AppHandler struct {
	http.Handler
	db model.DBHandler
}

type Success struct {
	Success bool `json:"success"`
}

var rd *render.Render = render.New()

func (a *AppHandler) Close() {
	a.db.Close()
}

func (a *AppHandler) GetMBTIHandler(rw http.ResponseWriter, r *http.Request) {
	var answer data.Answer
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&answer)
	if err != nil {
		rd.JSON(rw, http.StatusBadRequest, Success{false})
		return
	}

	var E, I, S, N, F, T, P, J int
	var mbti string
	for i := 0; i < 3; i++ {
		if answer.Response[i*4] == 1 {
			E += 1
		} else {
			I += 1
		}

		if answer.Response[i*4+1] == 1 {
			S += 1
		} else {
			N += 1
		}

		if answer.Response[i*4+2] == 1 {
			F += 1
		} else {
			T += 1
		}

		if answer.Response[i*4+3] == 1 {
			P += 1
		} else {
			J += 1
		}
	}

	if E > I {
		mbti += "e"
	} else {
		mbti += "i"
	}

	if S > N {
		mbti += "s"
	} else {
		mbti += "n"
	}

	if F > T {
		mbti += "f"
	} else {
		mbti += "t"
	}

	if P > J {
		mbti += "p"
	} else {
		mbti += "j"
	}

	champs := data.ChampData()
	goodMbti := data.GetGood(mbti)
	badMbti := data.GetBad(mbti)

	var res data.Result
	for _, champion := range champs {
		if champion.Line == answer.Line {
			if champion.MBTI == mbti {
				res.Name = champion.Name
			} else if champion.MBTI == goodMbti {
				res.GoodChamp = champion.Name
			} else if champion.MBTI == badMbti {
				res.BadChamp = champion.Name
			}
		}
	}

	rd.JSON(rw, http.StatusOK, res)
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Allow", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Next
		next.ServeHTTP(w, r)
		return
	})
}

func MakeHandler() *AppHandler {
	r := mux.NewRouter()
	r.Use(CORS)

	neg := negroni.Classic()
	neg.UseHandler(r)

	a := &AppHandler{
		Handler: neg,
		db:      model.NewDBHandler(),
	}

	r.HandleFunc("/mbti", a.GetMBTIHandler).Methods("POST")

	opts := middleware.RedocOpts{SpecURL: "/swagger.yml"}
	sh := middleware.Redoc(opts, nil)
	r.Handle("/docs", sh)
	r.Handle("/swagger.yml", http.FileServer(http.Dir("./")))

	return a
}
