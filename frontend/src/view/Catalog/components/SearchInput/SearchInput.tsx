import React, { useState } from 'react'

interface SearchInputProps {
  onSubmit?: (value: string) => void
}

export default ({ onSubmit }: SearchInputProps) => {
  const [value, setValue] = useState('')

  const handleChange: React.ChangeEventHandler<HTMLInputElement> = (event) => {
    event && setValue(event.target.value)
  }

  const handleSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    onSubmit && onSubmit(value)
    event.preventDefault()
  }

  return (
    <form onSubmit={handleSubmit}>
      {' '}
      <label>
        <span>Search: </span>
        <input type="text" value={value} onChange={handleChange} />{' '}
      </label>
      <input type="submit" value="Search" />
    </form>
  )
}
