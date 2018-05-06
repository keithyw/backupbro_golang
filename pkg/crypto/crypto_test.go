package crypto_test

import (
	"testing"
	"backupBro/pkg/crypto"
)

func Test_Hash(t *testing.T) {
	t.Run("Can hash and compare", should_be_able_to_hash_and_compare_strings)
	t.Run("Can detect unequal hashes", should_return_error_when_comparing_unequal_hashes)
	t.Run("Generates a different salt every time", should_generate_a_different_salt_each_time)
}

func should_be_able_to_hash_and_compare_strings(t *testing.T) {
	c := crypto.Hash{}
	testInput := "testInput"

	generatedHash, generatedError := c.Generate(testInput)
	compareError := c.Compare(generatedHash, testInput)

	if generatedError != nil {
		t.Error("Error generating hash")
	}

	if testInput == generatedHash {
		t.Error("Generated hash is the same as the input")
	}

	if compareError != nil {
		t.Error("Error comparing hash to input")
	}
}

func should_return_error_when_comparing_unequal_hashes(t *testing.T) {
	c := crypto.Hash{}
	testInput := "testInput"
	testCompare := "testCompare"

	generatedHash, generatedError := c.Generate(testInput)
	compareError := c.Compare(generatedHash, testCompare)

	if generatedError != nil {
		t.Error("Error generating hash")
	}

	if testInput == generatedHash {
		t.Error("Generated hash is same as input")
	}

	if compareError == nil {
		t.Error("Compare should not have been successful")
	}
}

func should_generate_a_different_salt_each_time(t *testing.T) {
	c := crypto.Hash{}
	testInput := "testInput"

	hash1, _ := c.Generate(testInput)
	hash2, _ := c.Generate(testInput)

	if hash1 == hash2 {
		t.Error("Hash should not be equal")
	}
}