import React, { useEffect, useState } from 'react'
import { Link, useNavigate } from 'react-router-dom'
import { Spinner } from 'react-bootstrap'
import axios from 'axios'


function EquipmentList() {
    const [equipmentList, setEquipmentList] = useState([])
    const [isInfoLoaded, setIsInfoLoaded] = useState(false)
    const navigate = useNavigate()

    useEffect(() => {
        async function fetchData() {
            const res = await axios.get("http://localhost:4000/equipment")

            if (res.status === 200) {
                setIsInfoLoaded(true)

                res.data ? setEquipmentList(res.data) : navigate("/equipment/create")

            }

        }

        fetchData()
    }, [])

    const handleDeleteEquipment = async (id) => {

    }

    return (
        <div className='d-flex flex-column justify-content-center align-items-center vh-100 bg-light'>
            {
                isInfoLoaded ? (
                    <div className='bg-white w-75 shadow border p-4'>
                        <h1 className='text-center'> List of Equipments </h1>
                        <div className="d-flex justify-content-between">
                            <Link to='/equipment/create' className='btn btn-outline-success my-4'>Add Equipment</Link>
                            <Link to='/' className='btn btn-outline-dark my-4'>Back</Link>
                        </div>
                        <table className='table table-hover text-center'>
                            <thead>
                                <tr>
                                    <th>No.</th>
                                    <th>Name</th>
                                    <th>Quantity</th>
                                    <th>Associated Lab</th>
                                    <th>Actions</th>
                                </tr>
                            </thead>
                            <tbody>
                                {
                                    equipmentList.map((eq, i) => (
                                        <tr key={i}>
                                            <td>{i + 1}</td>
                                            <td>{eq.name}</td>
                                            <td>{eq.quantity}</td>
                                            <td>{eq.associated_laboratory}</td>
                                            <td>
                                                <Link to={`/equipment/${eq.id}/edit`} className="btn btn-outline-primary me-3">Edit</Link>
                                                <button onClick={() => handleDeleteEquipment(eq.id)} className="btn btn-outline-danger me-4">Delete</button>
                                                <Link to={`/equipment/${eq.id}/details`} className="btn btn-outline-dark">Details</Link>
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

export default EquipmentList