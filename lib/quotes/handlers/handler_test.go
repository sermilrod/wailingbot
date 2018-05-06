// +build integration

package quote_test

import (
	"testing"

	"github.com/sermilrod/wailingbot/lib/quotes/handlers"
	"github.com/franela/goblin"
)

var resType string = "in_channel"
var respText string = "test quote"

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Quote", func() {
		g.It("Should return 200 when creating a new quoute", func() {
		}) // should be 201, but slack API expects a 200
		g.It("Should return 401 when passing an invalid token")
		g.It("Should return 400 when not passing any token")
	})
	g.Describe("Random", func() {
		g.It("Should return 200 when fetching a random quote") // should be 201, but slack API expects a 200
	})
	g.Describe("sendResponse", func() {
		g.It("Should serialize quote Response as JSON", func() {
			r := quote.BuildResponse(resType, make([]quote.Attachements, 0, 1), respText)
			g.Assert(len(r.Attachements)).Equal(0)
			g.Assert(r.ResponseType).Equal(resType)
			g.Assert(r.Text).Equal(respText)
		})
	})
	g.Describe("checkRequest", func() {
		g.It("Should return nil if everything is OK")
		g.It("Should return 401 when passing an invalid token")
		g.It("Should return 400 when not passing any token")
	})
}
