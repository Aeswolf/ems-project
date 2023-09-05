import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import 'bootstrap/dist/css/bootstrap.min.css';

import Home from './components/Home'
import FacultyList from './components/faculty/FacultyList';
import FacultyCreate from './components/faculty/FacultyCreate';
import FacultyEdit from './components/faculty/FacultyEdit';
import FacultyDetails from './components/faculty/FacultyDetails';
import DeptList from './components/department/DeptList';
import DeptCreate from './components/department/DeptCreate';
import DeptEdit from './components/department/DeptEdit';
import DeptDetails from './components/department/DeptDetails';
import LabList from './components/laboratory/LabList';
import LabCreate from './components/laboratory/LabCreate';
import LabEdit from './components/laboratory/LabEdit';
import LabDetails from './components/laboratory/LabDetails';
import EquipmentList from './components/equipment/EquipmentList';
import EquipmentCreate from './components/equipment/EquipmentCreate';
import EquipmentEdit from './components/equipment/EquipmentEdit';
import EquipmentDetails from './components/equipment/EquipmentDetails';
import EmployeeList from './components/employee/EmployeeList';
import EmployeeCreate from './components/employee/EmployeeCreate';
import EmployeeEdit from './components/employee/EmployeeEdit';
import EmployeeDetails from './components/employee/EmployeeDetails';

function App() {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<Home />} />
        </Routes>
        <Routes>
          <Route path="/faculty" element={<FacultyList />} />
          <Route path="/faculty/create" element={<FacultyCreate />} />
          <Route path="/faculty/:id/edit" element={<FacultyEdit />} />
          <Route path="/faculty/:id/details" element={<FacultyDetails />} />
        </Routes>
        <Routes>
          <Route path="/department" element={<DeptList />} />
          <Route path="/department/create" element={<DeptCreate />} />
          <Route path="/department/:id/edit" element={<DeptEdit />} />
          <Route path="/department/:id/details" element={<DeptDetails />} />
        </Routes>
        <Routes>
          <Route path="/laboratory" element={<LabList />} />
          <Route path="/laboratory/create" element={<LabCreate />} />
          <Route path="/laboratory/:id/edit" element={<LabEdit />} />
          <Route path="/laboratory/:id/details" element={<LabDetails />} />
        </Routes>
        <Routes>
          <Route path="/equipment" element={<EquipmentList />} />
          <Route path="/equipment/create" element={<EquipmentCreate />} />
          <Route path="/equipment/:id/edit" element={<EquipmentEdit />} />
          <Route path="/equipment/:id/details" element={<EquipmentDetails />} />
        </Routes>
        <Routes>
          <Route path="/employee" element={<EmployeeList />} />
          <Route path="/employee/create" element={<EmployeeCreate />} />
          <Route path="/employee/:id/edit" element={<EmployeeEdit />} />
          <Route path="/employee/:id/details" element={<EmployeeDetails />} />
        </Routes>
      </BrowserRouter>
    </div>
  )
}

export default App
