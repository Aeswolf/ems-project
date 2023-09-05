package structs

import (
	"time"

	"github.com/google/uuid"
)

type equipmentProps struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	IsDeleted            bool      `json:"is_deleted"`
	Quantity             int       `json:"quantity"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	AssociatedLaboratory string    `json:"associated_laboratory"`
}

type equipmentOptions func(*Equipment)

// struct for faculty
type Equipment struct {
	equipmentProps
}

func defaultEquipmentProps() *Equipment {
	props := equipmentProps{
		IsDeleted: false,
	}

	return &Equipment{
		equipmentProps: props,
	}
}

func NewEquipment(opts ...equipmentOptions) *Equipment {

	d := defaultEquipmentProps()

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithAssociatedLaboratoryName(n string) equipmentOptions {
	return func(e *Equipment) {
		e.AssociatedLaboratory = n
	}
}

func WithEquipmentCreationTime() equipmentOptions {
	return func(e *Equipment) {
		e.CreatedAt = time.Now()
	}
}

func WithEquipmentDeletionTime(t time.Time) equipmentOptions {
	return func(e *Equipment) {
		e.DeletedAt = t
	}
}

func WithEquipmentIsDeleted() equipmentOptions {
	return func(e *Equipment) {
		e.IsDeleted = true
	}
}

func WithEquipmentName(n string) equipmentOptions {
	return func(e *Equipment) {
		if e.Name != n {
			e.Name = n
		}
	}
}

func WithEquipmentUpdatedTime() equipmentOptions {
	return func(e *Equipment) {
		e.UpdatedAt = time.Now()
	}
}

func WithEquipmentID() equipmentOptions {
	return func(e *Equipment) {
		e.ID = uuid.New()
	}
}

func WithQuantity(n int) equipmentOptions {
	return func(e *Equipment) {
		e.Quantity = n
	}
}

func (e *Equipment) UpdateEquipment(opts ...equipmentOptions) {
	for _, opt := range opts {
		opt(e)
	}
}
