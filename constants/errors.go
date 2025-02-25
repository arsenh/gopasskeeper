package constants

// error messages
const (
	ErrInternalReadline           = "internal readLine error"
	ErrFatalReadingMasterPassword = "fatal error when reading master password"
	ErrPasswordMismatch           = "❌ Error: Passwords do not match. Please try again."
	ErrCantCreateFileInHomeDir    = "unable to create password file in home directory"
	ErrUnableDeleteHistoryFile    = "unable to delete history file %s, please do it manually for security reasons"
	ErrPasswordFileIsEmpty        = "password file is empty"
	ErrInvalidJsonFormat          = "invalid json format"
	ErrMasterHashIsEmpty          = "master hash is empty"
	ErrGetHomaDir                 = "unable to get home directory"
	ErrCreateFileInHomeDir        = "unable to create file in home directory"
	ErrOpenFile                   = "unable to open existing file in path %s\n"
	ErrInvalidCommand             = "invalid input. use the 'help' command to view detailed instructions and additional information"
	ErrAccessOptionalParam        = "internal error: invalid access to optional parameter"
	ErrCiphertextTooShort         = "ciphertext too short"
	ErrInternalCrypto             = "internal crypto error"
	ErrInternalSerialization      = "internal serialization error"
	ErrNoFunctionAssigned         = "No function assigned"
	ErrInvalidOperationAction     = "invalid operation (action) is selected"
)
