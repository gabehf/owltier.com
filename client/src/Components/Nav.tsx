import Menu from './Menu'
import './Css/Menus.css'
const Items = [
    {link: '/accounts/login', text: 'sign in / sign up'},
    {link: '/build/split-region', text: 'build tier list'},
    {link: '/community-rankings', text: 'community rankings'},
]

export default function Options(props: {current: string}) {
    return (
        <Menu 
            title='Navigation' 
            items={Items} 
            current={props.current} 
            align='right'
        />
    )
}