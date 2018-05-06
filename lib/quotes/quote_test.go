// +build integration

package quote_test

import (
	"testing"

	"github.com/franela/goblin"
)

func Test(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("FindByText", func() {
		g.BeforeEach(func() {
			//
		})
		g.AfterEach(func() {
			//defer db.Close()
		})
		g.It("Should populate Quote struct with record found by text", func() {

		})
		g.It("Should leave default values to Quote struct if not found")
	})
	g.Describe("Save", func() {
		g.It("Should create a new record")
	})
	g.Describe("Exists", func() {
		g.It("Should return true if a quote exists")
		g.It("Should return false if a quote dees not exist")
	})
	g.Describe("Random", func() {
		g.It("Should return a random record")
	})
	g.Describe("Parse", func() {
		g.It("Should return error if unable to parse")
		g.It("Should return the quote struct if success")
	})
}
