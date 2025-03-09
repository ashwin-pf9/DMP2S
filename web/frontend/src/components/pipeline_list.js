import React from "react";

const PipelineList = ({ pipelines, onSelect }) => {
  return (
    <div className="pipeline-list">
      {pipelines.map((pipeline) => (
        <div
          key={pipeline.id}
          className="pipeline-card"
          onClick={() => onSelect(pipeline.id,pipeline.name)}
        >
          <h3>{pipeline.name}</h3>
        </div>
      ))}
    </div>
  );
};

export default PipelineList;


// import React from "react";
// import PipelineCard from "./pipeline_card";

// const PipelineList = ({ pipelines, onSelect, selectedPipeline }) => {
//   return (
//     <div className="pipeline-list">
//       {pipelines.length === 0 ? (
//         <p className="no-pipelines">No pipelines available</p>
//       ) : (
//         pipelines.map((pipeline) => (
//           <PipelineCard
//             key={pipeline.id}
//             pipeline={pipeline}
//             onClick={() => onSelect(pipeline.id)}
//             isSelected={pipeline.id === selectedPipeline}
//           />
//         ))
//       )}
//     </div>
//   );
// };

// export default PipelineList;
