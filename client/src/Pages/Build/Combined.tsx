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
import { ListContext } from '../Build';
import { useOutletContext } from 'react-router-dom';
import TierBreak from '../../Components/TierBreak';

// TODO fix bottom margin

export default function Combined() {
    const lc = useOutletContext<ListContext>()
    const [list, setList] = lc
    const {combined} = list
    const teams = combined
    const [items, setItems] = useState(teams)
    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
          })
        )
    let breakId = -1
    return (
        <div>
           <DndContext
                sensors={sensors}
                collisionDetection={closestCenter}
                onDragEnd={handleDragEnd}
            >
                <SortableContext
                    items={items}
                    strategy={verticalListSortingStrategy}
                >
                        {items.map((id) => {
                            breakId++
                            return (
                                <>
                                <TeamCard listContext={lc} id={id} breakId={breakId}/>
                                {/* should not render the very last tier break */}
                                {breakId != items.length-1 ? <TierBreak listContext={lc} id={breakId} /> : null}
                                </>
                                )
                        })}
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
            
            let narr = arrayMove(items, oldIndex, newIndex);
            setList({
                format: 'combined',
                na: list.na,
                apac: list.apac,
                combined: narr,
                breaks: list.breaks
            })
            return narr
          });
        }
    }
}