import axios from 'axios';

const api = axios.create({
  baseURL: 'http://localhost:5000/api',
  timeout: 10000,
  withCredentials: true,
});

// Token helpers
const getAccessToken = () => localStorage.getItem('accessToken');

const refreshAccessToken = async () => {
  const response = await axios.post('/refresh');
  const { accessToken } = response.data;

  localStorage.setItem('accessToken', accessToken);
  return accessToken;
};

// Request interceptor
api.interceptors.request.use(
  (config) => {
    const token = getAccessToken();
    if (token) {
      config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;
    if (error.response?.status === 401 && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        const newToken = await refreshAccessToken();
        originalRequest.headers['Authorization'] = `Bearer ${newToken}`;
        return api(originalRequest);
      } catch (err) {
        console.error('Token refresh failed', err);
        // Optionally handle logout here
        return Promise.reject(err);
      }
    }

    return Promise.reject(error);
  }
);

export default api;
