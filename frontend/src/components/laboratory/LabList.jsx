import React, { useEffect, useState } from 'react'
import { Spinner } from 'react-bootstrap'
import { Link, useNavigate } from 'react-router-dom'
import axios from 'axios'

function LabList() {
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const [labs, setLabs] = useState([])
    const navigate = useNavigate()


    useEffect(() => {
        async function fetchData() {
            const res = await axios.get("http://localhost:4000/laboratory")

            if (res.status === 200) {
                setIsInfoLoaded(true)

                res.data ? setLabs(res.data) : navigate("/laboratory/create")
            }
        }

        fetchData()
    }, [])


    const handleLabDelete = async (id) => {
        const confirm = window.confirm("Are you sure you want to delete?")

        if (confirm) {
            const res = await axios.delete(`http://localhost:4000/laboratory/${id}`)

            if (res.data) {
                alert(`Lab with id ${id} has been ${res.data.message}`)
            }
            location.reload()
        }

    }

    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isInfoLoaded ? (
                    <div className='bg-white w-75 shadow border p-4'>
                        <h1 className='text-center'>List of Laboratories</h1>
                        <div className="d-flex justify-content-between">
                            <Link to='/laboratory/create' className='btn btn-outline-success my-4'>Add Lab</Link>
                            <Link to='/' className='btn btn-outline-dark my-4'>Back</Link>
                        </div>
                        <table className='table table-hover text-center'>
                            <thead>
                                <tr>
                                    <th>No.</th>
                                    <th>Name of Lab</th>
                                    <th>Associated Department</th>
                                    <th>Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    labs.map((l, i) => (
                                        <tr key={i}>
                                            <td>{i + 1}</td>
                                            <td>{l.name}</td>
                                            <td>{l.associated_department}</td>
                                            <td>
                                                <Link to={`/laboratory/${l.id}/edit`} className='btn btn-outline-primary me-3'>Edit</Link>
                                                <button onClick={() => handleLabDelete(l.id)} className='btn btn-outline-danger me-3'>Delete</button>
                                                <Link to={`/laboratory/${l.id}/details`} className='btn btn-outline-dark'>Details</Link>
                                            </td>
                                        </tr>
                                    ))
                                }
                            </tbody>
                        </table>
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

export default LabList