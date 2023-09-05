import React, { useEffect, useState } from 'react'
import { useNavigate, useParams, Link } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'
import axios from 'axios'

function LabEdit() {
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const [oldValues, setOldValues] = useState({})
    const [labInfo, setLabInfo] = useState({})
    const navigate = useNavigate()
    const { id } = useParams()

    const handleFormSubmit = async (e) => {
        e.preventDefault()

        if (labInfo.name !== oldValues.name || labInfo.associated_dept !== oldValues.associated_dept) {
            const res = await axios.put(`http://localhost:4000/laboratory/${id}`, labInfo)

            if (res.data) {
                alert(`Lab with id ${id} has been ${res.data.message}`)
            }
        }

        navigate("/laboratory")
    }

    const handleInputChange = e => {
        const { name, value } = e.target
        setLabInfo({ ...labInfo, [name]: value })
    }

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/laboratory/${id}`)

            if (res.data) {
                setOldValues(res.data)
                setLabInfo(res.data)
                setIsInfoLoaded(true)
            }
        }

        fetchData()
    }, [])


    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center flex-column bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h3 className='text-center mb-4'>Edit Lab</h3>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='mb-3 row'>
                                <label htmlFor='name' className='mb-1 fs-4 fw-bold mx-3'>Name</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={labInfo.name} placeholder='Enter lab name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='mb-3 row'>
                                <label htmlFor='associated_department' className='mb-1 fs-4 mx-3 fw-bold'>Associated Department</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={labInfo.associated_department} placeholder='Enter department name' name='associated_department' type='text' id='associated_department' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/laboratory" className='btn btn-outline-dark w-25'>Back</Link>
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

export default LabEdit