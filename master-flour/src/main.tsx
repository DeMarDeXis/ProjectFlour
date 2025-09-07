import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <App />
  </StrictMode>,
)

// TODO: hacer correctamente el deploy (10.08.2025)
// TODO: src/components/Notifications.tsx (13.08.2025)
//  retries for websocket
// TODO: Si apiFetch funciona bien, hacer todos fetches con el
// TODO: correctamente terminar el session si el usuario no esta logueado y no tenga token (10.08.2025)
//  Si es un error 401, redirigir a la pagina de login
//  Entonces, si funciona bien, quitalo

// TODO[Deploy]:
//  заменить некоторые значения на переменные окружения
//  pensar sobre mejoras para docker que estan en el docs