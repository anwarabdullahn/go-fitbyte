package entities

import "github.com/google/uuid"

type PreferenceType string
type WeightUnitType string
type HeightUnitType string

const (
	// PreferenceType
	PreferenceCardio PreferenceType = "CARDIO"
	PreferenceWeight PreferenceType = "WEIGHT"

	// WeightUnitType
	WeightKG  WeightUnitType = "KG"
	WeightLBS WeightUnitType = "LBS"

	// HeightUnitType
	HeightCM   HeightUnitType = "CM"
	HeightINCH HeightUnitType = "INCH"
)

type User struct {
	ID         uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Preference PreferenceType `gorm:"type:preference_type;default:null;" json:"preference" validate:"oneof=CARDIO WEIGHT"`
	Email      string         `gorm:"not null;uniqueIndex" json:"email" validate:"required,email"`
	Password   string         `gorm:"not null" json:"password" validate:"required,min=8,max=32"`
	Name       string         `gorm:"size:60;default:null;column:name;" json:"name" validate:"omitempty,min=2,max=60"`
	WeightUnit WeightUnitType `gorm:"type:weight_unit_type;column:weightunit;default:null" json:"weightUnit" validate:"oneof=KG LBS"`
	HeightUnit HeightUnitType `gorm:"type:height_unit_type;column:heightunit;default:null" json:"heightUnit" validate:"oneof=CM INCH"`
	Weight     int            `gorm:"default:null" json:"weight" validate:"min=10,max=1000"`
	Height     int            `gorm:"default:null" json:"height" validate:"min=3,max=250"`
	ImageURI   string         `gorm:"default:null;size:255;column:imageuri" json:"imageUri" validate:"omitempty,url"`

	// Relasi One-to-One ke UserFile
	UserFile UserFile `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"userFile"`
}
type UserDeleteRequest struct {
	ID uuid.UUID `json:"id"`
}

type AuthRequest struct {
	Email    string `gorm:"not null" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"password" validate:"required,min=8,max=32"`
}

type UpdateProfile struct {
	Preference PreferenceType `gorm:"type:preference_type;not null" json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit WeightUnitType `gorm:"type:weight_unit_type;not null;column:weightunit" json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit HeightUnitType `gorm:"type:height_unit_type;not null;column:heightunit" json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     int            `gorm:"not null" json:"weight" validate:"required,min=10,max=1000"`
	Height     int            `gorm:"not null" json:"height" validate:"required,min=3,max=250"`
	Name       string         `gorm:"size:60;default:null;column:name;" json:"name" validate:"omitempty,max=60" example:""`
	ImageURI   string         `gorm:"size:255;column:imageuri;default:null;" json:"imageUri" validate:"omitempty,url" example:""`
}

type UserFile struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	URI    string `gorm:"size:255;not null" json:"uri" validate:"required"`
	UserID uint   `gorm:"uniqueIndex" json:"userId" validate:"required"` // one-to-one
}

type UserFileResponse struct {
	URI string `json:"uri"`
}

type UserFileRequest struct {
	File string `json:"file" validate:"required,file,ext=jpg,jpeg,png"`
}
