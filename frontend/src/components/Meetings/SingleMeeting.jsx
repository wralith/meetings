import { deleteToAPI } from '../../handlers/formApi'
import { useState, useEffect } from 'react'
import MeetingModal from './MeetingModal'

export default function SingleMeeting({ meeting }) {
  const [isUpdate, setIsUpdate] = useState(false)

  useEffect(() => {
    setIsUpdate(isUpdate)
  }, [isUpdate])

  const onDelete = () => {
    deleteToAPI(meeting.id)
  }

  const date = Date.parse(meeting.created_at)
  const dateDay = new Date(date).toDateString()
  const dateHour = new Date(date).toTimeString().split(' ')

  // const dateCreated = new Date(Date.parse(meeting.created_at));
  return (
    <section className="flex flex-col gap-1 text-left">
      {isUpdate && <MeetingModal id={meeting.id} click={() => setIsUpdate(!isUpdate)} method="put" />}
      <span className="top-3 relative badge badge-secondary">New</span>
      <h2 className="font-bold w-full bg-accent p-3">{meeting.title}</h2>
      <div className="flex flex-row justify-between items-center">
        <p className="p-3">{meeting.body}</p>
        <div className="flex flex-col text-right">
          <small className="px-3">{dateDay}</small>
          <small className="px-3">{dateHour[0]}</small>
        </div>
      </div>
      <div className="btn-group justify-end px-3">
        <button
          onClick={() => {
            setIsUpdate(!isUpdate)
          }}
          className="btn-sm btn-outline btn">
          Update
        </button>
        <button onClick={onDelete} className="btn-sm btn-outline btn btn-error">
          Delete
        </button>
      </div>
    </section>
  )
}
