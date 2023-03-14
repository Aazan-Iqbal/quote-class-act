// Filename: internal/models/responses.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// The Question model will represent a single question in our questions table
type Quote struct {
	QuoteID   int64
	Quote    string
	Author string
	CreatedAt    time.Time
}

// The QuestionModel type will encapsulate the
// DB connection pool that will be initialized
// in the main() function
type QuoteModel struct {
	DB *sql.DB
}

// The Insert() function stores a quote into the questions table
func (m *QuoteModel) Insert(response string, quote_id int64) (int64, error) {
	var id int64
	statement :=
		`
							INSERT INTO question_responses(question_id, response_code_id)
							VALUES($1, (SELECT response_code_id FROM response_codes WHERE code =  $2))
							RETURNING question_id
	            `
		// select represents looking for the phrase inputted so that we get the 'id'
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, quote_id, response).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// func (m *QuestionModel) Get() (*Question, error) {
// 	var q Question

// 	statement :=
// 	            `
// 							SELECT question_id, body
// 							FROM questions
// 							ORDER BY RANDOM()
// 							LIMIT 1
// 	            `
// 	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
// 	defer cancel()
// 	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.QuestionID, &q.Body)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &q, nil
//}
