import { 
  createBrowserRouter,
  createRoutesFromElements,
  Route,
  RouterProvider,
} from 'react-router-dom'
import Home from './Pages/Home'
import './App.css'
import Build from './Pages/Build'
import Login from './Pages/Login'
import Register from './Pages/Register'
import SplitRegion from './Pages/Build/SplitRegion'
import Combined from './Pages/Build/Combined'
import NARegion from './Pages/Build/NARegion'
import APACRegion from './Pages/Build/APACRegion'


const router = createBrowserRouter(
  createRoutesFromElements(
    <Route path="/" element={<Home />}>
      <Route path="build" element={<Build />}>
        <Route path="split-region" element={<SplitRegion />} />
        <Route path="combined" element={<Combined />} />
        <Route path="na" element={<NARegion />} />
        <Route path="apac" element={<APACRegion />} />
      </Route>
      <Route path='accounts'>
        <Route path='login' element={<Login />} />
        <Route path='register' element={<Register />} />
      </Route>
    </Route>
  )
)

function App() {

  return (
    <>
      <RouterProvider router={router}/>
    </>
  );
}


export default App
