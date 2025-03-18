import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import usePipelines from "../hooks/use_pipelines";
import PipelineList from "../components/pipeline_list";
import "../styles/dashboard.css";
import TopBar from "../components/top_bar";

const Dashboard = () => {
  const {
    pipelines,
    loading,
    error,
  } = usePipelines();

  const navigate = useNavigate();
  const [userName, setUserName] = useState("User"); // Default value

  useEffect(() => {
    // Fetch user data from local storage or API
    const storedUser = localStorage.getItem("user_name"); // Example: Fetch from local storage
    if (storedUser) {
      setUserName(storedUser);
    }
  }, []);

  const handlePipelineClick = (pipelineId, pipelineName) => {
    navigate(`/pipeline/${pipelineId}`, {state: {pipelineName}});
  };

   // If backend is not reachable, prevent rendering the dashboard
   if (error) {
    if (error.includes("Unauthorized")) {
      alert("Session expired. Redirecting to login...");
      navigate("/login"); // Redirect user to login page
      return null; // Prevent rendering the current component
    }

    if (
      error.includes("Failed to fetch") ||
      error.includes("Backend is unreachable. Please try again later.")
    ) {
      return (
        <div className="dashboard-error">
          <h2>Backend is unreachable</h2>
          <p>Please check your internet connection or try again later.</p>
        </div>
      );
    }
  
    
  }
  

  //This error will occur when server is connected to wifi network
  if(error && error.includes("Invalid pipeline data format")) {
    return (
      <div className="dashboard">
      <div className="top-bar">
        <h1 className="project-name">Manufacturing Pipeline</h1>
        <div className="user-profile">
          <img src="/profile-icon.png" alt="User Profile" className="profile-icon" />
        </div>
      </div>
      <div className="dashboard-error">
        <h2>Failed to fetch pipelines</h2>
        <p>Please try again later.</p>
      </div>
      </div>
    );
  }

  return (
    <div className="dashboard">
       
       <TopBar></TopBar>

       <h2 className="welcome-message">Welcome, {userName}!</h2>

      {loading && <p className="loading-message">Loading pipelines...</p>}
      {error && <p className="error-message">Error: {error}</p>}

      {/* <PipelineCreator onCreate={handleCreatePipeline} /> */}
      <button onClick={() => navigate("/pipeline/create")} className="create-pipeline-button">
      Create Pipeline
    </button>

      <div className="dashboard-content">
        <PipelineList pipelines={pipelines} onSelect={handlePipelineClick} />
      </div>
      </div>
  );
};

export default Dashboard;







// import React, { useEffect, useState } from "react";
// import { fetchPipelines, fetchStages, createPipeline } from "../api/api";
// import PipelineCard from "../components/pipeline_card";
// import StageList from "../components/stage_list";
// import "../styles/dashboard.css";

// const Dashboard = () => {
//   const [pipelines, setPipelines] = useState([]);
//   const [selectedPipeline, setSelectedPipeline] = useState(null);
//   const [stages, setStages] = useState([]);
//   const [clicked, setClicked] = useState(false);
//   const [loading, setLoading] = useState(true);
//   const [error, setError] = useState(null);
//   const [newPipelineName, setNewPipelineName] = useState("");

//   useEffect(() => {
//     loadPipelines();
//   }, []);
//     const loadPipelines = async () => {
//       try {
//         setLoading(true);
//         setError(null);
//         const data = await fetchPipelines();
//         console.log("Fetched Pipelines:", data);

//         if (!data || !Array.isArray(data)) {
//           throw new Error("Invalid pipeline data format");
//         }

//         setPipelines(data);
//       } catch (error) {
//         console.error("Error fetching pipelines:", error);
//         setError(error.message);
//         setPipelines([]);
//       } finally {
//         setLoading(false);
//       }
//     };

   

//   const handlePipelineClick = async (pipelineId) => {
//     if (selectedPipeline === pipelineId) return;

//     setSelectedPipeline(pipelineId);
//     setClicked(true);

//     try {
//       const data = await fetchStages(pipelineId);
//       console.log(`Fetched stages for pipeline ${pipelineId}:`, data);

//       if (!data || !Array.isArray(data)) {
//         throw new Error("Invalid stage data format");
//       }

//       setStages(data);
//     } catch (error) {
//       console.error("Error fetching stages:", error);
//       setStages([]);
//     }
//   };

//   //
//   const handleCreatePipeline = async () => {
//     if (!newPipelineName.trim()) {
//       alert("Please enter a pipeline name");
//       return;
//     }

//     try {
//       await createPipeline(newPipelineName);
//       setNewPipelineName(""); // Reset input field
//       loadPipelines(); // Refresh pipeline list after creation
//     } catch (error) {
//       alert("Failed to create pipeline");
//     }
//   };
//   //

//   return (
//     <div className="dashboard">
//       {/* 🔹 Top Bar */}
//       <div className="top-bar">
//         <h1 className="project-name">Manufacturing Pipeline</h1>
//         <div className="user-profile">
//           <img src="/profile-icon.png" alt="User Profile" className="profile-icon" />
//         </div>
//       </div>

//       {/* 🔹 Debugging Messages */}
//       {loading && <p className="loading-message">Loading pipelines...</p>}
//       {error && <p className="error-message">Error: {error}</p>}

//       {/* 🔹 Create Pipeline Section */}
//       <div className="create-pipeline">
//         <input
//           type="text"
//           placeholder="Enter pipeline name"
//           value={newPipelineName}
//           onChange={(e) => setNewPipelineName(e.target.value)}
//           className="pipeline-input"
//         />
//         <button onClick={handleCreatePipeline} className="create-button">
//           Create Pipeline
//         </button>
//       </div>

//       {/* 🔹 Main Content (Pipelines + Stages) */}
//       <div className={`dashboard-content ${clicked ? "shift-left" : ""}`}>
        
//         {/* 🔹 Left Side: Pipeline Cards */}
//         <div className="pipeline-list">
//           {pipelines.length === 0 && !loading ? (
//             <p className="no-pipelines">No pipelines available</p>
//           ) : (
//             pipelines.map((pipeline) => (
//               <PipelineCard
//                 key={pipeline.id}
//                 pipeline={pipeline}
//                 onClick={() => handlePipelineClick(pipeline.id)}
//                 isSelected={pipeline.id === selectedPipeline}
//               />
//             ))
//           )}
//         </div>

//         {/* 🔹 Right Side: Stage List (Appears when a pipeline is clicked) */}
//         {selectedPipeline && (
//           <div className="stages-container">
//             <StageList stages={stages} />
//           </div>
//         )}

//       </div>
//     </div>
//   );
// };

// export default Dashboard;
