package dictionary_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-redis/redis/v8"
)

// SetDefinition sets a word-definition pair in Redis, but returns an error if the word or definition is empty.
func SetDefinition(ctx context.Context, rdb *redis.Client, word string, definition string) error {
	if len(word) == 0 {
		return errors.New("invalid word length")
	}
	if len(definition) == 0 {
		return errors.New("invalid definition length")
	}
	return rdb.Set(ctx, word, definition, 0).Err()
}

func TestAdd(t *testing.T) {
	// Create a Redis client for testing
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// Test success scenario
	err := SetDefinition(ctx, rdb, "test_word", "test_definition")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test failure scenario - invalid word length
	err = SetDefinition(ctx, rdb, "", "test_definition")
	if err == nil {
		t.Error("Expected error for invalid word length, but got nil")
	}
}
func TestGet(t *testing.T) {
	// Create a Redis client for testing
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// Add a test entry
	err := rdb.Set(ctx, "test_word", "test_definition", 0).Err()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test success scenario
	definition, err := rdb.Get(ctx, "test_word").Result()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if definition != "test_definition" {
		t.Errorf("Expected definition 'test_definition', but got '%s'", definition)
	}

	// Test failure scenario - word not found
	_, err = rdb.Get(ctx, "nonexistent_word").Result()
	if err == nil {
		t.Error("Expected error for nonexistent word, but got nil")
	}
}

func TestRemove(t *testing.T) {
	// Create a Redis client for testing
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.Background()

	// Add a test entry
	err := rdb.Set(ctx, "test_word", "test_definition", 0).Err()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test success scenario
	err = rdb.Del(ctx, "test_word").Err()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test failure scenario - word not found
	err = rdb.Del(ctx, "nonexistent_word").Err()
	if err == nil {
		t.Error("Expected error for nonexistent word, but got nil")
	}
}
