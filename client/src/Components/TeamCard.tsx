import './Css/TeamCard.css'
import './Css/TeamStyles.css'
import { useSortable } from '@dnd-kit/sortable'
import {CSS} from '@dnd-kit/utilities'
import { ListContext } from '../Pages/Build'

// TODO figure out how to make transitions work smoothly


export default function TeamCard(props: {id: string, listContext: ListContext, breakId: number}) {
    let teamCss = props.id.replaceAll(' ', '').toLowerCase()
    const {
        attributes,
        listeners,
        setNodeRef,
        transform,
        transition,
        isDragging,
      } = useSortable({id: props.id,
            transition: null});
      const style = {
        transform: CSS.Transform.toString(transform),
        transition,
        opacity: isDragging ? 0.5 : 1
      };
    return (
      <>
        <div className={`team-card ${teamCss}`} style={style} ref={setNodeRef} {...attributes} {...listeners}>
            <h3>{props.id}</h3>
        </div>
      </>
    )
}