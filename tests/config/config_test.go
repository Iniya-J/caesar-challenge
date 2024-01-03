Based on the explanation and test scenarios, here are the unit tests in Go:

```go
package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	// Scenario 1: Environment variables are set correctly
	os.Setenv("BASE_URL", "http://example.com/")
	os.Setenv("TOKEN_CODENATION", "12345")
	if BaseUrl != "http://example.com/" || GenerateUrl != "http://example.com/generate-data" || SubmitUrl != "http://example.com/submit-solution" || TokenCodeNation != "12345" {
		t.Errorf("Incorrect values for environment variables")
	}

	// Scenario 2: Environment variables are not set
	os.Clearenv()
	if BaseUrl != "" || GenerateUrl != "generate-data" || SubmitUrl != "submit-solution" || TokenCodeNation != "" {
		t.Errorf("Variables not initialized to empty strings when environment variables are not set")
	}

	// Scenario 3: Environment variables are set to unexpected values
	os.Setenv("BASE_URL", "Not a URL")
	os.Setenv("TOKEN_CODENATION", "Not a token")
	if BaseUrl != "Not a URL" || GenerateUrl != "Not a URLgenerate-data" || SubmitUrl != "Not a URLsubmit-solution" || TokenCodeNation != "Not a token" {
		t.Errorf("Incorrect values for environment variables when set to unexpected values")
	}

	// Scenario 4: Environment variables are set to edge case values
	os.Setenv("BASE_URL", "http://verylongurl.com/")
	os.Setenv("TOKEN_CODENATION", "verylongtoken")
	if BaseUrl != "http://verylongurl.com/" || GenerateUrl != "http://verylongurl.com/generate-data" || SubmitUrl != "http://verylongurl.com/submit-solution" || TokenCodeNation != "verylongtoken" {
		t.Errorf("Incorrect values for environment variables when set to edge case values")
	}

	// Edge Case 1: Environment variables are set to special characters
	os.Setenv("BASE_URL", "http://exa&mple.com/")
	os.Setenv("TOKEN_CODENATION", "123$%^")
	if BaseUrl != "http://exa&mple.com/" || GenerateUrl != "http://exa&mple.com/generate-data" || SubmitUrl != "http://exa&mple.com/submit-solution" || TokenCodeNation != "123$%^" {
		t.Errorf("Incorrect values for environment variables when set to special characters")
	}

	// Edge Case 2: Environment variables are set to non-ASCII characters
	os.Setenv("BASE_URL", "http://例子.com/")
	os.Setenv("TOKEN_CODENATION", "令牌")
	if BaseUrl != "http://例子.com/" || GenerateUrl != "http://例子.com/generate-data" || SubmitUrl != "http://例子.com/submit-solution" || TokenCodeNation != "令牌" {
		t.Errorf("Incorrect values for environment variables when set to non-ASCII characters")
	}

	// Edge Case 3: Environment variables are set to extremely long values
	longURL := "http://" + string(make([]byte, 2100)) + ".com/"
	longToken := string(make([]byte, 2100))
	os.Setenv("BASE_URL", longURL)
	os.Setenv("TOKEN_CODENATION", longToken)
	if BaseUrl != longURL || GenerateUrl != longURL+"generate-data" || SubmitUrl != longURL+"submit-solution" || TokenCodeNation != longToken {
		t.Errorf("Incorrect values for environment variables when set to extremely long values")
	}

	// Edge Case 4: Environment variables are set to values with leading or trailing whitespace
	os.Setenv("BASE_URL", " http://example.com/ ")
	os.Setenv("TOKEN_CODENATION", " 12345 ")
	if BaseUrl != " http://example.com/ " || GenerateUrl != " http://example.com/ generate-data" || SubmitUrl != " http://example.com/ submit-solution" || TokenCodeNation != " 12345 " {
		t.Errorf("Incorrect values for environment variables when set to values with leading or trailing whitespace")
	}
}
```
This test suite covers all the scenarios and edge cases mentioned. It checks the values of the global variables in the `config` package after setting the environment variables to different values. It uses the `os.Setenv` function to set the environment variables and `os.Clearenv` to clear them.