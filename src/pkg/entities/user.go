package entities

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
	Preference PreferenceType `gorm:"type:preference_type;not null" json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	Email      string         `gorm:"not null" json:"email" validate:"required,email"`
	Password   string         `gorm:"not null" json:"password" validate:"required,min=8,max=32"`
	Name       string         `gorm:"size:60" json:"name" validate:"omitempty,min=2,max=60"`
	WeightUnit WeightUnitType `gorm:"type:weight_unit_type;not null;column:weightunit" json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit HeightUnitType `gorm:"type:height_unit_type;not null;column:heightunit" json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     int            `gorm:"not null" json:"weight" validate:"required,min=10,max=1000"`
	Height     int            `gorm:"not null" json:"height" validate:"required,min=3,max=250"`
	ImageURI   string         `gorm:"size:255;column:imageuri" json:"imageUri" validate:"omitempty,url"`
}

type UpdateProfile struct {
	Preference PreferenceType `gorm:"type:preference_type;not null" json:"preference" validate:"required,oneof=CARDIO WEIGHT"`
	WeightUnit WeightUnitType `gorm:"type:weight_unit_type;not null;column:weightunit" json:"weightUnit" validate:"required,oneof=KG LBS"`
	HeightUnit HeightUnitType `gorm:"type:height_unit_type;not null;column:heightunit" json:"heightUnit" validate:"required,oneof=CM INCH"`
	Weight     int            `gorm:"not null" json:"weight" validate:"required,min=10,max=1000"`
	Height     int            `gorm:"not null" json:"height" validate:"required,min=3,max=250"`
	Name       string         `gorm:"size:60" json:"name" validate:"omitempty,min=2,max=60" example:""`
	ImageURI   string         `gorm:"size:255;column:imageuri" json:"imageUri" validate:"omitempty,url" example:""`
}
