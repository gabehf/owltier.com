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

// TODO fix styles

export default function Combined() {
    const lc = useOutletContext<ListContext>()
    const [list, setList] = lc
    const {apac} = list
    const [items, setItems] = useState(apac)
    const sensors = useSensors(
        useSensor(PointerSensor),
        useSensor(KeyboardSensor, {
            coordinateGetter: sortableKeyboardCoordinates,
          })
        )
    let breakId = 0
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
                            breakId += 1
                            return <TeamCard id={id} listContext={lc} breakId={breakId} />
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
                format: 'apac',
                na: list.na,
                apac: narr,
                breaks: list.breaks
            })
            return narr
          });
        }
    }
}