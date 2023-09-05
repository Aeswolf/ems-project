package structs

type FacultyRequest struct {
	Name string `json:"name"`
}

type DepartmentRequest struct {
	Name              string `json:"name"`
	AssociatedFaculty string `json:"associated_faculty"`
}

type LabRequest struct {
	Name                 string `json:"name"`
	AssociatedDepartment string `json:"associated_department"`
}

type EquipmentRequest struct {
	Name                 string `json:"name"`
	Quantity             int    `json:"quantity"`
	AssociatedLaboratory string `json:"associated_laboratory"`
}

type response struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

type EmployeeRequest struct {
	FirstName            string `json:"first_name"`
	LastName             string `json:"last_name"`
	Email                string `json:"email"`
	AssociatedLaboratory string `json:"associated_laboratory"`
}

func NewResponse(id, msg string) response {
	return response{
		ID:      id,
		Message: msg + " successfully",
	}
}
