import { Outlet } from 'react-router-dom'
import Header from '../Components/Header'
import Footer from '../Components/Footer'

export default function Home() {
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