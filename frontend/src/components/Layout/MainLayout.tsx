import { useState } from 'react'
import { Outlet } from 'react-router-dom'
import Header from './Header'
import Sidebar from './Sidebar'

export default function MainLayout() {
  const [selectedDatabase, setSelectedDatabase] = useState<string>()
  const [selectedTable, setSelectedTable] = useState<string>()

  return (
    <div className="h-screen flex flex-col bg-background">
      <Header />
      <div className="flex-1 flex overflow-hidden">
        <Sidebar
          selectedDatabase={selectedDatabase}
          selectedTable={selectedTable}
          onSelectDatabase={setSelectedDatabase}
          onSelectTable={(db, table) => {
            setSelectedDatabase(db)
            setSelectedTable(table)
          }}
        />
        <main className="flex-1 overflow-auto">
          <Outlet context={{ selectedDatabase, selectedTable }} />
        </main>
      </div>
    </div>
  )
}

