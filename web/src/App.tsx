
import { Routes, Route, Outlet } from 'react-router-dom'
import routes from './router'

function App() {

  return (
    <>
      <Routes>
        {routes?.map((route, index) => (
          <Route
            key={index}
            path={route.path}
            element={route.element}
            children={route.children?.map((child, childIndex) => (
              <Route
                key={childIndex}
                path={child.path}
                element={child.element}
              />
            ))}
          />
        ))}
      </Routes>
      <Outlet /> {/* 用于嵌套路由渲染 */}
    </>
  )
}

export default App
