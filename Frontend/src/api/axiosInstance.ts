import { apiEndpointApi } from '@/config/config';
import axios from 'axios';

const api = axios.create({
  baseURL: apiEndpointApi,
  timeout: 10000,
  withCredentials: true,
});

// Token helpers
const getAccessToken = () => localStorage.getItem('accessToken');


const refreshAccessToken = async () => {
  try {
    const response = await api.post(
      '/refresh',
    );

    const { access_token } = response.data;
    console.log("🚀 ~ refreshaccess_token ~ access_token:", access_token);
    if(!access_token){
      throw "un-auth"
    }
    localStorage.setItem('accessToken', access_token);
    return access_token;
  } catch (error) {
    console.error("Error refreshing access token:", error);
    throw error;
  }
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
      console.log("🚀 ~ originalRequest._retry:", !originalRequest._retry)
      console.log("🚀 ~ isLoginOrRegisterRoute:", isLoginOrRegisterRoute)
      console.log("🚀 ~ error.response?.status:", error.response?.status)
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
