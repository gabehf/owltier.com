import { Link } from 'react-router-dom'

interface MenuItem {
    link: string
    text: string
}

export default function Menu(props: {title: string, items: Array<MenuItem>, current: string, align: string}) {
    let menuitems = []
    for (let item of props.items) {
        menuitems.push(<Link to={`${item['link']}`} className='nav-link'>{item['text']}</Link>)
    }
    let align = props.align == 'left' ? 'nav-left' : 'nav-right'
    return (
        <div className={`basic-nav ${align}`}>
            <div className="nav-header">{props.title}</div>
            {menuitems}
        </div>
    )
}