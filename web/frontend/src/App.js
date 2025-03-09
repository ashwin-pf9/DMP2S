import React from "react";
import { BrowserRouter as Router, Routes, Route, Navigate } from "react-router-dom";
import Login from "./pages/login";
import Register from "./pages/registration";
import Dashboard from "./pages/dashboard";
import PipelineStages from "./pages/pipeline_stages";
import PipelineCreator from "./components/pipeline_creator";
import EditProfile from "./components/user_profile";

const PrivateRoute = ({ element }) => {
  return localStorage.getItem("token") ? element : <Navigate to="/login" />;
};

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/dashboard" element={<PrivateRoute element={<Dashboard />} />} />
        <Route path="/pipeline/create" element={<PrivateRoute element={<PipelineCreator />} />} />
        <Route path="/pipeline/:id" element={<PipelineStages />} />
        <Route path="//edit-profile" element={<EditProfile />} />
        <Route path="*" element={<Navigate to="/login" />} />
      </Routes>
    </Router>
  );
}

//<PipelineCreator onCreate={handleCreatePipeline} /> 

export default App;
