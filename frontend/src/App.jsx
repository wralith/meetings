import Layout from './components/UI/Layout'
import { QueryClient, QueryClientProvider } from 'react-query'
import Meetings from './views/Meetings'

const queryClient = new QueryClient()

function App() {
  return (
    <Layout>
      <QueryClientProvider client={queryClient}>
        <Meetings />
      </QueryClientProvider>
    </Layout>
  )
}

export default App
