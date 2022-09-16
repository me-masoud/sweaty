package Migration

import (
	"gorm.io/gorm"
)
type Sweater struct {
	gorm.Model
	Email  string
	Name string
}
