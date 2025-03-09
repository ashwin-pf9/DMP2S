import React from "react";
import StageList from "./stage_list";

const StageViewer = ({ stages }) => {
  return (
    <div className="stages-container">
      <StageList stages={stages} />
    </div>
  );
};

export default StageViewer;
