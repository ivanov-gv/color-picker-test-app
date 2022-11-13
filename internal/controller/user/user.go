package user

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"
	"github.com/sirupsen/logrus"
	"net/http"
)

type controller struct {
	ctx         context.Context
	log         logrus.FieldLogger
	userService service.UserInterface
	validate    *validator.Validate
}

func NewUserController(router *mux.Router, userService service.UserInterface, validator *validator.Validate,
	ctx context.Context, log logrus.FieldLogger) {
	c := &controller{
		ctx:         ctx,
		log:         log,
		userService: userService,
		validate:    validator,
	}

	router.HandleFunc("", c.create().ServeHTTP).Methods("POST")
}

func (c *controller) create() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		user, err := c.userService.Create(c.ctx)
		if err != nil {
			c.log.Errorf("can't create user: %w", err)
			http.Error(w, "can't create user: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(user.ToDto())
		if err != nil {
			c.log.Errorf("can't encode user to dto: %w", err)
			http.Error(w, "can't encode user to dto: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
	return http.HandlerFunc(fn)
}
