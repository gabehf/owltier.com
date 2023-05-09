import './Css/Forms.css'
import { Link } from 'react-router-dom'

export default function LoginForm() {
    return (
        <div className="basic-form">
            <label htmlFor="username">username</label>
            <input type="text" name="username" id="username" />
            <label htmlFor="password">password</label>
            <input type="password" name="password" id="password" />
            <button type="submit">submit</button>
            <Link to='/accounts/register' className='switch'>dont have an account?</Link>
            <div className="form-err"></div>
        </div>
    )
}