package errs

type Error struct {
	Message string            `json:"message"`
	Code    int8              `json:"code"`
	Details map[string]string `json:"details"`
	Err     error             `json:"-"`
}
