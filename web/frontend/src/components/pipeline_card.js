import React from "react";
import "../styles/pipeline_card.css";

const PipelineCard = ({ pipeline, isSelected, onClick }) => {
  return (
    <div
      className={`pipeline-card ${isSelected ? "selected" : ""}`}
      onClick={onClick}
    >
      {pipeline.name}
    </div>
  );
};

export default PipelineCard;





// import React from "react";
// import "../styles/dashboard.css";

// const PipelineCard = ({ pipeline, onClick }) => {
//   return (
//     <div className="pipeline-card" onClick={() => onClick(pipeline.id)}>
//       <h3>{pipeline.name}</h3>
//       <p>{pipeline.description || "No description available"}</p>
//     </div>
//   );
// };

// export default PipelineCard;
