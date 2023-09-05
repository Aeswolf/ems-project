package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	d "github.com/Aeswolf/equipment-database-management/database"
	s "github.com/Aeswolf/equipment-database-management/structs"
	"github.com/Aeswolf/equipment-database-management/utils"
)

// struct for the api server
type APIServer struct {
	Port  string
	Store *d.PostgresStore
}

// constructor for the new api server
func NewServer(p string, s *d.PostgresStore) *APIServer {
	return &APIServer{
		Port:  p,
		Store: s,
	}
}

// handler for creating a faculty
func (a *APIServer) HandleCreateFaculty(w http.ResponseWriter, r *http.Request) error {
	facultyReq := new(s.FacultyRequest)

	if err := json.NewDecoder(r.Body).Decode(facultyReq); err != nil {
		return err
	}

	if facultyReq.Name == "" {
		return fmt.Errorf("name field is empty")
	}

	f := s.NewFaculty(s.WithFacultyCreationTime(), s.WithFacultyName(facultyReq.Name), s.WithFacultyUpdatedTime(), s.WithFacultyID(), s.WithFacultyDeletionTime(time.Time{}))
	id, err := a.Store.AddFaculty(f)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "created")

	return utils.WriteJSON(w, http.StatusCreated, res)
}

// handler for getting all the faculties in the database
func (a *APIServer) HandleGetFaculties(w http.ResponseWriter, r *http.Request) error {
	faculties, err := a.Store.GetFaculties()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, faculties)
}

// handler for getting a single faculty
func (a *APIServer) HandleGetFaculty(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	f, err := a.Store.GetFaculty(reqParams["id"])

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, f)
}

// handler for deleting a faculty
func (a *APIServer) HandleDeleteFaculty(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	err := a.Store.RemoveFaculty(id)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "deleted")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for updating a single faculty
func (a *APIServer) HandleUpdateFaculty(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	facultyReq := new(s.FacultyRequest)

	if err := json.NewDecoder(r.Body).Decode(facultyReq); err != nil {
		return err
	}

	if facultyReq.Name == "" {
		return fmt.Errorf("no field value")
	}

	f := s.NewFaculty(s.WithFacultyName(facultyReq.Name))

	if err := a.Store.UpdateFaculty(id, f); err != nil {
		return err
	}

	res := s.NewResponse(id, "updated")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for creating a department
func (a *APIServer) HandleCreateDepartment(w http.ResponseWriter, r *http.Request) error {
	departmentReq := new(s.DepartmentRequest)

	if err := json.NewDecoder(r.Body).Decode(departmentReq); err != nil {
		return err
	}

	if departmentReq.Name == "" || departmentReq.AssociatedFaculty == "" {
		return fmt.Errorf("name field is empty")
	}

	d := s.NewDepartment(
		s.WithAssociatedFacultyName(departmentReq.AssociatedFaculty),
		s.WithDepartmentCreationTime(),
		s.WithDepartmentID(),
		s.WithDepartmentUpdatedTime(),
		s.WithDepartmentDeletionTime(time.Time{}),
		s.WithDepartmentName(departmentReq.Name),
	)

	id, err := a.Store.AddDepartment(d)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "created")
	return utils.WriteJSON(w, http.StatusCreated, res)
}

// handler for getting all the department in the database
func (a *APIServer) HandleGetDepartments(w http.ResponseWriter, r *http.Request) error {
	departments, err := a.Store.GetDepartments()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, departments)
}

// handler for getting a single faculty
func (a *APIServer) HandleGetDepartment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	d, err := a.Store.GetDepartment(reqParams["id"])

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, d)
}

// handler for the deletion of a department
func (a *APIServer) HandleDeleteDepartment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	if err := a.Store.RemoveDepartment(id); err != nil {
		return err
	}

	res := s.NewResponse(id, "deleted")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for updating a department
