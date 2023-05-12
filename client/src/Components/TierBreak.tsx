import { useState } from 'react';
import './Css/TierBreak.css'

export default function TierBreak() {
    const [isActive, setActive] = useState(false);
    return (
        <div className='tier-break'>
            <div 
                className={'triangle-right' + (isActive ? ' active' : '')}
                onClick={() => setActive(!isActive)}
            >
            </div>
            <hr className={'break-line' + (isActive ? ' active' : '')}/>
        </div>
    )
}