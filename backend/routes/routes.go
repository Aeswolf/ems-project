package routes

import (
	"github.com/gorilla/mux"

	"github.com/Aeswolf/equipment-database-management/api"
	"github.com/Aeswolf/equipment-database-management/utils"
)

func Access(a *api.APIServer, r *mux.Router) {
	// for faculty
	r.HandleFunc("/faculty", utils.MakeHTTPHandler(a.HandleCreateFaculty)).Methods("POST")
	r.HandleFunc("/faculty", utils.MakeHTTPHandler(a.HandleGetFaculties)).Methods("GET")
	r.HandleFunc("/faculty/{id}", utils.MakeHTTPHandler(a.HandleGetFaculty)).Methods("GET")
	r.HandleFunc("/faculty/{id}", utils.MakeHTTPHandler(a.HandleUpdateFaculty)).Methods("PUT")
	r.HandleFunc("/faculty/{id}", utils.MakeHTTPHandler(a.HandleDeleteFaculty)).Methods("DELETE")

	// for department
	r.HandleFunc("/department", utils.MakeHTTPHandler(a.HandleCreateDepartment)).Methods("POST")
	r.HandleFunc("/department", utils.MakeHTTPHandler(a.HandleGetDepartments)).Methods("GET")
	r.HandleFunc("/department/{id}", utils.MakeHTTPHandler(a.HandleGetDepartment)).Methods("GET")
	r.HandleFunc("/department/{id}", utils.MakeHTTPHandler(a.HandleDeleteDepartment)).Methods("DELETE")
	r.HandleFunc("/department/{id}", utils.MakeHTTPHandler(a.HandleUpdateDepartment)).Methods("PUT")

	// for laboratory
	r.HandleFunc("/laboratory", utils.MakeHTTPHandler(a.HandleCreateLaboratory)).Methods("POST")
	r.HandleFunc("/laboratory", utils.MakeHTTPHandler(a.HandleGetLaboratories)).Methods("GET")
	r.HandleFunc("/laboratory/{id}", utils.MakeHTTPHandler(a.HandleGetLaboratory)).Methods("GET")
	r.HandleFunc("/laboratory/{id}", utils.MakeHTTPHandler(a.HandleDeleteLaboratory)).Methods("DELETE")
	r.HandleFunc("/laboratory/{id}", utils.MakeHTTPHandler(a.HandleUpdateLaboratory)).Methods("PUT")

	// for equipment
	r.HandleFunc("/equipment", utils.MakeHTTPHandler(a.HandleCreateEquipment)).Methods("POST")
	r.HandleFunc("/equipment", utils.MakeHTTPHandler(a.HandleGetEquipments)).Methods("GET")
	r.HandleFunc("/equipment/{id}", utils.MakeHTTPHandler(a.HandleGetEquipment)).Methods("GET")
	r.HandleFunc("/equipment/{id}", utils.MakeHTTPHandler(a.HandleDeleteEquipment)).Methods("DELETE")
	r.HandleFunc("/equipment/{id}", utils.MakeHTTPHandler(a.HandleUpdateEquipment)).Methods("PUT")

	// for employee
	r.HandleFunc("/employee", utils.MakeHTTPHandler(a.HandleCreateEmployee)).Methods("POST")
	r.HandleFunc("/employee", utils.MakeHTTPHandler(a.HandleGetEmployees)).Methods("GET")
	r.HandleFunc("/employee/{id}", utils.MakeHTTPHandler(a.HandleGetEmployee)).Methods("GET")
	r.HandleFunc("/employee/{id}", utils.MakeHTTPHandler(a.HandleDeleteEmployee)).Methods("DELETE")
	r.HandleFunc("/employee/{id}", utils.MakeHTTPHandler(a.HandleUpdateEmployee)).Methods("PUT")
}
