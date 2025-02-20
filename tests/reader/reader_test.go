Here is the unit test code for the provided source code:

```go
package reader

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestReadAnswer(t *testing.T) {
	// Create a temporary file for testing
	tempFile, err := ioutil.TempFile("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write some data to the file
	data := []byte("Hello, World!")
	if _, err := tempFile.Write(data); err != nil {
		t.Fatal(err)
	}
	tempFile.Close()

	// Test reading the file
	readerAnswer, err := ReadAnswer(tempFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	// Check the data
	if string(readerAnswer.Data) != string(data) {
		t.Errorf("Expected %s, got %s", data, readerAnswer.Data)
	}

	// Check the file info
	if readerAnswer.Info.Name() != tempFile.Name() {
		t.Errorf("Expected %s, got %s", tempFile.Name(), readerAnswer.Info.Name())
	}

	// Test reading a non-existent file
	_, err = ReadAnswer("non_existent_file")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	// Test reading a directory
	_, err = ReadAnswer(os.TempDir())
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
```

This test code covers the following scenarios:

- Reading a valid file
- Reading a non-existent file
- Reading a directory

For the other scenarios, you would need to set up the appropriate conditions on your file system, which is beyond the scope of this example. However, you can use the same pattern as shown above to test those scenarios. For example, to test reading a file with different permissions, you could use the `os.Chmod` function to change the permissions of the test file before reading it.