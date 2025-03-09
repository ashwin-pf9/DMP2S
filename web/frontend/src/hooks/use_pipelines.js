import { useState, useEffect } from "react";
import { fetchPipelines, fetchStages, createPipeline } from "../api/api";
import { useNavigate } from "react-router-dom";

const usePipelines = () => {
  const [pipelines, setPipelines] = useState([]);
  const [selectedPipeline, setSelectedPipeline] = useState(null);
  const [stages, setStages] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    loadPipelines();
  }, []);

  const loadPipelines = async () => {
    try {
      setLoading(true);
      setError(null);
      const data = await fetchPipelines();
      if (!Array.isArray(data)) throw new Error("Invalid pipeline data format");
      setPipelines(data);
    } catch (error) {
      setError(error.message);
      setPipelines([]);
    } finally {
      setLoading(false);
    }
  };

  const handlePipelineClick = async (pipelineId) => {
    if (selectedPipeline === pipelineId) return;
    setSelectedPipeline(pipelineId);

    try {
      const data = await fetchStages(pipelineId);
      if (!Array.isArray(data)) throw new Error("Invalid stage data format");
      setStages(data);
    } catch {
      setStages([]);
    }
  };

  const handleCreatePipeline = async (name) => {
    if (!name.trim()) return;
    try {
      await createPipeline(name);
      loadPipelines(); // Refresh list
    } catch (error) {
      if (error.message === "Unauthorized") {
          alert("Session expired. Redirecting to login...");
          navigate("/login");
        } else {
          alert(`Failed to create pipeline ${name}: ${error.message}`);
        }
      }
  };

  return {
    pipelines,
    selectedPipeline,
    stages,
    loading,
    error,
    handlePipelineClick,
    handleCreatePipeline,
  };
};

export default usePipelines;
