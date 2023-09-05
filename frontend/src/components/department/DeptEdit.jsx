import React, { useEffect, useState } from 'react'
import { Link, useParams, useNavigate } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'
import axios from 'axios'

function DeptEdit() {
    const initialDetails = {
        name: "",
        associated_faculty: ""
    }

    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const [deptInfo, setDeptInfo] = useState(initialDetails)
    const [oldValues, setOldValues] = useState(initialDetails)
    const { id } = useParams()
    const navigate = useNavigate()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/department/${id}`)
            setDeptInfo(res.data)
            setOldValues(res.data)
            setIsInfoLoaded(true)
        }

        fetchData()
    }, [])

    const handleFormSubmit = async (e) => {
        e.preventDefault()

        if (deptInfo.name !== oldValues.name || deptInfo.associated_faculty !== oldValues.associated_faculty) {
            const res = await axios.put(`http://localhost:4000/department/${id}`, deptInfo)

            if (res.status === 200) {
                alert(`Department with id ${id} has been ${res.data.message}`)
            }
        }

        navigate("/department")
    }

    const handleInputChange = e => {
        const { name, value } = e.target
        setDeptInfo({ ...deptInfo, [name]: value })
    }

    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-3'>Edit Department</h3>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='mb-3 row'>
                                <label htmlFor='name' className='mx-3 fs-4 fw-bold'>Name</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={deptInfo.name} placeholder='Enter department name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='mb-3 row'>
                                <label htmlFor='associated_faculty' className='mx-3 fs-4 fw-bold'>Associated Faculty</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={deptInfo.associated_faculty} placeholder='Enter faculty name' name='associated_faculty' type='text' id='associated_faculty' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/department" className='btn btn-outline-dark w-25'>Back</Link>
                            </div>
                        </form>
                    </div>
                ) : (
                    <Spinner
                        variant="dark"
                        animation="border"
                        role="status"
                    />
                )
            }
        </div>
    )
}

export default DeptEdit