import { useEffect, useState } from 'react'

const API_BASE = 'http://localhost:8080'

function App() {
  const [tasks, setTasks] = useState([])
  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [status, setStatus] = useState('pending')
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  // Fetch all tasks from the backend
  const fetchTasks = async () => {
    try {
      const res = await fetch(`${API_BASE}/tasks`)
      const data = await res.json()
      setTasks(data.data || [])
    } catch (err) {
      setError('Failed to fetch tasks. Is the Go server running on ' + API_BASE + '?')
    }
  }

  useEffect(() => {
    fetchTasks()
  }, [])

  // Handle creating a new task
  const handleSubmit = async (e) => {
    e.preventDefault()
    if (!title.trim()) {
      setError('Task title is required')
      return
    }

    setLoading(true)
    setError('')

    try {
      const res = await fetch(`${API_BASE}/tasks`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, description, status }),
      })

      if (!res.ok) {
        const errData = await res.json()
        throw new Error(errData.error || 'Failed to create task')
      }

      setTitle('')
      setDescription('')
      setStatus('pending')

      fetchTasks()
    } catch (err) {
      setError(err.message)
    } finally {
      setLoading(false)
    }
  }

  // Handle updating just the status of an existing task
  const handleUpdateStatus = async (task, newStatus) => {
    try {
      await fetch(`${API_BASE}/tasks/${task.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          title: task.title,
          description: task.description,
          status: newStatus,
        }),
      })
      fetchTasks()
    } catch (err) {
      setError('Failed to update task')
    }
  }

  // Handle deleting a task
  const handleDelete = async (id) => {
    try {
      await fetch(`${API_BASE}/tasks/${id}`, { method: 'DELETE' })
      fetchTasks()
    } catch (err) {
      setError('Failed to delete task')
    }
  }

  return (
    <div>
      <h1>Task Manager</h1>

      <form onSubmit={handleSubmit}>
        <div>
          <label>Task Name: </label>
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Enter task name"
          />
        </div>

        <div>
          <label>Description: </label>
          <input
            type="text"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="Enter description"
          />
        </div>

        <div>
          <label>Status: </label>
          <select value={status} onChange={(e) => setStatus(e.target.value)}>
            <option value="pending">Pending</option>
            <option value="completed">Completed</option>
            <option value="not started">Not Started</option>
          </select>
        </div>

        <button type="submit" disabled={loading}>
          {loading ? 'Adding...' : 'Add Task'}
        </button>
      </form>

      {error && <p style={{ color: 'red' }}>{error}</p>}

      <h2>All Tasks</h2>

      {tasks.length === 0 ? (
        <p>No tasks yet.</p>
      ) : (
        <ul>
          {tasks.map((task) => (
            <li key={task.id}>
              <span>
                <strong>{task.title}</strong> — {task.description}
              </span>
              <select
                value={task.status}
                onChange={(e) => handleUpdateStatus(task, e.target.value)}
              >
                <option value="pending">Pending</option>
                <option value="completed">Completed</option>
                <option value="not started">Not Started</option>
              </select>
              <button onClick={() => handleDelete(task.id)}>Delete</button>
            </li>
          ))}
        </ul>
      )}
    </div>
  )
}

export default App