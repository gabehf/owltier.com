import TeamCard from '../../Components/TeamCard'

export default function SplitRegion() {
    return (
        <div style={{
            display: 'flex',
            width: '600px',
            justifyContent: 'space-between',
            }}>
            <div className="na">
           <TeamCard team="San Francisco Shock" /> 
           <TeamCard team="Houston Outlaws" /> 
           <TeamCard team="Boston Uprising" /> 
           <TeamCard team="Los Angeles Gladiators" /> 
           <TeamCard team="New York Excelsior" /> 
           <TeamCard team="Washington Justice" /> 
           <TeamCard team="Vegas Eternal" /> 
           <TeamCard team="Toronto Defiant" /> 
           <TeamCard team="Atlanta Reign" /> 
           <TeamCard team="Florida Mayhem" /> 
           <TeamCard team="London Spitfire" /> 
           <TeamCard team="Los Angeles Valiant" /> 
           <TeamCard team="Vancouver Titans" /> 
           </div>
            <div className="apac">
           <TeamCard team="Seoul Infernal" /> 
           <TeamCard team="Guangzhou Charge" /> 
           <TeamCard team="Hangzhou Spark" /> 
           <TeamCard team="Shanghai Dragons" /> 
           <TeamCard team="Dallas Fuel" /> 
           <TeamCard team="Seoul Dynasty" /> 
           <TeamCard team="Chengdu Hunters" /> 
           </div>
        </div>
    )
}