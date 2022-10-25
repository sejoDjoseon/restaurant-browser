import axios, {
  AxiosError,
  AxiosInstance,
  AxiosRequestConfig,
  AxiosResponse,
} from 'axios'
import { AppError } from 'models/AppError'

declare module 'axios' {
  interface AxiosResponse<T = any> extends Promise<T> {}
}

const SERVER_URL = `http://localhost:8000`

abstract class HttpClient {
  protected readonly axiosInstance: AxiosInstance

  protected constructor() {
    this.axiosInstance = axios.create({ baseURL: SERVER_URL, timeout: 30000 })
    this.initializeInterceptors()
  }

  private initializeInterceptors = () => {
    // this.axiosInstance.interceptors.request.use(this.logRequest)
    // this.axiosInstance.interceptors.response.use(this.logResponse)

    this.axiosInstance.interceptors.response.use(
      this.handleResponse,
      this.handleError,
    )
  }

  private handleResponse = ({ data }: AxiosResponse) => data

  private handleError = (
    error: AxiosError<{ message?: string; type?: string; extra?: any }>,
  ) => {
    const out = new AppError()

    if (error.response) {
      out.message = error.response.data?.message || ''
      out.extra = error.response.data?.extra
      out.code = error.response.status?.toString()
    } else {
      out.message = error.message
      out.code = error.code
    }

    switch (out.code) {
      case '400':
      case '401':
      case '403':
      case '404':
        return Promise.reject(out)
      case 'ECONNABORTED':
        console.warn('connection error', out)
        break
      default:
        console.error('Backend error: ', out)
        break
    }

    out.handled = true
    throw out
  }

  private logRequest = (r: AxiosRequestConfig) => {
    const data = JSON.stringify(r.data)
    console.log(`HTTP ==> ${r.method} ${r.url}  data:${data}`)
    return r
  }

  private logResponse = (r: AxiosResponse) => {
    const data = JSON.stringify(r.data)
    console.log(
      `HTTP <== ${r.config.method} ${r.config.url} status:${r.status} response:<${data}>`,
    )
    return r
  }
}

export default HttpClient
