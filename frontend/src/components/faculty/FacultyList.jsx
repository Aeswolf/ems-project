import React, { useEffect, useState } from 'react'
import { Button, Spinner } from 'react-bootstrap'
import { Link, useNavigate } from 'react-router-dom'

import axios from 'axios';

function FacultyList() {
    const [faculties, setFaculties] = useState([])
    const [isLoading, setIsLoading] = useState(true)
    const navigate = useNavigate()

    async function handleFacultyDelete(id) {
        const confirm = window.confirm("Do want to delete?")

        if (confirm) {
            const res = await axios.delete(`http://localhost:4000/faculty/${id}`)

            if (res.status === 200) {
                alert(`Faculty with id ${id} is ${res.data.message}`)
                location.reload()
            }
        }
    }

    useEffect(() => {
        async function fetchFaculties() {
            const res = await axios.get("http://localhost:4000/faculty")

            if (res.status === 200) {
                setIsLoading(false)

                res.data ? setFaculties(res.data) : navigate("/faculty/create")
            }

        }


        fetchFaculties();
    }, [])
    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isLoading ? (
                    <div>
                        <Spinner
                            variant="dark"
                            role="status"
                            animation="border"
                            aria-hidden="true"
                        />
                    </div>
                ) : (
                    <>

                        <div className='bg-white w-75 shadow border p-4'>
                            <h1 className='text-center'> List of Faculties </h1>
                            <div className="d-flex justify-content-between">
                                <Link to='/faculty/create' className='btn btn-outline-success my-4'>Add Faculty</Link>
                                <Link to='/' className='btn btn-outline-dark my-4'>Back</Link>
                            </div>
                            <table className='table table-hover text-center'>
                                <thead>
                                    <tr>
                                        <th>No.</th>
                                        <th>Name of Faculty</th>
                                        <th>Actions</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        faculties.map((faculty, index) => (
                                            <tr key={index}>
                                                <td>{index + 1}</td>
                                                <td>{faculty.name}</td>
                                                <td>
                                                    <Link to={`/faculty/${faculty.id}/edit`} className='btn btn-outline-primary me-3'>Edit</Link>
                                                    <button onClick={() => handleFacultyDelete(faculty.id)} className='btn btn-outline-danger me-3'>Delete</button>
                                                    <Link to={`/faculty/${faculty.id}/details`} className='btn btn-outline-dark'>Details</Link>
                                                </td>
                                            </tr>
                                        ))
                                    }
                                </tbody>
                            </table>
                        </div>
                    </>
                )
            }
        </div>
    )
}

export default FacultyList