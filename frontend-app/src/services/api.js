import axios from 'axios';

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL, // Backend URL from .env
});

export const registerUser = async (userData) => {
  const response = await api.post('/register', userData);
  return response.data;
};

export const loginUser = async (loginData) => {
  const response = await api.post('/login', loginData);
  return response.data;
};

export default api;
