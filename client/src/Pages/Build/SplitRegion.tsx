import { useState } from 'react'
import TeamCard from '../../Components/TeamCard'
import {
    arrayMove,
    SortableContext,
    sortableKeyboardCoordinates,
    verticalListSortingStrategy,
} from '@dnd-kit/sortable'
import {
    DndContext, 
    closestCenter,
    KeyboardSensor,
    PointerSensor,
    useSensor,
    useSensors,
  } from '@dnd-kit/core';

const na = [
     "San Francisco Shock",
     "Houston Outlaws",
     "Boston Uprising",
     "Los Angeles Gladiators",
     "New York Excelsior",
     "Washington Justice",
     "Vegas Eternal",
     "Toronto Defiant",
     "Atlanta Reign",
     "Florida Mayhem",
     "London Spitfire",
     "Los Angeles Valiant",
     "Vancouver Titans"
]

const apac = [
    "Seoul Infernal",
    "Guangzhou Charge",
    "Hangzhou Spark",
    "Shanghai Dragons",
    "Dallas Fuel",
    "Seoul Dynasty",
    // "Chengdu Hunters",
]

export default function SplitRegion() {
    const [nateams, setNA] = useState(na)
    const [apacteams, setAPAC] = useState(apac)
    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
          })
        )
    return (
        <div style={{
            display: 'flex',
            width: '800px',
            justifyContent: 'space-between',
            }}>
            <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
            >
                <div className="na">
                    <SortableContext
                        items={nateams}
                        strategy={verticalListSortingStrategy}
                    >
                            {nateams.map((id) => <TeamCard id={id} />)}
                    </SortableContext>
                </div>
            </DndContext>
            <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEndApac}
            >
                <div className="apac">
                    <SortableContext
                        items={apacteams}
                        strategy={verticalListSortingStrategy}
                    >
                            {apacteams.map((id) => <TeamCard id={id} />)}
                    </SortableContext>
                </div>
            </DndContext>
        </div>
    )

    function handleDragEnd(event: { active: any; over: any; }) {
        const {active, over} = event;
        
        if (active.id !== over.id) {
          setNA((items: string | any[]) => {
            const oldIndex = items.indexOf(active.id);
            const newIndex = items.indexOf(over.id);
            
            return arrayMove(nateams, oldIndex, newIndex);
          });
        }
    }
    
    function handleDragEndApac(event: { active: any; over: any; }) {
        const {active, over} = event;
        
        if (active.id !== over.id) {
          setAPAC((items: string | any[]) => {
            const oldIndex = items.indexOf(active.id);
            const newIndex = items.indexOf(over.id);
            
            return arrayMove(apacteams, oldIndex, newIndex);
          });
        }
    }
    
}