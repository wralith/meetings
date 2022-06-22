export default function Layout(props) {
  return (
    <div className="py-10 min-w-full min-h-[100vh] flex justify-center items-center text-center">
      {props.children}
    </div>
  )
}
