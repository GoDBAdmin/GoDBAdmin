import { useEffect } from 'react'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import Login from './pages/Login'
import Dashboard from './pages/Dashboard'
import QueryPage from './pages/QueryPage'
import MainLayout from './components/Layout/MainLayout'
import TableView from './components/DatabaseBrowser/TableView'

function PrivateRoute({ children }: { children: React.ReactNode }) {
  const token = localStorage.getItem('token')
  return token ? <>{children}</> : <Navigate to="/login" replace />
}

function App() {
  useEffect(() => {
    // Apply saved theme
    const theme = localStorage.getItem('theme')
    if (theme === 'dark') {
      document.documentElement.classList.add('dark')
    }
  }, [])

  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route
          path="/"
          element={
            <PrivateRoute>
              <MainLayout />
            </PrivateRoute>
          }
        >
          <Route index element={<Navigate to="/dashboard" replace />} />
          <Route path="dashboard" element={<Dashboard />} />
          <Route path="query" element={<QueryPage />} />
          <Route
            path="database/:db/table/:table"
            element={<TableView />}
          />
        </Route>
      </Routes>
    </BrowserRouter>
  )
}

export default App

