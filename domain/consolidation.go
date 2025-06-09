package domain

import (
	"time"

	"gorm.io/gorm"
)

type Consolidation struct {
	gorm.Model
	ID               uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName         string         `gorm:"type:varchar(255);not null" json:"full_name"`
	Phone            string         `gorm:"type:varchar(50)" json:"phone"`
	Address          string         `gorm:"type:varchar(255)" json:"address"`
	Age              int            `json:"age"`
	AttendsCellGroup bool           `gorm:"default:false" json:"attends_cell_group"`
	CallDay          string         `gorm:"type:varchar(150)" json:"call_day"`   // Can be improved to time.Time if specific date format is known
	CallTime         string         `gorm:"type:varchar(150)" json:"call_time"`  // Can be improved to time.Time if specific time format is known
	VisitDay         string         `gorm:"type:varchar(150)" json:"visit_day"`  // Can be improved to time.Time if specific date format is known
	VisitTime        string         `gorm:"type:varchar(150)" json:"visit_time"` // Can be improved to time.Time if specific time format is known
	InvitedBy        string         `gorm:"type:varchar(150)" json:"invited_by"`
	Consolidator     string         `gorm:"type:varchar(150)" json:"consolidator"`
	DocumentType     string         `gorm:"type:varchar(150)" json:"document_type"`
	DocumentNumber   string         `gorm:"type:varchar(150)" json:"document_number"`
	MaritalStatus    string         `gorm:"type:varchar(150)" json:"marital_status"`
	Petition         string         `gorm:"type:text" json:"petition"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
