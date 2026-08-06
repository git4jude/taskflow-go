import { useEffect, useState } from "react";
import TaskForm from "./components/TaskForm";
import TaskList from "./components/TaskList";
import { createTask, deleteTask, getTasks, updateTask } from "./services/taskApi";

function extractErrorMessage(err, fallback) {
  return err?.response?.data?.error || fallback;
}

export default function App() {
  const [tasks, setTasks] = useState([]);
  const [loading, setLoading] = useState(true);
  const [listError, setListError] = useState("");
  const [formError, setFormError] = useState("");
  const [editingTask, setEditingTask] = useState(null);
  const [submitting, setSubmitting] = useState(false);

  async function loadTasks() {
    setLoading(true);
    setListError("");
    try {
      const data = await getTasks();
      setTasks(data || []);
    } catch (err) {
      setListError(extractErrorMessage(err, "Failed to load tasks. Is the API running?"));
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadTasks();
  }, []);

  async function handleSubmit(values) {
    setSubmitting(true);
    setFormError("");
    try {
      if (editingTask) {
        await updateTask(editingTask.id, values);
      } else {
        await createTask(values);
      }
      setEditingTask(null);
      await loadTasks();
    } catch (err) {
      setFormError(extractErrorMessage(err, "Failed to save task."));
    } finally {
      setSubmitting(false);
    }
  }

  async function handleDelete(id) {
    if (!window.confirm("Delete this task?")) return;
    setListError("");
    try {
      await deleteTask(id);
      if (editingTask?.id === id) setEditingTask(null);
      await loadTasks();
    } catch (err) {
      setListError(extractErrorMessage(err, "Failed to delete task."));
    }
  }

  function handleEdit(task) {
    setFormError("");
    setEditingTask(task);
  }

  function handleCancel() {
    setFormError("");
    setEditingTask(null);
  }

  return (
    <div className="page">
      <header className="page-header">
        <h1>TaskFlow</h1>
        <p>Simple, focused task tracking for your team.</p>
      </header>

      <main className="container">
        <TaskForm
          editingTask={editingTask}
          onSubmit={handleSubmit}
          onCancel={handleCancel}
          submitting={submitting}
        />
        {formError && <p className="state-message state-error">{formError}</p>}

        <TaskList
          tasks={tasks}
          loading={loading}
          error={listError}
          onEdit={handleEdit}
          onDelete={handleDelete}
        />
      </main>
    </div>
  );
}
