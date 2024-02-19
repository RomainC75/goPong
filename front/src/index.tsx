import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import { AuthProviderWrapper } from "./context/auth.context";
import AuthPage from "./pages/Auth.page";
import IsPrivate from "./components/helper/isPrivate";
import { SocketProviderWrapper } from "./context/socket.context";
import React from "react";
import { GameProvider } from "./hooks/useGame";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);
root.render(
  // <React.StrictMode>
  <BrowserRouter>
    <AuthProviderWrapper>
      <Routes>
        <Route path="/auth" element={<AuthPage />} />
        <Route
          path="/"
          element={
            <IsPrivate>
              <GameProvider>
                <SocketProviderWrapper>
                  <App />
                </SocketProviderWrapper>
              </GameProvider>
            </IsPrivate>
          }
        />
      </Routes>
    </AuthProviderWrapper>
  </BrowserRouter>
  // {/* </React.StrictMode> */}
);

reportWebVitals();
