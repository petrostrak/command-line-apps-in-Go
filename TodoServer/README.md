#### Api TODO Server
This command-line tool connects to a REST API using Go's `net/http` package. It uses more advanced concepts such as the `http.Client` and `http.Request` types to fine-tune specific
connection parameters like headers and timeouts, and the `encoding/json` package to parse JSON response data. Finally, it covers several testing techniques to test the API server as well as the command-line client application, including local tests, simulated
responses, mock servers, and integration tests.

#### Testing the REST API Server
One approach for testing `HTTP` servers is testing each handler function individually by using the type `httptest.ResponseRecorder` . This type allows the recording of an `HTTP` response for analysis or tests. This approach is useful if you’re using the `DefaultServeMux` as the server multiplexer.

Because we implemented your own multiplexer function, `newMux` , we can use a different approach that allows integrated testing, including the route dispatching. We’ll use the type `httptest.Server` and instantiate a test server providing the multiplexer function as input. This approach creates a test server with an `URL` that simulates your server, allowing you to make requests similarly to using curl on the actual server. Then you can analyze and test the responses to ensure the server works as designed.