package entities

// Types definitions
type ActionDenied struct{}
type DuplicatedData struct{}
type NonExistentData struct{}
type NotFoundDataError struct{}
type EmailOrPasswordError struct{}
type EmailAlreadyExistsError struct{}
type InvalidOrExpiredTokenError struct{}

/// Types definitions extensions
// Action Denied error
func (m *ActionDenied) Error() string {
	return "Action denied for false attribution."
}
// Duplicated data error
func (m *DuplicatedData) Error() string {
	return "Possible duplicate data injection."
}
// Non existent data error
func (m *NonExistentData) Error() string {
	return "The requested source data not found."
}
// Not found data error
func (m *NotFoundDataError) Error() string {
	return "Requested data not found..."
}
// Email or password error
func (m *EmailOrPasswordError) Error() string {
	return "Incorrect Email or Password..."
}
// Email already exists error
func (m *EmailAlreadyExistsError) Error() string {
	return "Email already exists on database..."
}
// Invalid or expired token error
func (m *InvalidOrExpiredTokenError) Error() string {
	return "Invalid or expired JWT"
}

