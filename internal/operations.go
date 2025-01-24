package token

import (
	"database/sql"
	"fmt"

	db "github.com/vivek-344/token-management-system/db/queries"
	"github.com/vivek-344/token-management-system/models"
)

// InitializeTokens initializes the tokens in the database.
// If tokens already exist, it resets their usage counts if they weren't updated today.
func InitializeTokens(q *db.Queries) error {
	// Check if tokens are already initialized
	count, err := q.GetTokenCount()
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if count == 0 {
		// Insert 1000 tokens
		for i := 1; i <= 1000; i++ {
			tokenID := fmt.Sprintf("Token %d", i)
			err = q.AddToken(tokenID)
			if err != nil {
				return err
			}
		}

		fmt.Println("Tokens initialized.")
	} else {
		// Reset usage counts
		err = q.ResetUsageCount()
		if err != nil {
			return err
		}
	}
	return nil
}

// SimulateOperations simulates the token usage operations
func SimulateOperations(q *db.Queries, operations int) error {
	for i := 0; i < operations; i++ {
		// Select token with least usage count
		token, err := selectLeastUsedToken(q)
		if err != nil {
			return err
		}

		// Increment token usage count
		err = q.IncrementTokenUsage(token.TokenID)
		if err != nil {
			return err
		}
	}
	return nil
}

// selectLeastUsedToken selects a token with the least usage count.
// If multiple tokens have the same least usage count, it selects one at random.
func selectLeastUsedToken(q *db.Queries) (*models.Token, error) {
	// Find the minimum usage count
	minUsage, err := q.GetLeastUsageCount()
	if err != nil {
		return &models.Token{}, err
	}

	// Return a random token with least usage count
	return q.GetTokenWithLeastUsage(minUsage)
}

// DisplayResults displays the usage counts for each token and identifies the least-used tokens
func DisplayResults(q *db.Queries) error {
	fmt.Println("\nToken Usage Counts:")
	tokens, err := q.GetTokensByUsageCount()
	if err != nil {
		return err
	}

	usageMap := make(map[int][]string)
	minUsage := -1

	for _, token := range tokens {
		fmt.Printf("%s: %d uses\n", token.TokenID, token.UsageCount)

		usageMap[token.UsageCount] = append(usageMap[token.UsageCount], token.TokenID)
		if minUsage == -1 || token.UsageCount < minUsage {
			minUsage = token.UsageCount
		}
	}

	// Identify least-used tokens
	leastUsedTokens := usageMap[minUsage]
	fmt.Println("\nLeast Used Token(s):")
	for _, tokenID := range leastUsedTokens {
		fmt.Printf("%s (%d uses)\n", tokenID, minUsage)
	}

	return nil
}
