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

    // Check if the request is for login or register routes
    const isLoginOrRegisterRoute = originalRequest.url.includes("/login") || originalRequest.url.includes("/register");

    // If the status is 401 and the route is not login/register, handle token refresh
    if (error.response?.status === 401 && !isLoginOrRegisterRoute && !originalRequest._retry) {
      originalRequest._retry = true;
      try {
        // Attempt to refresh the access token
        const newToken = await refreshAccessToken();

        // Set the new token in the request header
        originalRequest.headers['Authorization'] = `Bearer ${newToken}`;

        // Retry the original request with the new token
        return api(originalRequest);
      } catch (err) {
        // If token refresh fails, reject the promise
        return Promise.reject(err);
      }
    }

    // If the error is not a 401 or the request is for login/register, just reject
    return Promise.reject(error);
  }
);

export default api;
