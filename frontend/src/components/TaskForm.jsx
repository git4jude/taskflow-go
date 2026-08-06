import { useEffect, useState } from "react";

const emptyTask = {
  title: "",
  description: "",
  status: "todo",
  priority: "medium",
  assigned_to: "",
  due_date: "",
};

function toFormValues(task) {
  if (!task) return emptyTask;
  return {
    title: task.title || "",
    description: task.description || "",
    status: task.status || "todo",
    priority: task.priority || "medium",
    assigned_to: task.assigned_to || "",
    due_date: task.due_date ? task.due_date.slice(0, 10) : "",
  };
}

export default function TaskForm({ editingTask, onSubmit, onCancel, submitting }) {
  const [values, setValues] = useState(toFormValues(editingTask));

  useEffect(() => {
    setValues(toFormValues(editingTask));
  }, [editingTask]);

  function handleChange(e) {
    const { name, value } = e.target;
    setValues((prev) => ({ ...prev, [name]: value }));
  }

  function handleSubmit(e) {
    e.preventDefault();
    const payload = {
      ...values,
      due_date: values.due_date || null,
    };
    onSubmit(payload);
  }

  return (
    <form className="card task-form" onSubmit={handleSubmit}>
      <h2 className="card-title">{editingTask ? "Edit Task" : "Add Task"}</h2>

      <div className="form-group">
        <label htmlFor="title">Title</label>
        <input
          id="title"
          name="title"
          type="text"
          value={values.title}
          onChange={handleChange}
          required
        />
      </div>

      <div className="form-group">
        <label htmlFor="description">Description</label>
        <textarea
          id="description"
          name="description"
          rows={3}
          value={values.description}
          onChange={handleChange}
        />
      </div>

      <div className="form-row">
        <div className="form-group">
          <label htmlFor="status">Status</label>
          <select id="status" name="status" value={values.status} onChange={handleChange}>
            <option value="todo">To Do</option>
            <option value="in_progress">In Progress</option>
            <option value="done">Done</option>
          </select>
        </div>

        <div className="form-group">
          <label htmlFor="priority">Priority</label>
          <select id="priority" name="priority" value={values.priority} onChange={handleChange}>
            <option value="low">Low</option>
            <option value="medium">Medium</option>
            <option value="high">High</option>
          </select>
        </div>
      </div>

      <div className="form-row">
        <div className="form-group">
          <label htmlFor="assigned_to">Assigned To</label>
          <input
            id="assigned_to"
            name="assigned_to"
            type="text"
            value={values.assigned_to}
            onChange={handleChange}
          />
        </div>

        <div className="form-group">
          <label htmlFor="due_date">Due Date</label>
          <input
            id="due_date"
            name="due_date"
            type="date"
            value={values.due_date}
            onChange={handleChange}
          />
        </div>
      </div>

      <div className="form-actions">
        {editingTask && (
          <button type="button" className="btn btn-secondary" onClick={onCancel}>
            Cancel
          </button>
        )}
        <button type="submit" className="btn btn-primary" disabled={submitting}>
          {editingTask ? "Save Changes" : "Add Task"}
        </button>
      </div>
    </form>
  );
}
