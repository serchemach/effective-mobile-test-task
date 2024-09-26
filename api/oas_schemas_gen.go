// Code generated by ogen, DO NOT EDIT.

package api

// InfoGetBadRequest is response for InfoGet operation.
type InfoGetBadRequest struct{}

func (*InfoGetBadRequest) infoGetRes() {}

// InfoGetInternalServerError is response for InfoGet operation.
type InfoGetInternalServerError struct{}

func (*InfoGetInternalServerError) infoGetRes() {}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/SongDetail
type SongDetail struct {
	ReleaseDate string    `json:"releaseDate"`
	Text        string    `json:"text"`
	Patronymic  OptString `json:"patronymic"`
}

// GetReleaseDate returns the value of ReleaseDate.
func (s *SongDetail) GetReleaseDate() string {
	return s.ReleaseDate
}

// GetText returns the value of Text.
func (s *SongDetail) GetText() string {
	return s.Text
}

// GetPatronymic returns the value of Patronymic.
func (s *SongDetail) GetPatronymic() OptString {
	return s.Patronymic
}

// SetReleaseDate sets the value of ReleaseDate.
func (s *SongDetail) SetReleaseDate(val string) {
	s.ReleaseDate = val
}

// SetText sets the value of Text.
func (s *SongDetail) SetText(val string) {
	s.Text = val
}

// SetPatronymic sets the value of Patronymic.
func (s *SongDetail) SetPatronymic(val OptString) {
	s.Patronymic = val
}

func (*SongDetail) infoGetRes() {}
