import axios from "axios";

const API_URL = import.meta.env.VITE_API_URL || "http://localhost:5000";

const client = axios.create({
  baseURL: `${API_URL}/api`,
  headers: { "Content-Type": "application/json" },
});

export async function getTasks() {
  const { data } = await client.get("/tasks");
  return data;
}

export async function createTask(task) {
  const { data } = await client.post("/tasks", task);
  return data;
}

export async function updateTask(id, task) {
  const { data } = await client.put(`/tasks/${id}`, task);
  return data;
}

export async function deleteTask(id) {
  await client.delete(`/tasks/${id}`);
}
