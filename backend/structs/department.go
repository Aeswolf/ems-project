package structs

import (
	"time"

	"github.com/google/uuid"
)

type departmentProps struct {
	ID                uuid.UUID `json:"id"`
	Name              string    `json:"name"`
	IsDeleted         bool      `json:"is_deleted"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	DeletedAt         time.Time `json:"deleted_at"`
	AssociatedFaculty string    `json:"associated_faculty"`
}

type departmentOptions func(*Department)

// struct for Department
type Department struct {
	departmentProps
}

func NewDepartment(opts ...departmentOptions) *Department {

	d := &Department{}

	for _, opt := range opts {
		opt(d)
	}

	return d
}

func WithDepartmentCreationTime() departmentOptions {
	return func(d *Department) {
		d.CreatedAt = time.Now()
	}
}

func WithDepartmentDeletionTime(t time.Time) departmentOptions {
	return func(d *Department) {
		d.DeletedAt = t
	}
}

func WithDepartmentIsDeleted() departmentOptions {
	return func(d *Department) {
		d.IsDeleted = true
	}
}

func WithDepartmentName(n string) departmentOptions {
	return func(d *Department) {
		if d.Name != n {
			d.Name = n
		}
	}
}

func WithDepartmentUpdatedTime() departmentOptions {
	return func(d *Department) {
		d.UpdatedAt = time.Now()
	}
}

func WithDepartmentID() departmentOptions {
	return func(d *Department) {
		d.ID = uuid.New()
	}
}

func WithAssociatedFacultyName(n string) departmentOptions {
	return func(d *Department) {
		if d.AssociatedFaculty != n {
			d.AssociatedFaculty = n
		}
	}
}

func (d *Department) UpdateDepartment(opts ...departmentOptions) {
	for _, opt := range opts {
		opt(d)
	}
}
