export default function BlurOverlay(props) {
  return (
    <div onClick={props.onClick} className="flex justify-center items-center fixed left-0 top-0 min-w-full min-h-full z-10 backdrop-blur-sm backdrop-brightness-25"></div>
  )
}
