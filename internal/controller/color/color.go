package color

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/ivanov-gv/color-picker-test-app/internal/model"
	"github.com/ivanov-gv/color-picker-test-app/internal/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type controller struct {
	ctx          context.Context
	log          logrus.FieldLogger
	colorService service.ColorInterface
	validate     *validator.Validate
}

func NewColorController(router *mux.Router, service service.ColorInterface, validator *validator.Validate,
	ctx context.Context, log logrus.FieldLogger) {
	c := &controller{
		ctx:          ctx,
		log:          log,
		colorService: service,
		validate:     validator,
	}

	router.HandleFunc("", c.add().ServeHTTP).Methods("POST")
	router.HandleFunc("", c.get().ServeHTTP).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", c.delete().ServeHTTP).Methods("DELETE")
}

func (c *controller) add() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		dto := &model.ColorDto{}
		err := json.NewDecoder(r.Body).Decode(&dto)
		if err != nil {
			c.log.Errorf("can't parse req: %w", err)
			http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = c.validate.StructCtx(c.ctx, dto)
		if err != nil {
			c.log.Errorf("invalid req: %w", err)
			http.Error(w, "bad json: "+err.Error(), http.StatusBadRequest)
			return
		}

		// vars user_id
		userIdStr := mux.Vars(r)["user_id"]
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.log.Errorf("invalid user_id: %w", err)
			http.Error(w, "bad path var user_id: "+err.Error(), http.StatusBadRequest)
			return
		}

		color := dto.FromDto()
		color, err = c.colorService.Add(c.ctx, userId, color)
		if err != nil {
			c.log.Errorf("can't add color: %w", err)
			http.Error(w, "can't add color: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(color.ToDto())
		if err != nil {
			c.log.Errorf("can't encode color to dto: %w", err)
			http.Error(w, "can't encode color to dto: "+err.Error(), http.StatusInternalServerError)
			return
		}

	}
	return http.HandlerFunc(fn)
}

func (c *controller) get() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userIdStr := mux.Vars(r)["user_id"]
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.log.Errorf("invalid user_id: %w", err)
			http.Error(w, "bad path var user_id: "+err.Error(), http.StatusBadRequest)
			return
		}

		colors, err := c.colorService.GetAll(c.ctx, userId)
		if err != nil {
			c.log.Errorf("can't get color: %w", err)
			http.Error(w, "can't get color: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = json.NewEncoder(w).Encode(model.ColorAll(colors).ToDto())
		if err != nil {
			c.log.Errorf("can't encode color to dto: %w", err)
			http.Error(w, "can't encode color to dto: "+err.Error(), http.StatusInternalServerError)
			return
		}

	}
	return http.HandlerFunc(fn)
}

func (c *controller) delete() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userIdStr := mux.Vars(r)["user_id"]
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			c.log.Errorf("invalid user_id: %w", err)
			http.Error(w, "bad path var user_id: "+err.Error(), http.StatusBadRequest)
			return
		}

		colorIdStr := mux.Vars(r)["id"]
		colorId, err := strconv.Atoi(colorIdStr)
		if err != nil {
			c.log.Errorf("invalid color id: %w", err)
			http.Error(w, "bad path var color id: "+err.Error(), http.StatusBadRequest)
			return
		}

		err = c.colorService.Delete(c.ctx, userId, colorId)
		if err != nil {
			c.log.Errorf("can't delete color: %w", err)
			http.Error(w, "can't delete color: "+err.Error(), http.StatusBadRequest)
			return
		}

	}
	return http.HandlerFunc(fn)
}
