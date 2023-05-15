import './Css/Menus.css'
import Menu from './Menu'

const Items = [
    {link: '/build/split-region', text: 'split region'},
    {link: '/build/combined', text: 'combined'},
    {link: '/build/na', text: 'na only'},
    {link: '/build/apac', text: 'apac only'},
]

export default function Options(props: {current: string}) {
    return (
        <Menu 
            title='Options' 
            items={Items} 
            current={props.current} 
            align='left'
            internal
        />
    )
}