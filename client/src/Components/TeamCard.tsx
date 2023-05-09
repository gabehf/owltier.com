import './Css/TeamCard.css'
import './Css/TeamStyles.css'

export default function TeamCard(props: {team: string}) {
    let teamCss = props.team.replaceAll(' ', '').toLowerCase()
    return (
        <div className={`team-card ${teamCss}`}>
            <h3>{props.team}</h3>
        </div>
    )
}