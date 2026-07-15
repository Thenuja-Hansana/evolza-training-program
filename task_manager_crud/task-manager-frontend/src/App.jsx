import { useEffect, useState } from 'react'
import './index.css'

const API_BASE = 'http://localhost:8080'

function App() {
  const [tasks, setTasks] = useState([])
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState('pending')

  const fetchTasks = async () => {
    const res = await fetch(`${API_BASE}/tasks`)
    const data = await res.json()
    setTasks(data.data || [])
  }

  useEffect(() => {
    fetchTasks()
  }, [])

  const handleSubmit = async (e) => {
    e.preventDefault()

    await fetch(`${API_BASE}/tasks`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ title, description, status }),
    })

    setTitle('')
    setDescription('')
    setStatus('pending')
    fetchTasks()
  }

  return (
    <div>
      <h1>Task Manager</h1>

      <form onSubmit={handleSubmit}>
        <div>
          <label>Task Name: </label>
          <input value={title} onChange={(e) => setTitle(e.target.value)} />
        </div>

        <div>
          <label>Description: </label>
          <input value={description} onChange={(e) => setDescription(e.target.value)} />
        </div>

        <div>
          <label>Status: </label>
          <select value={status} onChange={(e) => setStatus(e.target.value)}>
            <option value="pending">Pending</option>
            <option value="completed">Completed</option>
            <option value="not started">Not Started</option>
          </select>
        </div>

        <button type="submit">Add Task</button>
      </form>

      <h2>All Tasks</h2>
      <ul>
        {tasks.map((task) => (
          <li key={task.id}>
            <span>
              <strong>{task.title}</strong> — {task.description} — [{task.status}]
            </span>
          </li>
        ))}
      </ul>
    </div>
  )
}

export default App