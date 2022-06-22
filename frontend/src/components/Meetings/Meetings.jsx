import SingleMeeting from './SingleMeeting'


export default function Meetings({isLoading, isError, data, error}) {

  if (isLoading) {
    return (
      <>
        <p>Loading</p>
      </>
    )
  }

  if (isError) {
    return (
      <>
        <p>{error.message}</p>
      </>
    )
  }
  return (
    <>
      {data.map(meeting => (
        <SingleMeeting key={meeting.id} meeting={meeting} />
      ))}
    </>
  )
}
