import { Database, Table, ChevronRight, ChevronDown } from 'lucide-react'
import { useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { useNavigate } from 'react-router-dom'
import { databaseApi } from '@/services/api'
import { cn } from '@/utils/cn'

interface SidebarProps {
  selectedDatabase?: string
  selectedTable?: string
  onSelectDatabase: (db: string) => void
  onSelectTable: (db: string, table: string) => void
}

export default function Sidebar({ selectedDatabase, selectedTable, onSelectDatabase, onSelectTable }: SidebarProps) {
  const [expandedDbs, setExpandedDbs] = useState<Set<string>>(new Set())

  const { data: databases, isLoading } = useQuery({
    queryKey: ['databases'],
    queryFn: async () => {
      const response = await databaseApi.getDatabases()
      return response.data
    },
  })

  const toggleDatabase = (dbName: string) => {
    const newExpanded = new Set(expandedDbs)
    if (newExpanded.has(dbName)) {
      newExpanded.delete(dbName)
    } else {
      newExpanded.add(dbName)
    }
    setExpandedDbs(newExpanded)
  }

  return (
    <div className="w-64 bg-card border-r border-border h-full overflow-y-auto">
      <div className="p-4 border-b border-border">
        <h2 className="text-lg font-semibold">Databases</h2>
      </div>
      {isLoading ? (
        <div className="p-4 text-muted-foreground">Loading...</div>
      ) : (
        <div className="py-2">
          {databases?.map((db) => {
            const isExpanded = expandedDbs.has(db.name)
            const isSelected = selectedDatabase === db.name
            return (
              <div key={db.name}>
                <div
                  className={cn(
                    "flex items-center gap-2 px-4 py-2 cursor-pointer hover:bg-accent",
                    isSelected && "bg-accent"
                  )}
                  onClick={() => {
                    toggleDatabase(db.name)
                    onSelectDatabase(db.name)
                  }}
                >
                  {isExpanded ? (
                    <ChevronDown className="w-4 h-4" />
                  ) : (
                    <ChevronRight className="w-4 h-4" />
                  )}
                  <Database className="w-4 h-4" />
                  <span className="flex-1 truncate">{db.name}</span>
                </div>
                {isExpanded && (
                  <DatabaseTables
                    dbName={db.name}
                    selectedTable={selectedTable}
                    onSelectTable={onSelectTable}
                  />
                )}
              </div>
            )
          })}
        </div>
      )}
    </div>
  )
}

interface DatabaseTablesProps {
  dbName: string
  selectedTable?: string
  onSelectTable: (db: string, table: string) => void
}

function DatabaseTables({ dbName, selectedTable, onSelectTable }: DatabaseTablesProps) {
  const navigate = useNavigate()
  const { data: tables, isLoading } = useQuery({
    queryKey: ['tables', dbName],
    queryFn: async () => {
      const response = await databaseApi.getTables(dbName)
      return response.data
    },
    enabled: !!dbName,
  })

  if (isLoading) {
    return <div className="pl-8 py-1 text-sm text-muted-foreground">Loading tables...</div>
  }

  return (
    <div>
      {tables?.map((table) => {
        const isSelected = selectedTable === table.name
        return (
          <div
            key={table.name}
            className={cn(
              "flex items-center gap-2 px-4 py-1.5 pl-12 cursor-pointer hover:bg-accent text-sm",
              isSelected && "bg-accent font-medium"
            )}
            onClick={() => {
              onSelectTable(dbName, table.name)
              navigate(`/database/${dbName}/table/${table.name}`)
            }}
          >
            <Table className="w-3.5 h-3.5" />
            <span className="flex-1 truncate">{table.name}</span>
            <span className="text-xs text-muted-foreground">{table.rows}</span>
          </div>
        )
      })}
    </div>
  )
}

