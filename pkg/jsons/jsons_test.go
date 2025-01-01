package jsons

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"io"
	"net/http"
	"testing"
)

func init() {
	// Set Gin to release mode to disable debug logging
	gin.SetMode(gin.ReleaseMode)
}

//easyjson:json
type TestStruct struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Email   string   `json:"email"`
	Tags    []string `json:"tags"`
	Enabled bool     `json:"enabled"`
}

// getTestData generates sample JSON data for testing
func getTestData() []byte {
	data := TestStruct{
		Name:    "test user",
		Age:     25,
		Email:   "test@example.com",
		Tags:    []string{"tag1", "tag2", "tag3"},
		Enabled: true,
	}
	b, _ := json.Marshal(data)
	return b
}

// BenchmarkGinShouldBindJSON benchmarks Gin's ShouldBindJSON method
func BenchmarkGinShouldBindJSON(b *testing.B) {
	// Create test context
	c, _ := gin.CreateTestContext(nil)
	testData := getTestData()

	b.ReportAllocs() // Report memory allocations
	b.ResetTimer()   // Reset timer before the actual benchmark

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		// Create new request for each iteration to simulate real-world usage
		c.Request = &http.Request{
			Body: io.NopCloser(bytes.NewBuffer(testData)),
		}
		if err := c.ShouldBindJSON(&obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkEasyJSON benchmarks easyjson's Unmarshal method
func BenchmarkEasyJSON(b *testing.B) {
	testData := getTestData()

	b.ReportAllocs() // Report memory allocations
	b.ResetTimer()   // Reset timer before the actual benchmark

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		if err := easyjson.Unmarshal(testData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkStdJSON benchmarks standard encoding/json package for comparison
func BenchmarkStdJSON(b *testing.B) {
	testData := getTestData()

	b.ReportAllocs() // Report memory allocations
	b.ResetTimer()   // Reset timer before the actual benchmark

	for i := 0; i < b.N; i++ {
		var obj TestStruct
		if err := json.Unmarshal(testData, &obj); err != nil {
			b.Fatal(err)
		}
	}
}
