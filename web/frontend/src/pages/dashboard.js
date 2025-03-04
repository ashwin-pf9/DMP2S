import React, { useEffect, useState } from "react";
import { fetchPipelines, fetchStages } from "../api/api";
import PipelineCard from "../components/pipeline_card";
import StageList from "../components/stage_list";
import "../styles/dashboard.css";

const Dashboard = () => {
  const [pipelines, setPipelines] = useState([]);
  const [selectedPipeline, setSelectedPipeline] = useState(null);
  const [stages, setStages] = useState([]);

  useEffect(() => {
    const loadPipelines = async () => {
      try {
        const data = await fetchPipelines();
        setPipelines(data);
      } catch (error) {
        console.error("Error fetching pipelines:", error);
      }
    };
    loadPipelines();
  }, []);

  const handlePipelineClick = async (pipelineId) => {
    setSelectedPipeline(pipelineId);
    try {
      const data = await fetchStages(pipelineId);
      setStages(data);
    } catch (error) {
      console.error("Error fetching stages:", error);
    }
  };

  return (
    <div className="dashboard">
      <h2>Dashboard</h2>
      <div className="pipeline-list">
        {pipelines.map((pipeline) => (
          <PipelineCard
            key={pipeline.id}
            pipeline={pipeline}
            onClick={handlePipelineClick}
          />
        ))}
      </div>
      {selectedPipeline && <StageList stages={stages} />}
    </div>
  );
};

export default Dashboard;













// import React, { useEffect, useState } from "react";
// import { useNavigate } from "react-router-dom";
// import API_BASE_URL from "../config";

// const Dashboard = () => {
//   console.log("dashboard function called..")
//   const [pipelines, setPipelines] = useState([]);
//   const navigate = useNavigate();

//   useEffect(() => {
//     const fetchPipelines = async () => {
//       const token = localStorage.getItem("token");
//       if (!token) {
//         navigate("/login");
//         return;
//       }

//       try {
//         const response = await fetch(`${API_BASE_URL}/pipelines`, {
//           headers: { Authorization: `Bearer ${token}` },
//         });

//         if (!response.ok) throw new Error("Failed to fetch pipelines");

//         const data = await response.json();
//         setPipelines(data);
//       } catch (error) {
//         console.error("Error fetching pipelines:", error);
//       }
//     };

//     fetchPipelines();
//   }, [navigate]);

//   return (
//     <div>
//       <h2>Dashboard</h2>
//       <ul>
//         {pipelines.map((pipeline) => (
//           <li key={pipeline.id}>{pipeline.name}</li>
//         ))}
//       </ul>
//     </div>
//   );
// };

// export default Dashboard;
