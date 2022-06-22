import { useForm } from 'react-hook-form'
import { postToAPI, putToAPI } from '../../handlers/formApi'

export default function MeetingForm({ method, id }) {
  // useForm hook implementation from react-hook-form
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm()

  let onSubmit

  if (method == "post") {
    onSubmit = data => {
      postToAPI(data)
      reset()
    }
  }
  if (method == "put") {
    onSubmit = (data) => {
      putToAPI(data, id)
      reset()
    }
  }

  return (
    <div className="z-50">
      <form onSubmit={handleSubmit(onSubmit)} className="flex flex-col gap-3" action="">
        <div className="form-control shadow-md">
          <label className="input-group input-group-vertical">
            <span className="label-text p-2 pl-4 bg-accent">Title</span>
            <input
              {...register('title', { required: true })}
              type="text"
              className="input input-bordered"
            />
            {errors.title && (
              <span className="justify-end py-2 text-sm text-primary">This field is required</span>
            )}
          </label>
        </div>
        <div className="form-control shadow-md">
          <label className="input-group input-group-vertical">
            <span className="label-text p-2 pl-4 bg-accent">Description</span>
            <textarea
              {...register('body', { required: true })}
              className="h-24 input input-bordered"
            />
            {errors.body && (
              <span className="justify-end py-2 text-sm text-primary">This field is required</span>
            )}
          </label>
        </div>
        <div className="form-control shadow-md">
          <label className="input-group input-group-vertical">
            <span className="label-text p-2 pl-4 bg-accent">Date</span>
            <input type="datetime-local"
              // {...register('body', { required: true })}
              className="input input-bordered"
            />
            {errors.body && (
              <span className="justify-end py-2 text-sm text-primary">This field is required</span>
            )}
          </label>
        </div>
        <button type="submit" className="btn btn-primary">
          Add
        </button>
      </form>
    </div>
  )
}
