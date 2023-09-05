package structs

import (
	"time"

	"github.com/google/uuid"
)

type labProps struct {
	ID                   uuid.UUID `json:"id"`
	Name                 string    `json:"name"`
	IsDeleted            bool      `json:"is_deleted"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	DeletedAt            time.Time `json:"deleted_at"`
	AssociatedDepartment string    `json:"associated_department"`
}

type labOptions func(*Lab)

// struct for laboratory
type Lab struct {
	labProps
}

func NewLab(opts ...labOptions) *Lab {

	d := &Lab{}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithLabCreationTime() labOptions {
	return func(l *Lab) {
		l.CreatedAt = time.Now()
	}
}

func WithLabDeletionTime(t time.Time) labOptions {
	return func(l *Lab) {
		l.DeletedAt = t
	}
}

func WithLabIsDeleted() labOptions {
	return func(l *Lab) {
		l.IsDeleted = true
	}
}

func WithLabName(n string) labOptions {
	return func(l *Lab) {
		if l.Name != n {
			l.Name = n
		}
	}
}

func WithLabUpdatedTime() labOptions {
	return func(l *Lab) {
		l.UpdatedAt = time.Now()
	}
}

func WithLabID() labOptions {
	return func(l *Lab) {
		l.ID = uuid.New()
	}
}

func WithAssociatedDeptName(n string) labOptions {
	return func(l *Lab) {
		if l.AssociatedDepartment != n {
			l.AssociatedDepartment = n
		}
	}
}

func (l *Lab) UpdateLab(opts ...labOptions) {
	for _, opt := range opts {
		opt(l)
	}
}