func (a *APIServer) HandleUpdateDepartment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	deptReq := new(s.DepartmentRequest)

	if err := json.NewDecoder(r.Body).Decode(deptReq); err != nil {
		return err
	}

	if deptReq.AssociatedFaculty == "" && deptReq.Name == "" {
		return fmt.Errorf("empty fields")
	}

	d := s.NewDepartment(
		s.WithAssociatedFacultyName(deptReq.AssociatedFaculty),
		s.WithDepartmentName(deptReq.Name),
	)

	if err := a.Store.UpdateDepartment(id, d); err != nil {
		return err
	}

	res := s.NewResponse(id, "updated")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for creating a department
func (a *APIServer) HandleCreateLaboratory(w http.ResponseWriter, r *http.Request) error {
	labReq := new(s.LabRequest)

	if err := json.NewDecoder(r.Body).Decode(labReq); err != nil {
		return err
	}

	if labReq.Name == "" || labReq.AssociatedDepartment == "" {
		return fmt.Errorf("name field is empty")
	}

	l := s.NewLab(
		s.WithAssociatedDeptName(labReq.AssociatedDepartment),
		s.WithLabCreationTime(),
		s.WithLabID(),
		s.WithLabUpdatedTime(),
		s.WithLabDeletionTime(time.Time{}),
		s.WithLabName(labReq.Name),
	)

	id, err := a.Store.AddLab(l)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "created")
	return utils.WriteJSON(w, http.StatusCreated, res)
}

// handler for getting a laboratory
func (a *APIServer) HandleGetLaboratory(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	l, err := a.Store.GetLab(reqParams["id"])

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, l)
}

// handler for getting all labs
func (a *APIServer) HandleGetLaboratories(w http.ResponseWriter, r *http.Request) error {
	labs, err := a.Store.GetLabs()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, labs)
}

// handler for deleting a laboratory
func (a *APIServer) HandleDeleteLaboratory(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	if err := a.Store.RemoveLaboratory(id); err != nil {
		return err
	}

	res := s.NewResponse(id, "deleted")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for updating a laboratory
func (a *APIServer) HandleUpdateLaboratory(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	labReq := new(s.LabRequest)

	if err := json.NewDecoder(r.Body).Decode(labReq); err != nil {
		return err
	}

	if labReq.Name == "" || labReq.AssociatedDepartment == "" {
		return fmt.Errorf("empty field")
	}

	l := s.NewLab(
		s.WithLabName(labReq.Name),
		s.WithAssociatedDeptName(labReq.AssociatedDepartment),
	)

	if err := a.Store.UpdateLab(id, l); err != nil {
		return err
	}

	res := s.NewResponse(id, "updated")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for creating a department
func (a *APIServer) HandleCreateEquipment(w http.ResponseWriter, r *http.Request) error {
	equipmentReq := new(s.EquipmentRequest)

	if err := json.NewDecoder(r.Body).Decode(equipmentReq); err != nil {
		return err
	}

	if equipmentReq.Name == "" || equipmentReq.AssociatedLaboratory == "" {
		return fmt.Errorf("name field is empty")
	}

	e := s.NewEquipment(
		s.WithAssociatedLaboratoryName(equipmentReq.AssociatedLaboratory),
		s.WithEquipmentCreationTime(),
		s.WithEquipmentID(),
		s.WithEquipmentUpdatedTime(),
		s.WithEquipmentDeletionTime(time.Time{}),
		s.WithEquipmentName(equipmentReq.Name),
		s.WithQuantity(equipmentReq.Quantity),
	)

	id, err := a.Store.AddEquipment(e)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "created")
	return utils.WriteJSON(w, http.StatusCreated, res)
}

// handler for getting all equipments in the database
func (a *APIServer) HandleGetEquipments(w http.ResponseWriter, r *http.Request) error {
	equipments, err := a.Store.GetEquipments()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, equipments)
}

// handler for getting a single equipment from the database
func (a *APIServer) HandleGetEquipment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	e, err := a.Store.GetEquipment(reqParams["id"])

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, e)
}

