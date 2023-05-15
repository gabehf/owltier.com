import { Outlet } from 'react-router-dom'
import Options from '../Components/Options'
import Nav from '../Components/Nav'
import './Css/Main.css'
import Share from '../Components/Share'
import React from 'react'
import ListDefaults from './Build/ListDefaults'

// TODO: List state only updates when a team is moved, SO
// we should define the initial state of the list here, and
// the build components will use that list as initial state

// TODO: Need to track tier splits

// TODO: Prop drilling list context all the way into TierBreak
// components feels super gross, should figure out a better way

export type ListState = { 
    format: string, 
    na: Array<string>, 
    apac: Array<string>, 
    combined: Array<string>,
    breaks: Array<number> 
}
export type ListContext = [list: ListState, setter: React.Dispatch<React.SetStateAction<ListState>>]

let initialList = {
    format: 'split',
    na: ListDefaults.na,
    apac: ListDefaults.apac,
    combined: ListDefaults.apac.concat(ListDefaults.na),
    // 20 teams + first blank spot
    breaks:  new Array(21).fill(0)
}

export default function Build() {
    const listContext = React.useState<ListState>(initialList)
    return (
        <div className="page-content">
            <div className="left-menus">
                <Options current='split-region' />
                <Share listContext={listContext} />
            </div>
            <Outlet context={listContext}/>
            <div className="right-menus">
                <Nav current='build' />
            </div>
        </div>
    )
}