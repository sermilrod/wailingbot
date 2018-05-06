// +build unit

package validate_test

import (
	"testing"

	"github.com/sermilrod/wailingbot/lib/validations"
	"github.com/franela/goblin"
)

var testToken string = "abcd1234"
var testText string = "test quote"
var testUserName string = "sermilrod"
var testGetCmd string = "/wwget"
var testAddCmd string = "/wwadd"

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SlackToken", func() {
		g.It("Should return nil if valid token", func() {
			x := make(map[string][]string)
			x["token"] = append(x["token"], testToken)
			g.Assert(validate.SlackToken(x, testToken)).Equal(nil)
		})
		g.It("Should return error if missing token", func() {
			x := make(map[string][]string)
			err := validate.SlackToken(x, testToken)
			g.Assert(err.Error()).Equal("Missing token")
		})
		g.It("Should return error if invalid token", func() {
			x := make(map[string][]string)
			x["token"] = append(x["token"], "invalidtoken")
			err := validate.SlackToken(x, testToken)
			g.Assert(err.Error()).Equal("Invalid token")
		})
	})
	g.Describe("SlackForm", func() {
		g.It("Should return nil if valid form", func() {
			x := make(map[string][]string)
			x["token"] = append(x["token"], testToken)
			x["text"] = append(x["text"], testText)
			x["user_name"] = append(x["user_name"], testUserName)
			x["command"] = append(x["command"], testAddCmd)
			g.Assert(validate.SlackForm(x, testToken, testGetCmd)).Equal(nil)
		})
		g.It("Should return error if token validation fails", func() {
			x := make(map[string][]string)
			x["token"] = append(x["token"], "invalidtoken")
			x["command"] = append(x["command"], testAddCmd)
			err := validate.SlackForm(x, testToken, testAddCmd)
			g.Assert(err.Error()).Equal("Invalid token")
		})
		g.It("Should return error if command is missing", func() {
			x := make(map[string][]string)
			err := validate.SlackForm(x, testToken, testGetCmd)
			g.Assert(err.Error()).Equal("Unable to parse command")
		})
		g.It("Should return error if text is missing", func() {
			x := make(map[string][]string)
			x["command"] = append(x["command"], testAddCmd)
			err := validate.SlackForm(x, testToken, testGetCmd)
			g.Assert(err.Error()).Equal("Missing text")
		})
		g.It("Should return error if user_name is missing", func() {
			x := make(map[string][]string)
			x["text"] = append(x["text"], testText)
			x["command"] = append(x["command"], testAddCmd)
			err := validate.SlackForm(x, testToken, testGetCmd)
			g.Assert(err.Error()).Equal("Missing user_name")
		})
	})
}
