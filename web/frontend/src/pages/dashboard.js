import React, { useEffect, useState } from "react";
import { fetchPipelines, fetchStages, createPipeline } from "../api/api";
import PipelineCard from "../components/pipeline_card";
import StageList from "../components/stage_list";
import "../styles/dashboard.css";

const Dashboard = () => {
  const [pipelines, setPipelines] = useState([]);
  const [selectedPipeline, setSelectedPipeline] = useState(null);
  const [stages, setStages] = useState([]);
  const [clicked, setClicked] = useState(false);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [newPipelineName, setNewPipelineName] = useState("");

  useEffect(() => {
    loadPipelines();
  }, []);
    const loadPipelines = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await fetchPipelines();
        console.log("Fetched Pipelines:", data);

        if (!data || !Array.isArray(data)) {
          throw new Error("Invalid pipeline data format");
        }

        setPipelines(data);
      } catch (error) {
        console.error("Error fetching pipelines:", error);
        setError(error.message);
        setPipelines([]);
      } finally {
        setLoading(false);
      }
    };

   

  const handlePipelineClick = async (pipelineId) => {
    if (selectedPipeline === pipelineId) return;

    setSelectedPipeline(pipelineId);
    setClicked(true);

    try {
      const data = await fetchStages(pipelineId);
      console.log(`Fetched stages for pipeline ${pipelineId}:`, data);

      if (!data || !Array.isArray(data)) {
        throw new Error("Invalid stage data format");
      }

      setStages(data);
    } catch (error) {
      console.error("Error fetching stages:", error);
      setStages([]);
    }
  };

  //
  const handleCreatePipeline = async () => {
    if (!newPipelineName.trim()) {
      alert("Please enter a pipeline name");
      return;
    }

    try {
      await createPipeline(newPipelineName);
      setNewPipelineName(""); // Reset input field
      loadPipelines(); // Refresh pipeline list after creation
    } catch (error) {
      alert("Failed to create pipeline");
    }
  };
  //

  return (
    <div className="dashboard">
      {/* 🔹 Top Bar */}
      <div className="top-bar">
        <h1 className="project-name">Manufacturing Pipeline</h1>
        <div className="user-profile">
          <img src="/profile-icon.png" alt="User Profile" className="profile-icon" />
        </div>
      </div>

      {/* 🔹 Debugging Messages */}
      {loading && <p className="loading-message">Loading pipelines...</p>}
      {error && <p className="error-message">Error: {error}</p>}

      {/* 🔹 Create Pipeline Section */}
      <div className="create-pipeline">
        <input
          type="text"
          placeholder="Enter pipeline name"
          value={newPipelineName}
          onChange={(e) => setNewPipelineName(e.target.value)}
          className="pipeline-input"
        />
        <button onClick={handleCreatePipeline} className="create-button">
          Create Pipeline
        </button>
      </div>

      {/* 🔹 Main Content (Pipelines + Stages) */}
      <div className={`dashboard-content ${clicked ? "shift-left" : ""}`}>
        
        {/* 🔹 Left Side: Pipeline Cards */}
        <div className="pipeline-list">
          {pipelines.length === 0 && !loading ? (
            <p className="no-pipelines">No pipelines available</p>
          ) : (
            pipelines.map((pipeline) => (
              <PipelineCard
                key={pipeline.id}
                pipeline={pipeline}
                onClick={() => handlePipelineClick(pipeline.id)}
                isSelected={pipeline.id === selectedPipeline}
              />
            ))
          )}
        </div>

        {/* 🔹 Right Side: Stage List (Appears when a pipeline is clicked) */}
        {selectedPipeline && (
          <div className="stages-container">
            <StageList stages={stages} />
          </div>
        )}

      </div>
    </div>
  );
};

export default Dashboard;
