import TaskItem from "./TaskItem";

export default function TaskList({ tasks, loading, error, onEdit, onDelete }) {
  return (
    <div className="card">
      <h2 className="card-title">Tasks</h2>

      {loading && <p className="state-message">Loading tasks…</p>}
      {error && !loading && <p className="state-message state-error">{error}</p>}
      {!loading && !error && tasks.length === 0 && (
        <p className="state-message">No tasks yet. Add one above to get started.</p>
      )}

      {!loading && !error && tasks.length > 0 && (
        <div className="table-wrapper">
          <table>
            <thead>
              <tr>
                <th>Title</th>
                <th>Description</th>
                <th>Status</th>
                <th>Priority</th>
                <th>Assigned To</th>
                <th>Due Date</th>
                <th>Actions</th>
              </tr>
            </thead>
            <tbody>
              {tasks.map((task) => (
                <TaskItem key={task.id} task={task} onEdit={onEdit} onDelete={onDelete} />
              ))}
            </tbody>
          </table>
        </div>
      )}
    </div>
  );
}
