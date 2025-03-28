import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import FormPage from "../src/components/FormPage";
import DisplayPage from "../src/components/DisplayPage";

const App: React.FC = () => {
  return (
      <Router>
        <Routes>
          <Route path="/" element={<FormPage />} />
          <Route path="/display" element={<DisplayPage />} />
        </Routes>
      </Router>
  );
};

export default App;