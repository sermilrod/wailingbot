package quote

import (
	"database/sql"
	"net/http"

	config "github.com/sermilrod/wailingbot/lib/configuration"
	quote "github.com/sermilrod/wailingbot/lib/quotes"
	"github.com/sermilrod/wailingbot/lib/validations"
	"github.com/labstack/echo"
)

// Attachements represent extra attachements to command response
type Attachements map[string]string

// Response formats a valid response for Slack API
type Response struct {
	Attachements []Attachements `json:"attachements"`
	ResponseType string         `json:"response_type"` // in_channel (visible to all channel) or ephemeral (visible to user requesting)
	Text         string         `json:"text"`
}

// QuoteHandler creates a new closure that persists quotes
func QuoteHandler(cfg *config.Configuration, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		respText := "Quote already in the wailing wall"
		quote := quote.Quote{}
		if msg, err := CheckRequest(c, cfg); err != nil {
			r := BuildResponse("ephemeral", make([]Attachements, 0, 1), msg)
			return c.JSONPretty(http.StatusOK, r, "  ")
		}
		err := quote.Parse(c.FormValue("text"))
		if err != nil {
			msg := "Unable to parse the quote: " + err.Error()
			r := BuildResponse("ephemeral", make([]Attachements, 0, 1), msg)
			return c.JSONPretty(http.StatusOK, r, "  ")
		}
		q, err := quote.Exists(db)
		if err != nil {
			msg := "Unable to validate quote existence: " + err.Error()
			r := BuildResponse("ephemeral", make([]Attachements, 0, 1), msg)
			return c.JSONPretty(http.StatusOK, r, "  ")
		}
		if !q {
			if err := quote.Save(db); err != nil {
				msg := "Unable to persist the quote: " + err.Error()
				r := BuildResponse("ephemeral", make([]Attachements, 0, 1), msg)
				return c.JSONPretty(http.StatusOK, r, "  ")
			}
			respText = "New quote added"
		}

		attcs, attc := make([]Attachements, 0, 1), make(Attachements)
		attc["text"] = "Quote owner: " + quote.Owner + ", quote date: " + quote.Date.Format("2006-01-02")
		attcs = append(attcs, attc)
		r := BuildResponse("ephemeral", attcs, respText)
		return c.JSONPretty(http.StatusOK, r, "  ")
	}
}

// RandomHandler returns a random quote from the database
func RandomHandler(cfg *config.Configuration, db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		quote := quote.Quote{}
		if msg, err := CheckRequest(c, cfg); err != nil {
			r := BuildResponse("in_channel", make([]Attachements, 0, 1), msg)
			return c.JSONPretty(http.StatusOK, r, "  ")
		}
		if err := quote.Random(db); err != nil {
			msg := "Unable to retrieve a random record: " + err.Error()
			r := BuildResponse("in_channel", make([]Attachements, 0, 1), msg)
			return c.JSONPretty(http.StatusOK, r, "  ")
		}
		r := BuildResponse("in_channel", make([]Attachements, 0, 1), quote.Text+", by "+quote.Owner+" "+quote.Date.Format("2006-01-02"))
		return c.JSONPretty(http.StatusOK, r, "  ")
	}
}

// BuildResponse populates Response struct
func BuildResponse(resType string, attcs []Attachements, respText string) (r Response) {
	return Response{Attachements: attcs, ResponseType: resType, Text: respText}
}

// CheckRequest performs a set of validations before letting the handler do operations
func CheckRequest(c echo.Context, cfg *config.Configuration) (msg string, err error) {
	params := make(map[string][]string)
	params["token"] = append(params["token"], c.FormValue("token"))
	params["text"] = append(params["text"], c.FormValue("text"))
	params["user_name"] = append(params["user_name"], c.FormValue("user_name"))
	params["command"] = append(params["command"], c.FormValue("command"))
	if err := validate.SlackForm(params, cfg.SlackToken, cfg.GetCmd); err != nil {
		return "Form validation failed: " + err.Error(), err
	}
	return msg, nil
}
