import RegisterForm from '../Components/RegisterForm'
import Nav from '../Components/Nav'

export default function Login() {
    return (
        <div className="page-content no-w">
        <RegisterForm />
        <Nav current='login'/>
        </div>
    )
}