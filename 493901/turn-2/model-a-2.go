// BenchmarkHttpServer benchmarks the HTTP server handler.
// It measures the time it takes to handle a single request at /hello.
// This benchmark simulates multiple client connections to see how the server performs.
func BenchmarkHttpServer(b *testing.B) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	client := &http.Client{}

	for i := 0; i < b.N; i++ {
		_, err := client.Get(ts.URL + "/hello")
		if err != nil {
			b.Errorf("Error performing request: %v", err)
		}
	}
}
