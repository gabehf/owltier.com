import { Outlet } from 'react-router-dom'
import Header from '../Components/Header'

export default function Home() {
    return (
        <div className='full-container'>
        <Header />
        <Outlet />
        </div>
    )
}