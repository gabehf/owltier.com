import { Outlet, useLocation, useNavigate } from 'react-router-dom'
import Header from '../Components/Header'
import Footer from '../Components/Footer'
import { useEffect } from 'react'

export default function Home() {
    const navigate = useNavigate()
    const location = useLocation()
    useEffect(() => {
        if (location.pathname == '/') {
            navigate('/build/split-region')
        }
    })
    return (
        <>
        <div className='full-container'>
            <Header />
            <Outlet /> 
        </div>
        <Footer />
        </>
    )
}