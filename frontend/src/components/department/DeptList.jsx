import React, { useEffect, useState } from 'react'
import axios from 'axios';
import { Link, useNavigate } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'

function DeptList() {
    const [depts, setDepts] = useState([])
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get(`http://localhost:4000/department`)
            const resData = res.data

            if (resData || resData === null) {
                setIsInfoLoaded(true)
            }

            if (resData === null) {
                navigate("/department/create")
                return
            }

            setDepts(resData)
        }

        fetchData()
    }, [])

    const handleDeleteDept = async (id) => {
        const confirm = window.confirm("Are sure you want to delete?")

        if (confirm) {
            const res = await axios.delete(`http://localhost:4000/department/${id}`)
            alert(`Department with id ${id} has been ${res.data.message}`)
            location.reload()
        }
    }
    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isInfoLoaded ? (
                    <>
                        <div className='bg-white w-75 shadow border p-4'>
                            <h1 className='text-center'> List of Departments </h1>
                            <div className="d-flex justify-content-between">
                                <Link to='/department/create' className='btn btn-outline-success my-4'>Add Department</Link>
                                <Link to='/' className='btn btn-outline-dark my-4'>Back</Link>
                            </div>
                            <table className='table table-hover text-center'>
                                <thead>
                                    <tr>
                                        <th>No.</th>
                                        <th>Name of Department</th>
                                        <th>Associated Faculty</th>
                                        <th>Actions</th>
                                    </tr>
                                </thead>
                                <tbody>
                                    {
                                        depts.map((d, i) => (
                                            <tr key={i}>
                                                <td>{i + 1}</td>
                                                <td>{d.name}</td>
                                                <td>{d.associated_faculty}</td>
                                                <td>
                                                    <Link to={`/department/${d.id}/edit`} className="btn btn-outline-primary me-3">Edit</Link>
                                                    <button onClick={() => handleDeleteDept(d.id)} className="btn btn-outline-danger me-4">Delete</button>
                                                    <Link to={`/department/${d.id}/details`} className="btn btn-outline-dark">Details</Link>
                                                </td>
                                            </tr>
                                        ))
                                    }
                                </tbody>
                            </table>
                        </div>
                    </>
                ) : (
                    <Spinner
                        variant="dark"
                        animation="border"
                        aria-hidden="true"
                    />
                )
            }
        </div>
    )
}

export default DeptList