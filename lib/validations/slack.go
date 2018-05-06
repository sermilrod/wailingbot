package validate

import (
	"errors"
)

// SlackToken verifies slack token and config token
func SlackToken(form map[string][]string, cfgToken string) (err error) {
	if _, ok := form["token"]; !ok {
		return errors.New("Missing token")
	}

	if form["token"][0] != cfgToken {
		return errors.New("Invalid token")
	}
	return nil
}

// SlackForm verifies that the bot has the info it needs to produce a response
func SlackForm(form map[string][]string, token string, getCmd string) (err error) {
	cmd, ok := form["command"]
	if !ok {
		return errors.New("Unable to parse command")
	}

	if cmd[0] != getCmd {
		if _, ok := form["text"]; !ok {
			return errors.New("Missing text")
		}

		if _, ok := form["user_name"]; !ok {
			return errors.New("Missing user_name")
		}
	}

	if err := SlackToken(form, token); err != nil {
		return err
	}

	return nil
}
