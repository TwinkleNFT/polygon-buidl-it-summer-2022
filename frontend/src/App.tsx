import { useState } from "react";

import Navibar from "./components/Navibar";
import Home from "./pages/Home";
import Playground from "./pages/Playground";
import Footer from "./components/Footer";
import { BrowserRouter, Routes, Route } from "react-router-dom";
function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <BrowserRouter>
        <Navibar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/playground" element={<Playground />} />
        </Routes>
        <Footer />
      </BrowserRouter>
    </>
  );
}

export default App;
