import React from 'react'
import { Link } from 'react-router-dom'

function Home() {
    return (
        <div className='d-flex justify-content-center align-items-center flex-column vh-100'>
            <div className='w-25 border bg-white shadow px-5 pt-3 pb-5 rounded d-flex justify-content-center flex-column'>
                <Link className='btn btn-outline-primary my-3 ' to="/faculty">faculty</Link>
                <Link className='btn btn-outline-primary mb-3' to="/department">department</Link>
                <Link className='btn btn-outline-primary mb-3' to="/laboratory">laboratory</Link>
                <Link className='btn btn-outline-primary mb-3' to="/equipment">equipment</Link>
                <Link className='btn btn-outline-primary' to="/employee">employee</Link>
            </div>
        </div>
    )
}

export default Home