import axios from 'axios';

const axiosInstance = axios.create({
    baseURL: process.env.VUE_APP_ROOT_API,
    // withCredentials : true
});

export default axiosInstance;