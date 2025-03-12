import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createPipeline, createStage } from "../api/api"; // API calls
import "../styles/pipeline_creator.css";


const PipelineCreator = () => {
  const [pipelineName, setPipelineName] = useState("");
  const [stages, setStages] = useState([]);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  // Add new stage fields dynamically
  const handleAddStage = () => {
    setStages([...stages, { name: "", description: "" }]);
  };

  // Handle stage field changes
  const handleStageChange = (index, field, value) => {
    const updatedStages = [...stages];
    updatedStages[index][field] = value;
    setStages(updatedStages);
  };

  // Handle form submission
  const handleSubmit = async () => {
    if (!pipelineName.trim()) {
      alert("Please enter a pipeline name.");
      return;
    }

    setLoading(true);

    try {
      // Step 1: Create pipeline
      const pipeline = await createPipeline(pipelineName);
      const pipelineId = pipeline.id;

      // Step 2: Add stages
      for (const stage of stages) {
        if (stage.name.trim()) {
          await createStage(pipelineId, stage);
        }
      }

      alert(`Pipeline \"${pipelineName}\" created successfully!`);
      navigate("/dashboard"); // Redirect to dashboard after success
    } catch (error) {
      alert(`Failed to create pipeline: ${error.message}`);
    }

    setLoading(false);
  };

  return (
    <div className="pipeline-creator-container">
      <h2>Create Pipeline</h2>
      <input
        type="text"
        placeholder="Enter pipeline name"
        value={pipelineName}
        onChange={(e) => setPipelineName(e.target.value)}
        className="pipeline-input"
      />

      <h3>Stages</h3>
      {stages.map((stage, index) => (
        <div key={index} className="stage-input">
          <input
            type="text"
            placeholder="Stage Name"
            value={stage.name}
            onChange={(e) => handleStageChange(index, "name", e.target.value)}
          />
          {/* <input
            type="text"
            placeholder="Stage Description"
            value={stage.description}
            onChange={(e) => handleStageChange(index, "description", e.target.value)}
          /> */}
        </div>
      ))}

      <button onClick={handleAddStage} className="add-stage-button">+ Add Stage</button>
      <button onClick={handleSubmit} className="submit-button" disabled={loading}>
        {loading ? "Creating..." : "Create Pipeline"}
      </button>
    </div>
  );
};

export default PipelineCreator;



// import React, { useState } from "react";
// import "../styles/pipeline_creator.css";

// const PipelineCreator = ({ onCreate }) => {
//   const [showForm, setShowForm] = useState(false);
//   const [newPipelineName, setNewPipelineName] = useState("");

//   const handleSubmit = () => {
//     if (newPipelineName.trim()) {
//       onCreate(newPipelineName);
//       setNewPipelineName("");
//       setShowForm(false); // Hide form after submission
//     } else {
//       alert("Please enter a valid pipeline name");
//     }
//   };

//   return (
//     <div>
//       {/* Fixed button at bottom-left */}
//       <button onClick={() => setShowForm(!showForm)} className="toggle-button">
//         {showForm ? "×" : "Create Pipeline"}
//       </button>

//       {/* Pipeline creation form (shown when button is clicked) */}
//       {showForm && (
//         <div className="pipeline-form-container">
//           <div className="pipeline-form">
//             <h3>Create New Pipeline</h3>
//             <input
//               type="text"
//               placeholder="Enter pipeline name"
//               value={newPipelineName}
//               onChange={(e) => setNewPipelineName(e.target.value)}
//               className="pipeline-input"
//             />
//             <div className="form-actions">
//               <button onClick={handleSubmit} className="submit-button">
//                 Submit
//               </button>
//               <button onClick={() => setShowForm(false)} className="cancel-button">
//                 Cancel
//               </button>
//             </div>
//           </div>
//         </div>
//       )}
//     </div>
//   );
// };

// export default PipelineCreator;
