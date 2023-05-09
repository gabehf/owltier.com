import './Css/Forms.css'
import { Link } from 'react-router-dom'

export default function RegisterForm() {
    return (
        <div className="basic-form">
            <label htmlFor="username">username</label>
            <input type="text" name="username" id="username" />
            <label htmlFor="password">password</label>
            <input type="password" name="password" id="password" />
            <label htmlFor="confpassword">confirm password</label>
            <input type="password" name="confpassword" id="confpassword" />
            <button type="submit">submit</button>
            <Link to='/accounts/login' className='switch'>already have an account?</Link>
            <div className="form-err"></div>
        </div>
    )
}