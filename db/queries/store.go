package db

import (
	"database/sql"
	"time"

	"github.com/vivek-344/token-management-system/models"
)

// Queries handles database operations for tokens
type Queries struct {
	db *sql.DB
}

// NewQueries creates a new Queries instance
func NewQueries(db *sql.DB) *Queries {
	return &Queries{db: db}
}

// GetTokenCount fetches the total number of tokens
func (q *Queries) GetTokenCount() (int, error) {
	var count int
	err := q.db.QueryRow(`SELECT COUNT(*) FROM token`).Scan(&count)
	return count, err
}

// GetAllTokens returns a slice of all the tokens
func (q *Queries) GetAllTokens() ([]models.Token, error) {
	getAllTokens := `SELECT token_id, usage_count, last_updated FROM token`
	rows, err := q.db.Query(getAllTokens)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []models.Token{}
	for rows.Next() {
		var token models.Token
		err := rows.Scan(&token.TokenID, &token.UsageCount, &token.LastUpdated)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

// GetTokensByUsageCount returns a slice of all the token sorted by usage count
func (q *Queries) GetTokensByUsageCount() ([]models.Token, error) {
	getTokensByUsageCount := `SELECT token_id, usage_count FROM token ORDER BY usage_count DESC, token_id ASC`
	rows, err := q.db.Query(getTokensByUsageCount)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []models.Token{}
	for rows.Next() {
		var token models.Token
		err := rows.Scan(&token.TokenID, &token.UsageCount)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

// GetLeastUsageCount returns the minimum usage count
func (q *Queries) GetLeastUsageCount() (int, error) {
	getLeastUsageCount := `SELECT MIN(usage_count) FROM token`

	var leastUsageCount int
	err := q.db.QueryRow(getLeastUsageCount).Scan(&leastUsageCount)
	if err != nil {
		return 0, err
	}

	return leastUsageCount, nil
}

// GetTokensWithLeastUsage returns a random token with least usage count
func (q *Queries) GetTokenWithLeastUsage(minUsage int) (*models.Token, error) {
	getTokensWithLeastUsage := `SELECT token_id, usage_count, last_updated FROM token WHERE usage_count = $1 ORDER BY RANDOM() LIMIT 1`

	var token models.Token
	err := q.db.QueryRow(getTokensWithLeastUsage, minUsage).Scan(&token.TokenID, &token.UsageCount, &token.LastUpdated)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &token, nil
}

// AddToken inserts a new token into the database
func (q *Queries) AddToken(tokenId string) error {
	addToken := `INSERT INTO token (token_id, usage_count, last_updated) VALUES ($1, $2, $3)`
	_, err := q.db.Exec(addToken, tokenId, 0, time.Now())
	return err
}

// ResetUsageCount resets the usage count to 0 if they weren't updated today.
func (q *Queries) ResetUsageCount() error {
	resetUsage := `UPDATE token SET usage_count = 0, last_updated = $1 WHERE DATE(last_updated) != DATE($1)`
	_, err := q.db.Exec(resetUsage, time.Now())
	return err
}

// IncrementTokenUsage increments the usage count of the specified token
func (q *Queries) IncrementTokenUsage(tokenID string) error {
	incrementTokenUsage := `UPDATE token SET usage_count = usage_count + 1, last_updated = $1 WHERE token_id = $2`
	_, err := q.db.Exec(incrementTokenUsage, time.Now(), tokenID)
	return err
}
