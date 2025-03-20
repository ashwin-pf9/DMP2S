import React, { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import { Trash2 } from "lucide-react";
import { fetchStages, startPipeline, deletePipeline } from "../api/api";
import "../styles/pipeline_stages.css";

const PipelineStages = () => {
  const { id } = useParams();
  const [stages, setStages] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const routeLocation = useLocation();
  const navigate = useNavigate();
  const [pipelineName] = useState(routeLocation.state?.pipelineName || "Pipeline");

  useEffect(() => {
    const getStages = async () => {
      try {
        setLoading(true);
        setError(null);
        const data = await fetchStages(id);
        if (!data || !Array.isArray(data)) {
          throw new Error("Invalid stage data format");
        }
        setStages(data);
      } catch (error) {
        setError(error.message);
      } finally {
        setLoading(false);
      }
    };
    getStages();
  }, [id]);

  useEffect(() => {
    // Establish WebSocket connection
    const socket = new WebSocket("ws://localhost:30080/ws/status-updates");

    socket.onmessage = (event) => {
      const update = JSON.parse(event.data);

      setStages((prevStages) =>
        prevStages.map((stage) =>
          stage.id === update.stage_id ? { ...stage, status: update.status } : stage
        )
      );
    };

    socket.onerror = (error) => {
      console.error("WebSocket Error:", error);
    };

    return () => {
      socket.close(); // Clean up WebSocket connection on component unmount
    };
  }, []);

  const handleStartPipeline = async () => {
    try {
      await startPipeline(id);
    } catch (error) {
      if (error.message === "Unauthorized") {
        alert("Session expired. Redirecting to login...");
        navigate("/login");
      } else {
        alert(`Failed to start pipeline ${id}: ${error.message}`);
      }
    }
  };

  const handleDeletePipeline = async () => {
    try {
      await deletePipeline(id);
    } catch (error) {
      if (error.message === "Unauthorized") {
        alert("Session expired. Redirecting to login...");
        navigate("/login");
      } else {
        alert(`Failed to delete pipeline ${id}: ${error.message}`);
      }
    }
  }

  return (
    <div className="pipeline-stages">
      <h2>{pipelineName}</h2>
      {loading && <p>Loading stages...</p>}
      {error && <p className="error">{error}</p>}
      <ul>
        {stages.map((stage) => (
          <li key={stage.id} className="stage-card">
            <h3>{stage.name}</h3>
            <p>Status: <span className="stage-status">{stage.status || "Pending"}</span></p>
          </li>
        ))}
      </ul>

      {/* <button className="start-button" onClick={handleStartPipeline}>Start Pipeline</button> */}
      <div className="buttons-container">
        <button className="start-button" onClick={handleStartPipeline}>Start Pipeline</button>
        <button className="delete-button" onClick={handleDeletePipeline}>
          <Trash2 size={20} /> Delete Pipeline
        </button>
      </div>
    </div>
  );
};

export default PipelineStages;
