package quote

import (
	"database/sql"
	"errors"
	"strings"
	"time"
)

// Quote represents the database schema
type Quote struct {
	ID        int
	Text      string
	Owner     string
	Date      time.Time
	CreatedAt time.Time
}

// FindByText queries the database looking for quotes maching a text pattern
func (q *Quote) FindByText(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT * FROM quotes WHERE text=$1", q.Text).
		Scan(&(q.ID), &(q.Text), &(q.Owner), &(q.Date), &(q.CreatedAt))
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		return err
	default:
		return nil
	}
}

// Save inserts a new record in the database
func (q *Quote) Save(db *sql.DB) (err error) {
	sqlStatement := `INSERT INTO quotes (id, text, owner, date, created_at)
									 VALUES (nextval('quotes_sequence'), $1, $2, $3, LOCALTIMESTAMP)
									 RETURNING id`
	err = db.QueryRow(sqlStatement, q.Text, q.Owner, q.Date).Scan(&(q.ID))
	if err != nil {
		return err
	}
	return nil
}

// Exists validates the existance of a quote
func (q *Quote) Exists(db *sql.DB) (res bool, err error) {
	if err := q.FindByText(db); err != nil {
		return false, err
	}
	if q.ID == 0 {
		return false, nil
	}
	return true, nil
}

// Random returns a random record from quotes
func (q *Quote) Random(db *sql.DB) (err error) {
	err = db.QueryRow("SELECT * FROM quotes ORDER BY random() limit 1").
		Scan(&(q.ID), &(q.Text), &(q.Owner), &(q.Date), &(q.CreatedAt))
	if err != nil {
		return err
	}
	return nil
}

// Parse maps the response to the Quote structure
func (q *Quote) Parse(text string) (err error) {
	command := strings.Fields(text)
	for index, element := range command {
		if index == 0 {
			if !strings.Contains(element, "@") {
				return errors.New("Parse error: Username missing")
			}
			q.Owner = element
		}
		if index == 1 {
			if strings.Contains(element, "[") {
				pDate := strings.Replace(strings.Replace(element, "[", "", -1), "]", "", -1)
				q.Date, err = time.Parse("2006-01-02", pDate)
				if err != nil {
					return errors.New("Parse error: Unable to parse quote date")
				}
				q.Text = strings.Join(command[index+1:], " ")
			} else {
				q.Text = strings.Join(command[index:], " ")
			}
			break
		}
	}
	return nil
}
