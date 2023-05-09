import { Outlet } from 'react-router-dom'
import Options from '../Components/Options'
import Nav from '../Components/Nav'
import './Css/Main.css'

export default function Build() {
    return (
        <div className="page-content">
            <div className="left-menus">
                <Options current='split-region' />
            </div>
            <Outlet />
            <div className="right-menus">
                <Nav current='build' />
            </div>
        </div>
    )
}