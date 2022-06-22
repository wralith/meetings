import BlurOverlay from '../UI/BlurOverlay'
import MeetingForm from './MeetingForm'

export default function MeetingModal(props) {
  return (
    <>
      <BlurOverlay onClick={props.click} />
      <div className="fixed sm:w-1/2 w-[90%] z-50 mt-40 p-10 shadow-xl rounded-2xl bg-base-100">
        <MeetingForm id={props.id} method={props.method} />
      </div>
    </>
  )
}
