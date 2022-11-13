package integration_test

import (
	"fmt"
	. "github.com/Eun/go-hit"
	"log"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const (
	host        = "http://app:8080"
	probe       = "/healtz"
	userHandler = "/user"
)

var userId int

func TestMain(m *testing.M) {
	if err := Do(Get(host+probe), Expect().Status().Equal(http.StatusOK)); err != nil {
		log.Fatalf("Integration tests: host %s is not available: %s", host, err)
	}

	err := Do(Description("Test user service"),
		Post(host+userHandler),
		Send().Body().String(""),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().JSON().JQ(".id").In(&userId))

	if err != nil {
		log.Fatalf("Integration tests: can't create user for tests: %s", err)
	}

	os.Exit(m.Run())
}

func TestColor(t *testing.T) {
	colorHandler := fmt.Sprintf("/user/%d/color", userId)

	color := map[string]any{
		"name": "my happy color",
		"HEX":  "#000000"}
	var colorId int

	Test(t,
		Description("Color create"),
		Post(host+colorHandler),
		Send().Body().JSON(color),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().JSON().JQ(".id").In(&colorId))

	color["id"] = colorId

	Test(t, Description("Color get"),
		Get(host+colorHandler),
		Send().Body().String(""),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(map[string]any{
			"colors": []map[string]any{color}}))

	Test(t,
		Description("Color create the same color"),
		Post(host+colorHandler),
		Send().Body().JSON(color),
		Expect().Status().Equal(http.StatusBadRequest))

	Test(t, Description("Color delete"),
		Delete(host+colorHandler+"/"+strconv.Itoa(colorId)),
		Send().Body().String(""),
		Expect().Status().Equal(http.StatusOK))

	Test(t, Description("Color get empty list"),
		Get(host+colorHandler),
		Send().Body().String(""),
		Expect().Status().Equal(http.StatusOK),
		Expect().Body().JSON().Equal(map[string]any{
			"colors": []map[string]any{}}))
}

func TestColorValidation(t *testing.T) {
	colorHandler := fmt.Sprintf("/user/%d/color", userId)
	var colorId int

	color := map[string]any{
		"name": "my happy color",
		"HEX":  "#000000"}

	colorSameName := map[string]any{
		"name": "my happy color",
		"HEX":  "#000001"}

	colorSameHex := map[string]any{
		"name": "my happy color!",
		"HEX":  "#000000"}

	colorWrongHex := map[string]any{
		"name": "my happy color!",
		"HEX":  "not a color"}

	colorEmptyName := map[string]any{
		"name": "",
		"HEX":  "#000000"}

	colorLongName := map[string]any{
		"name": "too long name for the color to create",
		"HEX":  "#000000"}

	wrongColors := []map[string]any{colorSameName, colorSameHex, colorWrongHex, colorEmptyName, colorLongName}

	Test(t,
		Description("Color create"),
		Post(host+colorHandler),
		Send().Body().JSON(color),
		Expect().Status().Equal(http.StatusOK),
		Store().Response().Body().JSON().JQ(".id").In(&colorId))

	for c := range wrongColors {
		Test(t,
			Description("Color create wrong colors"),
			Post(host+colorHandler),
			Send().Body().JSON(c),
			Expect().Status().Equal(http.StatusBadRequest))
	}

	Test(t, Description("Color delete"),
		Delete(host+colorHandler+"/"+strconv.Itoa(colorId)),
		Send().Body().String(""),
		Expect().Status().Equal(http.StatusOK))
}
