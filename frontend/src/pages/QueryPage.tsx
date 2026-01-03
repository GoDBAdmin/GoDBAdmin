import { useState } from 'react'
import { useOutletContext } from 'react-router-dom'
import { useMutation, useQuery } from '@tanstack/react-query'
import { databaseApi } from '@/services/api'
import Editor from '@monaco-editor/react'
import { Play, Loader2 } from 'lucide-react'

interface DashboardContext {
  selectedDatabase?: string
}

export default function QueryPage() {
  const { selectedDatabase } = useOutletContext<DashboardContext>()
  const [query, setQuery] = useState('')
  const [currentDb, setCurrentDb] = useState(selectedDatabase || '')

  const { data: databases } = useQuery({
    queryKey: ['databases'],
    queryFn: async () => {
      const response = await databaseApi.getDatabases()
      return response.data
    },
  })

  const queryMutation = useMutation({
    mutationFn: (q: string) => {
      if (!currentDb) {
        throw new Error('Please select a database')
      }
      return databaseApi.executeQuery({ database: currentDb, query: q })
    },
  })

  const handleExecute = () => {
    if (query.trim() && currentDb) {
      queryMutation.mutate(query)
    }
  }

  const result = queryMutation.data?.data

  return (
    <div className="h-full flex flex-col p-6">
      <div className="mb-4 flex items-center gap-4">
        <div className="flex-1">
          <label className="block text-sm font-medium mb-2">Database</label>
          <select
            value={currentDb}
            onChange={(e) => setCurrentDb(e.target.value)}
            className="px-3 py-2 border border-input rounded-md bg-background focus:outline-none focus:ring-2 focus:ring-ring"
          >
            <option value="">Select database</option>
            {databases?.map((db) => (
              <option key={db.name} value={db.name}>
                {db.name}
              </option>
            ))}
          </select>
        </div>
        <button
          onClick={handleExecute}
          disabled={!query.trim() || !currentDb || queryMutation.isPending}
          className="mt-6 px-4 py-2 bg-primary text-primary-foreground rounded-md hover:bg-primary/90 focus:outline-none focus:ring-2 focus:ring-ring disabled:opacity-50 disabled:cursor-not-allowed transition-colors flex items-center gap-2"
        >
          {queryMutation.isPending ? (
            <>
              <Loader2 className="w-4 h-4 animate-spin" />
              Executing...
            </>
          ) : (
            <>
              <Play className="w-4 h-4" />
              Execute
            </>
          )}
        </button>
      </div>

      <div className="flex-1 border border-border rounded-lg overflow-hidden mb-4">
        <Editor
          height="100%"
          defaultLanguage="sql"
          value={query}
          onChange={(value) => setQuery(value || '')}
          theme="vs-dark"
          options={{
            minimap: { enabled: false },
            fontSize: 14,
            wordWrap: 'on',
            automaticLayout: true,
          }}
        />
      </div>

      {queryMutation.isError && (
        <div className="mb-4 p-4 bg-destructive/10 border border-destructive/20 rounded-lg text-destructive">
          {queryMutation.error instanceof Error
            ? queryMutation.error.message
            : 'An error occurred'}
        </div>
      )}

      {result && (
        <div className="flex-1 overflow-auto border border-border rounded-lg">
          {result.error ? (
            <div className="p-4 bg-destructive/10 border-b border-destructive/20 text-destructive">
              {result.error}
            </div>
          ) : result.columns && result.columns.length > 0 ? (
            <div className="overflow-x-auto">
              <table className="w-full border-collapse">
                <thead className="bg-muted sticky top-0">
                  <tr>
                    {result.columns.map((col, idx) => (
                      <th
                        key={idx}
                        className="px-4 py-2 text-left border-b border-border font-semibold"
                      >
                        {col}
                      </th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {result.rows.map((row, rowIdx) => (
                    <tr key={rowIdx} className="hover:bg-accent/50">
                      {row.map((cell, cellIdx) => (
                        <td
                          key={cellIdx}
                          className="px-4 py-2 border-b border-border"
                        >
                          {cell === null ? (
                            <span className="text-muted-foreground italic">NULL</span>
                          ) : (
                            String(cell)
                          )}
                        </td>
                      ))}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          ) : (
            <div className="p-4 text-center text-muted-foreground">
              Query executed successfully. {result.affected > 0 && (
                <span>{result.affected} row(s) affected.</span>
              )}
            </div>
          )}
        </div>
      )}
    </div>
  )
}

