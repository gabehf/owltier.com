import { useState } from 'react';
import './Css/TierBreak.css'
import { ListContext } from '../Pages/Build';

export default function TierBreak(props: {listContext: ListContext, id: number}) {
    const [isActive, setActive] = useState(false);
    const [list, setList] = props.listContext
    return (
        <div className='tier-break'>
            <div 
                className={'triangle-right' + (isActive ? ' active' : '')}
                onClick={() => {
                    setActive(!isActive)
                    let nb = list.breaks
                    nb[props.id] = isActive ? 0 : 1
                    setList({
                        format: list.format,
                        na: list.na,
                        apac: list.apac,
                        breaks: nb
                    })
                }}
            >
            </div>
            <hr className={'break-line' + (isActive ? ' active' : '')}/>
        </div>
    )
}