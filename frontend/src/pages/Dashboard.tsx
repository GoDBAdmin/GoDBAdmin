import { useOutletContext } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { databaseApi } from '@/services/api'
import { Database, Table, FileText } from 'lucide-react'
import { Link } from 'react-router-dom'

interface DashboardContext {
  selectedDatabase?: string
  selectedTable?: string
}

export default function Dashboard() {
  const { selectedDatabase, selectedTable } = useOutletContext<DashboardContext>()

  const { data: databases, isLoading } = useQuery({
    queryKey: ['databases'],
    queryFn: async () => {
      const response = await databaseApi.getDatabases()
      return response.data
    },
  })

  const { data: tables } = useQuery({
    queryKey: ['tables', selectedDatabase],
    queryFn: async () => {
      if (!selectedDatabase) return []
      const response = await databaseApi.getTables(selectedDatabase)
      return response.data
    },
    enabled: !!selectedDatabase,
  })

  return (
    <div className="p-6 space-y-6">
      <div>
        <h2 className="text-2xl font-bold mb-2">Dashboard</h2>
        <p className="text-muted-foreground">
          Welcome to GoDBAdmin. Select a database from the sidebar to get started.
        </p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div className="p-6 bg-card border border-border rounded-lg">
          <div className="flex items-center gap-3 mb-2">
            <Database className="w-6 h-6 text-primary" />
            <h3 className="text-lg font-semibold">Databases</h3>
          </div>
          {isLoading ? (
            <p className="text-muted-foreground">Loading...</p>
          ) : (
            <p className="text-3xl font-bold">{databases?.length || 0}</p>
          )}
        </div>

        <div className="p-6 bg-card border border-border rounded-lg">
          <div className="flex items-center gap-3 mb-2">
            <Table className="w-6 h-6 text-primary" />
            <h3 className="text-lg font-semibold">Tables</h3>
          </div>
          <p className="text-3xl font-bold">{tables?.length || 0}</p>
        </div>

        <div className="p-6 bg-card border border-border rounded-lg">
          <div className="flex items-center gap-3 mb-2">
            <FileText className="w-6 h-6 text-primary" />
            <h3 className="text-lg font-semibold">Quick Actions</h3>
          </div>
          <div className="mt-4 space-y-2">
            <Link
              to="/query"
              className="block text-sm text-primary hover:underline"
            >
              Open Query Editor
            </Link>
          </div>
        </div>
      </div>

      {selectedDatabase && (
        <div className="mt-6 p-6 bg-card border border-border rounded-lg">
          <h3 className="text-lg font-semibold mb-4">
            Database: {selectedDatabase}
          </h3>
          {selectedTable && (
            <p className="text-muted-foreground">
              Selected table: <span className="font-medium">{selectedTable}</span>
            </p>
          )}
        </div>
      )}
    </div>
  )
}

