import MeetingModal from '../components/Meetings/MeetingModal'
import MeetingsLayout from '../components/Meetings/MeetingsLayout'
import { useState, useEffect } from 'react'

export default function Home() {
  const [isAdd, setIsAdd] = useState(false)

  useEffect(() => {
    setIsAdd(isAdd)
  }, [isAdd])

  return (
    <div className="flex flex-col gap-5 w-[90%] sm:w-1/2">
      <MeetingsLayout />
      <button
        onClick={() => {
          setIsAdd(!isAdd)
        }}
        className="btn btn-primary">
        Add
      </button>
      {isAdd && <MeetingModal click={() => setIsAdd(!isAdd)} method="post"/>}
    </div>
  )
}
