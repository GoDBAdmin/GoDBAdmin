import axios from 'axios'

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8090/api'

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Add token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// Handle auth errors
api.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      window.location.href = '/login'
    }
    return Promise.reject(error)
  }
)

export interface LoginRequest {
  host: string
  port: number
  username: string
  password: string
  database: string
}

export interface LoginResponse {
  token: string
}

export interface Database {
  name: string
}

export interface Table {
  name: string
  engine: string
  rows: number
  size: string
  comment: string
}

export interface Column {
  name: string
  type: string
  null: string
  key: string
  default: string | null
  extra: string
  comment: string
}

export interface Index {
  name: string
  columns: string[]
  unique: boolean
}

export interface TableStructure {
  columns: Column[]
  indexes: Index[]
}

export interface QueryRequest {
  database: string
  query: string
}

export interface QueryResponse {
  columns: string[]
  rows: any[][]
  affected: number
  error?: string
}

export interface TableDataRequest {
  database: string
  table: string
  page: number
  pageSize: number
  sortBy?: string
  sortDir?: string
}

export interface TableDataResponse {
  columns: string[]
  rows: any[][]
  total: number
  page: number
  pageSize: number
}

export const authApi = {
  login: (data: LoginRequest) => api.post<LoginResponse>('/auth/login', data),
}

export const databaseApi = {
  getDatabases: () => api.get<Database[]>('/databases'),
  getTables: (dbName: string) => api.get<Table[]>(`/databases/${dbName}/tables`),
  getTableStructure: (dbName: string, tableName: string) =>
    api.get<TableStructure>(`/databases/${dbName}/tables/${tableName}`),
  getTableData: (dbName: string, tableName: string, params: { page: number; pageSize: number; sortBy?: string; sortDir?: string }) =>
    api.get<TableDataResponse>(`/databases/${dbName}/tables/${tableName}/data`, { params }),
  executeQuery: (data: QueryRequest) => api.post<QueryResponse>('/query', data),
  createTable: (dbName: string, data: { tableName: string; createSQL: string }) =>
    api.post(`/databases/${dbName}/tables`, data),
  dropTable: (dbName: string, tableName: string) =>
    api.delete(`/databases/${dbName}/tables/${tableName}`),
}

export default api