// handler for updating an equipment in the database
func (a *APIServer) HandleUpdateEquipment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	equipmentReq := new(s.EquipmentRequest)

	if err := json.NewDecoder(r.Body).Decode(equipmentReq); err != nil {
		return err
	}

	if equipmentReq.Name == "" || equipmentReq.AssociatedLaboratory == "" {
		return fmt.Errorf("empty fields")
	}

	e := s.NewEquipment(
		s.WithAssociatedLaboratoryName(equipmentReq.AssociatedLaboratory),
		s.WithEquipmentName(equipmentReq.Name),
		s.WithQuantity(equipmentReq.Quantity),
	)

	if err := a.Store.UpdateEquipment(id, e); err != nil {
		return err
	}

	res := s.NewResponse(id, "updated")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for deleting an equipment from the database
func (a *APIServer) HandleDeleteEquipment(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	if err := a.Store.RemoveEquipment(id); err != nil {
		return err
	}

	res := s.NewResponse(id, "deleted")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for adding a new employee
func (a *APIServer) HandleCreateEmployee(w http.ResponseWriter, r *http.Request) error {
	empReq := new(s.EmployeeRequest)

	if err := json.NewDecoder(r.Body).Decode(empReq); err != nil {
		return err
	}

	if empReq.FirstName == "" || empReq.LastName == "" || empReq.Email == "" ||
		empReq.AssociatedLaboratory == "" {
		return fmt.Errorf("empty field")
	}

	emp := s.NewEmployee(
		s.WithEmployeeAssociatedLaboratoryName(empReq.AssociatedLaboratory),
		s.WithEmployeeCreationTime(),
		s.WithEmployeeID(),
		s.WithEmployeeUpdatedTime(),
		s.WithEmployeeDeletionTime(time.Time{}),
		s.WithEmployeeLastName(empReq.LastName),
		s.WithEmployeeFirstName(empReq.FirstName),
		s.WithEmployeeEmail(empReq.Email),
	)

	id, err := a.Store.AddEmployee(emp)

	if err != nil {
		return err
	}

	res := s.NewResponse(id, "created")
	return utils.WriteJSON(w, http.StatusCreated, res)
}

// handler for getting all the employees
func (a *APIServer) HandleGetEmployees(w http.ResponseWriter, r *http.Request) error {
	employees, err := a.Store.GetEmployees()

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, employees)
}

// handler for getting an employees
func (a *APIServer) HandleGetEmployee(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	emp, err := a.Store.GetEmployee(reqParams["id"])

	if err != nil {
		return err
	}

	return utils.WriteJSON(w, http.StatusOK, emp)
}

// handler for updating an employees
func (a *APIServer) HandleUpdateEmployee(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)
	empReq := new(s.EmployeeRequest)
	id := reqParams["id"]

	if err := json.NewDecoder(r.Body).Decode(empReq); err != nil {
		return err
	}

	if empReq.FirstName == "" || empReq.LastName == "" || empReq.Email == "" ||
		empReq.AssociatedLaboratory == "" {
		return fmt.Errorf("empty field")
	}

	emp := s.NewEmployee(
		s.WithEmployeeAssociatedLaboratoryName(empReq.AssociatedLaboratory),
		s.WithEmployeeEmail(empReq.Email),
		s.WithEmployeeFirstName(empReq.FirstName),
		s.WithEmployeeLastName(empReq.LastName),
	)

	if err := a.Store.UpdateEmployee(id, emp); err != nil {
		return err
	}

	res := s.NewResponse(id, "updated")

	return utils.WriteJSON(w, http.StatusOK, res)
}

// handler for deleting an employee
func (a *APIServer) HandleDeleteEmployee(w http.ResponseWriter, r *http.Request) error {
	reqParams := mux.Vars(r)

	id := reqParams["id"]

	if err := a.Store.RemoveEmployee(id); err != nil {
		return err
	}

	res := s.NewResponse(id, "deleted")

	return utils.WriteJSON(w, http.StatusOK, res)
}
