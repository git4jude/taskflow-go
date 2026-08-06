const STATUS_LABELS = {
  todo: "To Do",
  in_progress: "In Progress",
  done: "Done",
};

const PRIORITY_LABELS = {
  low: "Low",
  medium: "Medium",
  high: "High",
};

function formatDate(value) {
  if (!value) return "—";
  const date = new Date(value);
  if (Number.isNaN(date.getTime())) return "—";
  return date.toLocaleDateString(undefined, {
    year: "numeric",
    month: "short",
    day: "numeric",
  });
}

export default function TaskItem({ task, onEdit, onDelete }) {
  return (
    <tr>
      <td className="cell-title">{task.title}</td>
      <td className="cell-description">{task.description || "—"}</td>
      <td>
        <span className={`badge badge-status-${task.status}`}>
          {STATUS_LABELS[task.status] || task.status}
        </span>
      </td>
      <td>
        <span className={`badge badge-priority-${task.priority}`}>
          {PRIORITY_LABELS[task.priority] || task.priority}
        </span>
      </td>
      <td>{task.assigned_to || "—"}</td>
      <td>{formatDate(task.due_date)}</td>
      <td className="cell-actions">
        <button className="btn btn-outline btn-sm" onClick={() => onEdit(task)}>
          Edit
        </button>
        <button className="btn btn-danger btn-sm" onClick={() => onDelete(task.id)}>
          Delete
        </button>
      </td>
    </tr>
  );
}
