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

const teams = [
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
    "Vancouver Titans",
    "Seoul Infernal",
    "Guangzhou Charge",
    "Hangzhou Spark",
    "Shanghai Dragons",
    "Dallas Fuel",
    "Seoul Dynasty",
    // "Chengdu Hunters"
]

// TODO fix styles

export default function Combined() {
    const [items, setItems] = useState(teams)
    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
          })
        )
    return (
        <div 
            // style={{
            // display: 'flex',
            // flexDirection: 'column',
            // flexWrap: 'wrap',
            // height: '600px',
            // width: '700px',
            // }}
        >
           <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
            >
                <SortableContext
                    items={items}
                    strategy={verticalListSortingStrategy}
                >
                        {items.map((id) => <TeamCard id={id} />)}
                </SortableContext>
            </DndContext>
        </div>
    )

    function handleDragEnd(event: { active: any; over: any; }) {
        const {active, over} = event;
        
        if (active.id !== over.id) {
          setItems((is: string | any[]) => {
            const oldIndex = is.indexOf(active.id);
            const newIndex = is.indexOf(over.id);
            
            return arrayMove(items, oldIndex, newIndex);
          });
        }
    }
}