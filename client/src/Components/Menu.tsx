import { Link } from 'react-router-dom'

interface MenuItem {
    link?: string
    action?: React.MouseEventHandler<HTMLAnchorElement>
    text: string
}

export default function Menu(props: {
            title: string, 
            items: Array<MenuItem>, 
            current?: string, 
            align: string, 
            internal?: boolean
        }) {
    let menuitems = []
    for (let item of props.items) {
        if (props.internal) {
            menuitems.push(<Link to={`${item['link']}`} className='nav-link'>{item['text']}</Link>)
        } else {
            if (item.action !== undefined) {
                menuitems.push(<a onClick={item.action} className='nav-link pointer'>{item['text']}</a>)
            } else {
                menuitems.push(<a href={`${item['link']}`} className='nav-link'>{item['text']}</a>)
            }
        }
    }
    let align = props.align == 'left' ? 'nav-left' : 'nav-right'
    return (
        <div className={`basic-nav ${align}`}>
            <div className="nav-header">{props.title}</div>
            {menuitems}
        </div>
    )
}