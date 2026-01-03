import { useQuery } from '@tanstack/react-query'
import { databaseApi } from '@/services/api'
import { Loader2 } from 'lucide-react'
import { useState } from 'react'
import { useParams } from 'react-router-dom'

export default function TableView() {
  const { db: database, table } = useParams<{ db: string; table: string }>()
  
  if (!database || !table) {
    return <div className="p-4 text-muted-foreground">Invalid database or table</div>
  }
  const [page, setPage] = useState(1)
  const pageSize = 50

  const { data: tableData, isLoading, error: tableDataError } = useQuery({
    queryKey: ['tableData', database, table, page],
    queryFn: async () => {
      const response = await databaseApi.getTableData(database, table, {
        page,
        pageSize,
      })
      return response.data
    },
    enabled: !!database && !!table,
  })

  const { data: structure } = useQuery({
    queryKey: ['tableStructure', database, table],
    queryFn: async () => {
      const response = await databaseApi.getTableStructure(database, table)
      return response.data
    },
    enabled: !!database && !!table,
  })

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-64">
        <Loader2 className="w-8 h-8 animate-spin text-primary" />
      </div>
    )
  }

  if (tableDataError) {
    return (
      <div className="p-4 text-destructive">
        Error loading table data: {tableDataError instanceof Error ? tableDataError.message : 'Unknown error'}
      </div>
    )
  }

  if (!tableData) {
    return <div className="p-4 text-muted-foreground">No data available</div>
  }

  // Ensure columns and rows are arrays
  const columns = tableData.columns || []
  const rows = tableData.rows || []
  const totalPages = Math.ceil((tableData.total || 0) / pageSize)

  return (
    <div className="p-6 space-y-6">
      <div>
        <h2 className="text-2xl font-bold mb-2">
          Table: {table}
        </h2>
        <p className="text-muted-foreground">
          Database: {database} | Total rows: {tableData.total}
        </p>
      </div>

      {structure && structure.columns && structure.columns.length > 0 && (
        <div className="bg-card border border-border rounded-lg p-4">
          <h3 className="text-lg font-semibold mb-4">Structure</h3>
          <div className="overflow-x-auto">
            <table className="w-full border-collapse">
              <thead className="bg-muted">
                <tr>
                  <th className="px-4 py-2 text-left border border-border">Column</th>
                  <th className="px-4 py-2 text-left border border-border">Type</th>
                  <th className="px-4 py-2 text-left border border-border">Null</th>
                  <th className="px-4 py-2 text-left border border-border">Key</th>
                  <th className="px-4 py-2 text-left border border-border">Default</th>
                  <th className="px-4 py-2 text-left border border-border">Extra</th>
                </tr>
              </thead>
              <tbody>
                {structure.columns.map((col, idx) => (
                  <tr key={idx} className="hover:bg-accent/50">
                    <td className="px-4 py-2 border border-border font-medium">{col.name}</td>
                    <td className="px-4 py-2 border border-border">{col.type}</td>
                    <td className="px-4 py-2 border border-border">{col.null}</td>
                    <td className="px-4 py-2 border border-border">{col.key || '-'}</td>
                    <td className="px-4 py-2 border border-border">
                      {col.default || <span className="text-muted-foreground italic">NULL</span>}
                    </td>
                    <td className="px-4 py-2 border border-border">{col.extra || '-'}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      )}

      <div className="bg-card border border-border rounded-lg">
        <div className="p-4 border-b border-border">
          <h3 className="text-lg font-semibold">Data</h3>
        </div>
        <div className="overflow-x-auto">
          <table className="w-full border-collapse">
            <thead className="bg-muted sticky top-0">
              <tr>
                {columns.map((col, idx) => (
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
              {rows.length === 0 ? (
                <tr>
                  <td colSpan={columns.length} className="px-4 py-8 text-center text-muted-foreground">
                    No data available
                  </td>
                </tr>
              ) : (
                rows.map((row, rowIdx) => (
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
                ))
              )}
            </tbody>
          </table>
        </div>
        {totalPages > 1 && (
          <div className="p-4 border-t border-border flex items-center justify-between">
            <div className="text-sm text-muted-foreground">
              Page {page} of {totalPages} ({tableData.total} total rows)
            </div>
            <div className="flex gap-2">
              <button
                onClick={() => setPage((p) => Math.max(1, p - 1))}
                disabled={page === 1}
                className="px-3 py-1 border border-border rounded-md hover:bg-accent disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Previous
              </button>
              <button
                onClick={() => setPage((p) => Math.min(totalPages, p + 1))}
                disabled={page === totalPages}
                className="px-3 py-1 border border-border rounded-md hover:bg-accent disabled:opacity-50 disabled:cursor-not-allowed"
              >
                Next
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  )
}

