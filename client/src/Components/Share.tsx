import Menu from './Menu'
import './Css/Menus.css'
import type { ListContext } from '../Pages/Build'

function exportImage() {
    alert('Image exported!')
}

export default function Share(props: {current?: string, listContext: ListContext}) {
    const [list] = props.listContext
    const Items = [
        {action: save, text: 'save tier list'},
        {action: publish, text: 'publish tier list'},
        {link: 'https://twitter.com', text: 'share on twitter'},
        {action: exportImage, text: 'export as image'},
    ]
    function save() {
        alert('saved!')
    }
    function publish() {
        console.log(list)
    }
    return (
        <Menu 
            title='Share' 
            items={Items} 
            current={props.current} 
            align='left'
        />
    )
}