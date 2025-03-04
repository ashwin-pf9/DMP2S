import React from "react";
import "../styles/dashboard.css";

const StageList = ({ stages }) => {
  return (
    <div className="stages-container">
      <h3>Stages</h3>
      {stages.length === 0 ? (
        <p>No stages found</p>
      ) : (
        <ul>
          {stages.map((stage) => (
            <li key={stage.id}>{stage.name}</li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default StageList;
