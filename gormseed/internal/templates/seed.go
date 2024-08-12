package templates

func SeedTemplate() []byte {
	return []byte(`// Code generated by Gormseed (gorms). DO NOT EDIT.

package seeds

import (
	"path/filepath"
	"runtime"
)

type Seeds struct{}

func (*Seeds) Path() string {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	return basepath
}
`)
}

func SeederTemplate() []byte {
	return []byte(`// Code generated by Gormseed (gorms).

package seeds

import (
	"gorm.io/gorm"
)

// don't rename this function
func (s *Seeds) {{ .SeedFuncName }}(db *gorm.DB) error {
	// place your seed code here
	return nil
}

// don't rename this function
func (s *Seeds) {{ .RollbackFuncName }}(db *gorm.DB) error {
	// place your rollback code here
	return nil
}	
`)
}
