import React, { useEffect, useState } from 'react'
import { useParams, useNavigate, Link } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'
import axios from 'axios'

function FacultyEdit() {
    const [oldValue, setOldValue] = useState({})
    const [facultyInfo, setFacultyInfo] = useState({})
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const { id } = useParams()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/faculty/${id}`)

            if (res.data) {
                setFacultyInfo(res.data)
                setOldValue(res.data)
                setIsInfoLoaded(true)
            }
        }

        fetchData();
    }, [])

    const navigate = useNavigate()

    const handleInputChange = e => {
        setFacultyInfo({ ...facultyInfo, name: e.target.value })
    }

    const handleFormSubmit = async (e) => {
        e.preventDefault();

        if (facultyInfo.name !== "" && facultyInfo.name !== oldValue.name) {
            const res = await axios.put(`http://localhost:4000/faculty/${id}`, facultyInfo)

            alert(`Faculty with an id of ${res.data["id"]} has been updated`)
        }
        navigate("/faculty")
    }
    return (
        <div className='d-flex w-100 vh-100 justify-content-center align-items-center flex-column bg-light'>
            {
                isInfoLoaded ? (
                    <div className='w-50 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                        <h2 className='text-center my-3'>Edit Faculty</h2>
                        <form method='POST' onSubmit={handleFormSubmit} >
                            <div className='my-2 mx-4 row'>
                                <label htmlFor='name' className="fw-bold fs-4 mx-3">Name</label>
                                <div>
                                    <input required className='form-control form-control-lg' value={facultyInfo.name} placeholder='Enter faculty name' name='name' type='text' id='name' onChange={handleInputChange} />
                                </div>
                            </div>
                            <div className='col-auto d-flex justify-content-center align-items-center my-4 flex-column'>
                                <button type='submit' className='btn btn-outline-success mb-3 w-25'>Submit</button>
                                <Link to="/faculty" className='btn btn-outline-dark w-25'>Back</Link>
                            </div>
                        </form>
                    </div>
                ) : (
                    <div>
                        <Spinner
                            variant="dark"
                            role="status"
                            animation="border"
                            aria-hidden="true"
                        />
                    </div>
                )
            }
        </div>
    )
}

export default FacultyEdit