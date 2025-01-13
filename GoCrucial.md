Method validation at the start of each request:

if r.Method != http.MethodDelete {
    w.WriteHeader(http.StatusMethodNotAllowed)
    w.Write([]byte(fmt.Sprintf(`Cannot %s %s`, r.Method, r.URL.Path)))
    return
}           

Usage of fmt.Sprintf
fmt.Sprintf is used to format a string with placeholders and return the formatted string without printing it directly.

queryValues := r.URL.Query()
r.URL.Query():
Extracts the query string from the URL of the HTTP request.
Returns it as a url.Values object, which is essentially a map-like structure (map[string][]string) where:

Example: For a URL like:

http://example.com/api/resource?userId=123&role=admin
queryValues will be:

map[string][]string{
    "userId": {"123"},
    "role": {"admin"},
}

Core Logging Methods

Debug
Logs detailed debug-level information, useful for development and debugging.
Example: logger.Debug("This is a debug message")
Info
Logs general informational messages to indicate normal operations.
Example: logger.Info("Application started successfully")
Warn
Logs a warning message about something unexpected but not critical.
Example: logger.Warn("Disk space is running low")
Error
Logs an error message when something goes wrong but the application can continue.
Example: logger.Error("Failed to connect to the database")
Fatal
Logs a critical error message and then terminates the application.
Example: logger.Fatal("Unable to start server on port 8080")
Typically followed by os.Exit().
Panic
Logs a critical error message and then panics (causes the program to crash).
Example: logger.Panic("Unexpected null pointer exception")

logger.Log
This is the base logger instance, usually created and configured at the beginning of the application.


io.ReadAll(r io.Reader) ([]byte, error)
Reads all data from an io.Reader until EOF and returns it as a byte slice.
Example:
data, err := io.ReadAll(file)
fmt.Println(string(data)) // Prints file content

bodyByte, err := io.ReadAll(r.Body)


io.ReadAll(r.Body)
r.Body is an io.Reader that represents the body of an HTTP request. This might contain data like JSON, form data, or plain text sent by a client.
io.ReadAll reads the entire content of the r.Body until EOF (end of file) and returns it as a byte slice ([]byte).
Use this when you need to process the entire request body at once.
bodyByte
The result of io.ReadAll is assigned to bodyByte, which is a byte slice ([]byte).
err
If an error occurs during reading (e.g., if the body is too large or the connection is interrupted), err will hold the error. Always check and handle it properly.

When to Use:
Always use defer r.Body.Close() when you read from r.Body in an HTTP handler.


if err := json.Unmarshal(bodyByte, &body); err != nil {
    log.Errorf(`failed to unmarshal the request body : ERROR : %s`, err)
    utils.SendInternalServerErrorResponse(w)
    return
}

Parses JSON Data:
It attempts to convert (unmarshal) the raw JSON byte array (bodyByte) into a Go data structure (body).

err = validator.New().Struct(body)
	if err != nil {
		log.Errorf(`Expected fields are missing on the request body : ERROR : %s`, err.Error())
		utils.SendBadRequestResponse(w, "Expected fields are missing on the request body")
		return
	}

    validator.New()
    Creates a new instance of a validator to validate struct fields based on tags in the struct definition.
    .Struct(body)
    Method: Struct validates the body object against its defined struct tags.


Marshalling: Convert Go objects to JSON (e.g., for sending responses or saving data).
Unmarshalling: Convert JSON to Go objects (e.g., when processing incoming requests or reading data).