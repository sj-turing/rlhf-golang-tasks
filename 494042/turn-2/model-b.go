package main

type AppError struct {
	Code    int
	Message string
	Cause   error // Store the underlying error if present
}

func NewAppError(code int, message string) error {
	return &AppError{Code: code, Message: message}
}
func LoginHandler(username, password string, callback ErrorCallback) (bool, error) {
	// Validate input, check against database, etc.
	if err := ValidateCredentials(username, password); err != nil {
		callback(err)
		return false, err
	}
	// Authenticate and return success
	return true, nil
}

func (callback ErrorCallback) Log(err error) {
	log.Println("Error:", err)
	if appErr, ok := err.(*AppError); ok {
		log.Println("Code:", appErr.Code)
	}
	stack := make([]byte, 1024)
	runtime.Stack(stack, true)
	log.Println("Stack Trace:", string(stack))
}

func UnsafeQuery(query string, callback ErrorCallback) *sql.Rows {
	rows, err := db.Query(query) // Vulnerable to SQL injection
	if err != nil {
		callback(err)
		return nil, err
	}
	return rows, nil
}

func UnsafeTemplate(tpl string, data map[string]interface{}, callback ErrorCallback) (string, error) {
	// Vulnerable to XSS if data contains untrusted user inputs
	t, err := template.New("").Parse(tpl)
	if err != nil {
		callback(err)
		return "", err
	}
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		callback(err)
		return "", err
	}
	return buf.String(), nil
}
