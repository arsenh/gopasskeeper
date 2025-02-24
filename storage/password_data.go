package storage

type Password struct {
	Service  string
	UserName string
	Password string
	Note     string
}

type Passwords []Password

type PasswordJson struct {
	MasterKeyHash string `json:"master_key_hash"`
	Data          string `json:"data"`
}
