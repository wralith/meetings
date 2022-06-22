import Meetings from "./Meetings"
import { useQuery } from 'react-query'
import {fetchMeetings} from "../../handlers/formApi"

export default function MeetingsLayout() {
  const { isLoading, isError, data, error } = useQuery('meetings', fetchMeetings)

  return (
    <div className="h-[80vh] overflow-auto rounded-t-xl flex flex-col border-base-300 border shadow-md gap-3 pb-3">
      <h1 className="rounded-t-xl py-3 pl-3 text-left font-bold bg-accent">Meetings</h1>
      <Meetings isLoading={isLoading} isError={isError} data={data} error={error} />
    </div>
  )
}
