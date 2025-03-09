import React from "react";
import "../styles/stage_list.css"; // Make sure to style it properly

const StageList = ({ stages }) => {
  if (!stages || stages.length === 0) {
    return <p className="no-stages">No stages available</p>;
  }

  return (
    <div className="stage-list">
      {/* <h3>Pipeline Stages</h3>
      <ul>
        {stages.map((stage) => (
          <li key={stage.id} className="stage-card">
            <h4>{stage.name}</h4>
            <p>{stage.description || "No description available"}</p>
            <span className="stage-status">{stage.status}</span>
          </li>
        ))}
      </ul> */}
    </div>
  );
};

export default StageList;
