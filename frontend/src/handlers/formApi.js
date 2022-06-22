import axios from 'axios'

export async function postToAPI(data) {
  await axios.post('http://localhost:8080/meetings', data)
}
export async function putToAPI(data, id) {
  await axios.put(`http://localhost:8080/meetings/${id}`, data)
}
export async function deleteToAPI(id) {
  await axios.delete(`http://localhost:8080/meetings/${id}`)
}
export async function fetchMeetings() {
  const { data } = await axios.get('http://localhost:8080/meetings')
  // console.log(data)
  return data
}